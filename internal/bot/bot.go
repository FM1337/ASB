package bot

import (
	"fmt"

	"github.com/FM1337/ASB/internal/bot/handlers"
	"github.com/FM1337/ASB/internal/bot/memory"
	"github.com/FM1337/ASB/internal/models"
	"github.com/apex/log"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session  *discordgo.Session
	logger   log.Interface
	config   *models.ConfigDiscord
	handlers *handlers.Handlers
}

func NewBot(lg log.Interface, config *models.ConfigDiscord, data *memory.Data) *Bot {
	return &Bot{
		logger:   lg,
		config:   config,
		handlers: handlers.InitHandlers(lg, config, data),
	}
}

func (b *Bot) Start() (err error) {
	b.session, err = discordgo.New(fmt.Sprintf("Bot %s", b.config.Token))

	if err != nil {
		return err
	}

	// Set up intents
	b.session.Identify.Intents = discordgo.IntentMessageContent | discordgo.IntentsAllWithoutPrivileged

	// Add handlers
	b.session.AddHandler(b.handlers.Ready)
	b.session.AddHandler(b.handlers.MessageCreate)
	b.session.AddHandler(b.handlers.InteractionCreate)

	return err
}

func (b *Bot) Stop() {
}
