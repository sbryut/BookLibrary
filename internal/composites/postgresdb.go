package composites

import (
	"context"
	"database/sql"
	"library/internal/config"
	"library/pkg/client/postgres"
)

type PostgresDBComposite interface {
	DB() *sql.DB
}

type postgresDB struct {
	db *sql.DB
}

func NewPostgresDBComposite(ctx context.Context, cfg config.Config) (*postgresDB, error) {
	client, err := postgres.NewClient(ctx, cfg.DB)
	if err != nil {
		return nil, err
	}
	return &postgresDB{db: client}, nil
}

func (c postgresDB) DB() *sql.DB {
	return c.db
}
