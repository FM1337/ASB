package database

import (
	"context"
	"database/sql"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/FM1337/ASB/internal/ent"
	"github.com/FM1337/ASB/internal/models"
	"github.com/apex/log"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Open new postgres connection.
func Open(logger log.Interface, config models.ConfigDB) *ent.Client {
	var db *sql.DB
	var err error

	var attempt int
	for {
		attempt++
		db, err = sql.Open("pgx", config.URL)
		if err != nil {
			logger.WithError(err).WithField("attempt", attempt).Fatal("failed to open database connection")

			if attempt > 3 {
				logger.Fatal("failed to open database connection after 3 attempts")
			}
			time.Sleep(time.Second * 5)
		}
		break
	}

	logger.Info("connected to database")
	time.Sleep(3 * time.Second)

	// Create an ent.Driver from db.
	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(driver))
}

func Migrate(ctx context.Context, logger log.Interface) {
	logger.Info("initiating database schema migration")
	db := ent.FromContext(ctx)
	if db == nil {
		logger.Fatal("failed to get ent client from context")
		return
	}

	if err := db.Schema.Create(
		ctx,
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
	); err != nil {
		logger.WithError(err).Fatal("failed to create schema")
	}
	logger.Info("database schema migration complete")
}
