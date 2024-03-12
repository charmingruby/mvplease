package postgres

import (
	"fmt"

	"github.com/charmingruby/mvplease/internal/config"
	"github.com/jmoiron/sqlx"
)

func New(cfg *config.Config) (*sqlx.DB, error) {
	cfg.Logger.Info("Connecting to database...")

	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Name,
		cfg.Database.SSL,
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	cfg.Logger.Info("Connected to database.")

	return db, nil
}
