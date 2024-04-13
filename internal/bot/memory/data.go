package memory

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/FM1337/ASB/internal/ent"
	"github.com/FM1337/ASB/internal/ent/cooldown"
	"github.com/FM1337/ASB/internal/ent/server"
	"github.com/FM1337/ASB/internal/ent/serverconfig"
	"github.com/FM1337/ASB/internal/models"
)

type Data struct {
	words         []string
	blacklist     map[string][]int
	configuration map[string]models.ServerConfig
	enabled       map[string]bool
	cooldown      map[string][]models.Cooldown
	db            *ent.Client
}

// Loads the bot data from the database (blacklist, server config, etc)
func LoadData(db *ent.Client) (*Data, error) {
	data := &Data{
		words:         []string{},
		blacklist:     make(map[string][]int),
		configuration: make(map[string]models.ServerConfig),
		cooldown:      make(map[string][]models.Cooldown),
		db:            db,
	}

	words, err := db.WordBlacklist.Query().WithServer(func(sq *ent.ServerQuery) {
		sq.WithConfiguration()
	}).All(context.Background())
	if err != nil {
		return nil, err
	}

	for _, word := range words {
		index := slices.Index(data.words, word.Word)
		if index == -1 {
			data.words = append(data.words, word.Word)
			index = len(data.words) - 1
		}

		for _, srv := range word.Edges.Server {
			if _, ok := data.blacklist[srv.ServerID]; !ok {
				// if there's no existing array, set one and make the first value the index
				data.blacklist[srv.ServerID] = []int{index}
			} else if !slices.Contains(data.blacklist[srv.ServerID], index) {
				// if the index isn't already in the array, append it
				data.blacklist[srv.ServerID] = append(data.blacklist[srv.ServerID], index)
			}

			// Since we're here, might as well set the flag to see if the server is enabled or not
			data.enabled[srv.ServerID] = srv.Enabled

			// get the server config from the database
			config, err := db.ServerConfig.Query().Where(serverconfig.HasServerWith(server.ServerID(srv.ServerID))).First(context.Background())
			if err != nil {
				// if a server is stored without a config, we got a major problem
				return nil, err
			}

			data.configuration[srv.ServerID] = models.ServerConfig{
				RemoveRoles:          config.RemoveRoles,
				GiveRole:             config.GiveRole,
				Timeout:              config.Timeout,
				Kick:                 config.Kick,
				Ban:                  config.Ban,
				CheckInvites:         config.CheckInvites,
				CheckLinks:           config.CheckInvites,
				EnforceRatelimit:     config.Ratelimit,
				SendAlerts:           config.Alerts,
				FlagLinks:            config.FlagLinks,
				LogChannel:           config.LogChannel,
				GivenRole:            config.GivenRole,
				ExcludedChannels:     config.ExcludedChannels,
				ExcludedRoles:        config.ExcludedRoles,
				ExcluedUsers:         config.ExcludedUsers,
				RateLimitCount:       config.RatelimitMessage,
				RateLimitTime:        config.RatelimitTime,
				TimeoutTime:          config.TimeoutTime,
				BanMessageDeleteTime: config.BanDeleteMessageTime,
			}

			// get the cooldown info for the server from the database
			data.cooldown[srv.ServerID] = []models.Cooldown{}

			cooldowns, err := db.Cooldown.Query().Where(cooldown.HasServerWith(server.ServerID(srv.ServerID))).All(context.Background())
			if err != nil {
				return nil, err
			}

			for _, cooldown := range cooldowns {
				data.cooldown[srv.ServerID] = append(data.cooldown[srv.ServerID],
					models.Cooldown{
						UserId:  cooldown.UserID,
						HashId:  cooldown.Hash,
						Count:   cooldown.Count,
						ResetAt: cooldown.ResetsAt,
					},
				)
			}

		}
	}

	return data, nil
}

func (d *Data) AddServer(serverId, ownerId string) (error, error) {
	// Create a server config
	config, err := d.db.ServerConfig.Create().Save(context.Background())
	if err != nil {
		return err, nil
	}

	// Create the server
	_, err = d.db.Server.Create().
		SetServerID(serverId).
		SetOwnerID(ownerId).
		SetConfiguration(config).
		Save(context.Background())

	if err != nil {
		// remove the config that was just created
		err2 := d.db.ServerConfig.DeleteOne(config).Exec(context.Background())
		return err, err2
	}

	// Store in memory

	d.configuration[serverId] = models.ServerConfig{
		RemoveRoles:          config.RemoveRoles,
		GiveRole:             config.GiveRole,
		Timeout:              config.Timeout,
		Kick:                 config.Kick,
		Ban:                  config.Ban,
		CheckInvites:         config.CheckInvites,
		CheckLinks:           config.CheckInvites,
		EnforceRatelimit:     config.Ratelimit,
		SendAlerts:           config.Alerts,
		FlagLinks:            config.FlagLinks,
		LogChannel:           config.LogChannel,
		GivenRole:            config.GivenRole,
		ExcludedChannels:     config.ExcludedChannels,
		ExcludedRoles:        config.ExcludedRoles,
		ExcluedUsers:         config.ExcludedUsers,
		RateLimitCount:       config.RatelimitMessage,
		RateLimitTime:        config.RatelimitTime,
		TimeoutTime:          config.TimeoutTime,
		BanMessageDeleteTime: config.BanDeleteMessageTime,
	}

	d.enabled[serverId] = false
	d.blacklist[serverId] = []int{}
	d.cooldown[serverId] = []models.Cooldown{}

	return nil, nil
}

