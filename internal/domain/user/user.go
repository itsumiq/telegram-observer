package user

import "time"

type User struct {
	ID         string    `db:"id"`
	TelegramID int64     `db:"telegram_id"`
	Username   string    `db:"username"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func New(id string, telegramID int64, username string) *User {
	timeNow := time.Now().UTC()
	return &User{
		ID:         id,
		TelegramID: telegramID,
		Username:   username,
		CreatedAt:  timeNow,
		UpdatedAt:  timeNow,
	}
}
