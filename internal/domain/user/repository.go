package user

import "context"

// Repository represent contract that user repositories should implement.
type Repository interface {
	// Save saves user and return error
	Save(ctx context.Context, user *User) error

	// GetIDByTelegramID finds user by telegramID and returns id and error.
	GetIDByTelegramID(ctx context.Context, telegramID int64) (string, error)

	// GetTelegramIDByID finds user by id and returns telegramID and error.
	GetTelegramIDByID(ctx context.Context, id string) (int64, error)
}
