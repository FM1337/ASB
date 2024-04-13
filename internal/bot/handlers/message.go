package handlers

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strings"

	"github.com/FM1337/ASB/internal/models"
	"github.com/bwmarrin/discordgo"
	"mvdan.cc/xurls/v2"
)

func (h *Handlers) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.GuildID == "" {
		// This is a DM, which means we'll need to do some special handling.
		// TODO for now just don't do anything
		return
	}

	// Check if the bot is enabled for the server or if it should ignore the message

	if !h.data.Enabled(m.GuildID) || !h.data.ShouldIgnore(m.GuildID, m.ChannelID, m.Author.ID, m.Member.Roles) {
		// Do nothing
		return
	}

	// used to determine the level of suspiciousness the message/author is at.
	susLevel := 0
	// if this gets set to true, we stop all remaining checks and take an action
	instaAction := false

	// grab the config
	cfg := h.data.GetConfiguration(m.GuildID)

	// If enabled, messages containing links should increase the sus level
	if cfg.FlagLinks && xurls.Relaxed().MatchString(m.Content) {
		susLevel++
	}

	// we should check for blacklisted words and increase if any are found
	if h.data.CheckForBlacklistedContent(m.Content, m.GuildID) {
		susLevel++
	}

	if susLevel == 0 {
		// At this point if the level is still 0, it's more than likely not spam so we can bail here
		return
	}

	// if the sus level isn't 0 and ratelimit is turned on, let's start that magic

	if cfg.EnforceRatelimit {
		// hash the message
		hash := fmt.Sprintf("%x", md5.Sum([]byte(strings.ToLower(m.Content))))
		seenBefore, count := h.data.MessageSeenBefore(m.GuildID, m.Author.ID, hash)
		if count+1 >= cfg.RateLimitCount {
			// User has hit the limit, we need to action now
			instaAction = true
		} else if seenBefore {
			// okay we've seen this messag before from this user in the server within the
			// rate limit time, let's reset the timer for them and increment their count
		} else {
			// first time seeing this message, let's create a new entry for them.
		}
	}

	if instaAction {
		// Take action now.
		doAction(s, m, &cfg)
		return
	}

	if cfg.CheckInvites {
		// check for an invite
		hasInvite, inviteCode := grabDiscordLink(m.Content)
		if hasInvite {
			// get information about the invite
			invite, err := s.Invite(inviteCode)
			if err != nil {
				// error handling
				return
			}
			// increase the sus level by 1 just for containing an invite
			susLevel++

			if h.data.CheckForBlacklistedContent(invite.Guild.Name, m.ChannelID) {
				// Blackedlisted content found, sus goes up by 1
				susLevel++
			}
		}
	}

}

func grabDiscordLink(msg string) (bool, string) {
	r, _ := regexp.Compile("discord.gg/[a-zA-z0-9]{1,10}")

	match := r.FindString(strings.ToLower(msg))

	split := strings.Split(match, "/")

	if len(split) == 2 {
		match = split[1]
	}

	return match != "", match
}

func doAction(s *discordgo.Session, m *discordgo.MessageCreate, cfg *models.ServerConfig) {
	if cfg.Ban {
		// Note until discordgo implements the ban delete message seconds query param, we'll just default to 1 day
		// Easiest action to take, we ban and clear

		err := s.GuildBanCreateWithReason(m.ChannelID, m.Author.ID, "User was flagged as a spammer by security bot", 1)
		if err != nil {
			// likely a permission error, log it anyway (TODO logging)
		}

		return
	}
}
