package database

import (
	"Cart_Api_New/config"
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sipki-tech/database/migrations"
)

type Repo struct {
	sql *sqlx.DB
}

func New(ctx context.Context, cfg config.DBConfig) (*sqlx.DB, error) {
	migrates, err := migrations.Parse(cfg.Migrates)
	if err != nil {
		return nil, err
	}

	err = migrations.Run(ctx, cfg.Driver, &cfg.Postgres, migrations.Up, migrates)
	if err != nil {
		return nil, err
	}

	dsn, err := cfg.Postgres.DSN()
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Open(cfg.Driver, dsn)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil

}
