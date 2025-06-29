package user

import "context"

type Repository interface {
	Save(ctx context.Context, user *User) error
	GetIDByTelegramID(ctx context.Context, telegramID int64) (string, error)
	GetTelegramIDByID(ctx context.Context, id string) (int64, error)
}
