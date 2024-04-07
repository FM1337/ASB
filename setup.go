package main

import (
	"context"

	"github.com/FM1337/ASB/internal/bot"
	"github.com/FM1337/ASB/internal/bot/memory"
	"github.com/FM1337/ASB/internal/database"
	"github.com/FM1337/ASB/internal/ent"
	"github.com/apex/log"
)

func setup() {
	// Open a connection to the database and run any migrations
	// that need to occur
	db = database.Open(logger, cli.Flags.DB)
	ctx := ent.NewContext(log.NewContext(context.Background(), logger), db)
	database.Migrate(ctx, logger)

	// Load configurations from the database into memory
	data, err := memory.LoadData(db)
	if err != nil {
		logger.WithError(err).Fatal("Failed to load data")
	}

	// Create the bot
	asb = bot.NewBot(logger, &cli.Flags.Discord, data)
}
