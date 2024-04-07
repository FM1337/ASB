package handlers

import "github.com/bwmarrin/discordgo"

func (h *Handlers) Join(s *discordgo.Session, g *discordgo.GuildMemberAdd) {
	// TODO: do we want to do something if a spammer rejoins?
}
