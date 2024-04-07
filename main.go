package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/FM1337/ASB/internal/bot"
	"github.com/FM1337/ASB/internal/ent"
	"github.com/FM1337/ASB/internal/models"
	"github.com/apex/log"
	"github.com/lrstanley/clix"
)

var (
	cli = &clix.CLI[models.Flags]{
		Links: clix.GithubLinks("github.com/FM1337/ASB", "main", "https://github.com/FM1337/ASB"),
	}
	logger log.Interface

	db  *ent.Client
	asb *bot.Bot
)

func init() {
	cli.Parse()
	logger = cli.Logger

	if !cli.Flags.Configured {
		logger.Fatal("Bot not configured")
	}
}

func main() {
	// Run the required setup tasks
	setup()

	err := asb.Start()

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
		asb.Stop()
		db.Close()
		logger.Info("done cleaning up; exiting")
		os.Exit(0)
	case <-done:
	}
}
