package handlers

import "github.com/bwmarrin/discordgo"

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

}
