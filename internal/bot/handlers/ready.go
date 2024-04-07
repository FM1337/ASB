package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (h *Handlers) Ready(s *discordgo.Session, r *discordgo.Ready) {
	// Bot is loaded and connected

	for _, srv := range s.State.Guilds {
		if !h.data.ServerExists(srv.ID) {
			// server doesn't exist, add it
			go func(serverId, ownerId string) {
				err1, err2 := h.data.AddServer(serverId, ownerId)
				if err1 != nil {
					// todo set an embed with info
					s.ChannelMessageSend(h.config.ErrorChannel, fmt.Sprintf("error adding server %s", err1.Error()))
				}
				if err2 != nil {
					// todo set an embed with info
					s.ChannelMessageSend(h.config.ErrorChannel, fmt.Sprintf("extra error during adding server %s", err2.Error()))
				}
			}(srv.ID, srv.OwnerID)
			// if an error occurs, we'll need to add a command later for trying to add it manually
		}
	}

	h.logger.Infof("Bot is online and connected to %d servers", len(s.State.Guilds))
}
