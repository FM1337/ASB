package bot

import (
	"fmt"

	"github.com/FM1337/ASB/internal/ent"
	"github.com/FM1337/ASB/internal/models"
	"github.com/apex/log"
	"github.com/bwmarrin/discordgo"
	"github.com/redis/go-redis/v9"
)

type Bot struct {
	session *discordgo.Session
	logger  log.Interface
	db      *ent.Client
	rdb     *redis.Client
	config  models.ConfigDiscord
}

func NewBot(lg log.Interface, db *ent.Client, rdb *redis.Client, config models.ConfigDiscord) *Bot {
	return &Bot{
		logger: lg,
		db:     db,
		config: config,
		rdb:    rdb,
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
	b.session.AddHandler(b.messageCreate)

	return err
}

func (b *Bot) Stop() {
}

func (b *Bot) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

}
