package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"telegram-observer/internal/domain/user"
	"telegram-observer/internal/infrastructure/storage"

	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
)

type UserRepository struct {
	db sqlx.DB
}

func NewUserRepository(db sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(ctx context.Context, user *user.User) error {
	query := "INSERT INTO users (telegram_id, username, created_at, updated_at)" +
		"VALUES ($1, $2, $3, $4)"

	if err := r.db.GetContext(ctx, &user.ID, query, user.TelegramID, user.Username, user.CreatedAt, user.UpdatedAt); err != nil {
		var sqlErr sqlite3.Error
		if errors.As(err, &sqlErr) {
			return fmt.Errorf("userRepository.Save: %w", storage.ErrDublicateUser)
		}
		return fmt.Errorf("userRepository.Save: %w", err)
	}

	return nil
}

func (r *UserRepository) GetIDByTelegramID(ctx context.Context, telegramID int64) (string, error) {
	query := "SELECT id FROM users WHERE telegram_id = $1"
	userID := ""

	if err := r.db.GetContext(ctx, &userID, query, telegramID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userID, fmt.Errorf("userRepository.GetIDByTelegramID: %w", storage.ErrUserNotFound)
		}
		return userID, fmt.Errorf("userRepository.GetIDByTelegramID: %w", err)
	}

	return userID, nil
}

func (r *UserRepository) GetTelegramIDByID(ctx context.Context, id string) (int64, error) {
	query := "SELECT telegram_id FROM users WHERE id = $1"
	var telegramID int64

	if err := r.db.GetContext(ctx, &telegramID, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return telegramID, fmt.Errorf("userRepository.GetTelegramIDByID: %w", storage.ErrUserNotFound)
		}
		return telegramID, fmt.Errorf("userRepository.GetTelegramIDByID: %w", err)
	}

	return telegramID, nil
}
