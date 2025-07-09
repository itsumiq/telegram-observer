package sqlite

import (
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Connect(
	driverName string,
	dbPath string,
	migrationPath string,
	log *slog.Logger,
) (*sqlx.DB, error) {
	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationPath),
		fmt.Sprintf("%s://%s", driverName, dbPath),
	)
	if err != nil {
		log.Error("failed to create migrations")
		return nil, fmt.Errorf("sqlite.Connect: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error("failed to apply migrations")
		return nil, fmt.Errorf("sqlite.Connect: %w", err)
	}

	log.Info("all migrations successful applied")

	conn, err := sqlx.Connect(driverName, dbPath)
	if err != nil {
		log.Error("failed database connection")
		return nil, fmt.Errorf("sqlite.Connect: %w", err)
	}

	log.Info("successful connection to database")
	return conn, nil
}