func (d *Data) ServerExists(serverId string) bool {
	_, ok := d.configuration[serverId]
	return ok
}

func (d *Data) EnableServer(serverId string) (bool, error) {
	if !d.ServerExists(serverId) {
		return false, fmt.Errorf("your server doesn't exist in the bot's memory")
	}

	// check to see if the bot is already enabled

	if d.enabled[serverId] {
		return false, fmt.Errorf("bot already enabled")
	}

	// grab the config
	cfg := d.GetConfiguration(serverId)

	// this is to determine the highest possible level of sus that can
	// be detected based on enabled checks
	highestPossibleSusLevel := 0

	hasBlacklist := len(d.blacklist[serverId]) > 0

	if hasBlacklist {
		highestPossibleSusLevel++
	}

	if cfg.CheckLinks {
		highestPossibleSusLevel++
	}

	if cfg.CheckInvites {
		highestPossibleSusLevel++

		// if there's a blacklist, there's a possibility of the invite server name being in there.
		if hasBlacklist {
			highestPossibleSusLevel++
		}
	}

	if cfg.EnforceRatelimit {
		highestPossibleSusLevel++
	}

	if cfg.FlagLinks {
		highestPossibleSusLevel++
	}

	// We need at 3, if not then the bot can't reliably flag spammers

	if highestPossibleSusLevel < 3 {
		return false, fmt.Errorf("security configured too low, high chance of false flagging or missing actual spammers")
	}

	return d.changeServerStatus(serverId, true)
}

func (d *Data) DisableServer(serverId string) (bool, error) {
	if !d.ServerExists(serverId) {
		return false, fmt.Errorf("your server doesn't exist in the bot's memory")
	}

	// check to see if the bot is already disabled
	if !d.enabled[serverId] {
		return false, fmt.Errorf("bot already disabled")
	}

	return d.changeServerStatus(serverId, false)
}

func (d *Data) changeServerStatus(serverId string, status bool) (bool, error) {
	srv, err := d.db.Server.Query().Where(server.ServerID(serverId)).First(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return false, fmt.Errorf("server missing from database")
		}
		// TODO log the error to the error channel/probably sentry but don't send the exposed error
		return false, fmt.Errorf("unexpected server error occurred while fetching server")
	}

	_, err = srv.Update().SetEnabled(status).Save(context.TODO())

	if err != nil {
		return false, fmt.Errorf("unexpected server error occurred while saving server enabled status")
	}

	// update memory db
	d.enabled[serverId] = status

	return true, nil
}

func (d *Data) UpdateConfiguration(serverId string, config models.ServerConfig) {}

func (d *Data) GetConfiguration(serverId string) models.ServerConfig {
	return d.configuration[serverId]
}

// Checks to see if the provided content contains a blacklisted word or phrase
func (d *Data) CheckForBlacklistedContent(content, serverId string) bool {
	// convert the content to lower case
	content = strings.ToLower(content)

	// loop through the serverId's blacklist content to see if the content contains any of em
	for _, index := range d.blacklist[serverId] {
		if strings.Contains(content, d.words[index]) {
			return true
		}
	}

	return false
}

func (d *Data) Enabled(serverId string) bool {
	// extra error handling
	if !d.ServerExists(serverId) {
		return false
	}

	return d.enabled[serverId]
}

func (d *Data) ShouldIgnore(serverId, channelId, userId string, roleIds []string) bool {
	cfg := d.configuration[serverId]

	// Check if the message is in a channel or is a user for a server that we should ignore
	if slices.Contains(cfg.ExcludedChannels, channelId) || slices.Contains(cfg.ExcluedUsers, userId) {
		return true
	}

	// check for er
	return slices.ContainsFunc(cfg.ExcludedRoles, func(roleId string) bool {
		return slices.Contains(roleIds, roleId)
	})

}

func (d *Data) MessageSeenBefore(serverId, userId, hash string) (bool, int) {
	cooldowns := d.cooldown[serverId]
	index := slices.IndexFunc(cooldowns, func(c models.Cooldown) bool {
		return c.HashId == hash && c.UserId == userId
	})

	if index == -1 {
		return false, 0
	}

	return true, cooldowns[index].Count
}

// Cleans up unused data (such as unused words) TODO
func (d *Data) cleanup() {
}
