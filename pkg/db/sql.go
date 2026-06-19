package db

import (
	"context"
	"fmt"

	"github.com/TheDigitalMadness/notifications-service-go/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func buildPostgresString(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DB,
	)
}

func NewPostgres(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool

	var err error
	for i := 0; i < 5; i++ {
		p, pgErr := pgxpool.New(ctx, buildPostgresString(cfg))
		if pgErr != nil {
			err = pgErr
			continue
		}

		if err := p.Ping(ctx); err != nil {
			p.Close()
			continue
		}

		pool = p
		err = nil
	}

	if err != nil {
		return nil, fmt.Errorf("Postgres error: %w", err)
	}

	return pool, nil
}
