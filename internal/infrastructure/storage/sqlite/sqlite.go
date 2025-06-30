package sqlite

import (
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Connect(driverName string, dbPath string, log *slog.Logger) (*sqlx.DB, error) {
	conn, err := sqlx.Connect(driverName, dbPath)
	if err != nil {
		log.Error("failed database connection")
		return nil, fmt.Errorf("sqlite.Connect: %w", err)
	}

	log.Info("successful connection to database")
	return conn, nil
}
