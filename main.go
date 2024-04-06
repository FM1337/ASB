package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/FM1337/ASB/internal/bot"
	"github.com/FM1337/ASB/internal/database"
	"github.com/FM1337/ASB/internal/ent"
	"github.com/FM1337/ASB/internal/models"
	"github.com/apex/log"
	"github.com/lrstanley/clix"
	"github.com/redis/go-redis/v9"
)

var (
	cli = &clix.CLI[models.Flags]{
		Links: clix.GithubLinks("github.com/FM1337/ASB", "main", "https://github.com/FM1337/ASB"),
	}
	logger log.Interface

	db  *ent.Client
	rdb *redis.Client
)

func init() {
	cli.Parse()
	logger = cli.Logger

	if !cli.Flags.Configured {
		logger.Warn("Bot not configured")
		os.Exit(1)
	}

	db = database.Open(logger, cli.Flags.DB)
	ctx := ent.NewContext(log.NewContext(context.Background(), logger), db)
	database.Migrate(ctx, logger)

	rdb = redis.NewClient(&redis.Options{
		Addr:     cli.Flags.Redis.Addr,
		Password: cli.Flags.Redis.Password,
		DB:       cli.Flags.Redis.DB,
	})

	// Ping to verify we connected successfully
	_, err := rdb.Ping(context.Background()).Result()

	if err != nil {
		logger.WithError(err).Fatal("Failed to ping redis server")
	}
}

func main() {

	b := bot.NewBot(logger, db, rdb, cli.Flags.Discord)

	err := b.Start()

	if err != nil {
		logger.WithError(err).Fatal("Failed to start bot")
	}

	// Wait here until CTRL-C or other term signal is received.
	done := make(chan struct{})

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-signals:
		logger.Info("received SIGINT/SIGTERM/SIGQUIT, closing connections...")
		// Shut down the bot followed by the databases
		b.Stop()
		db.Close()
		logger.Info("done cleaning up; exiting")
		os.Exit(0)
	case <-done:
	}
}
