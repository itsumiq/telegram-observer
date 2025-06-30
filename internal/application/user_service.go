package application

import (
	"context"
	"fmt"
	"telegram-observer/internal/domain/user"
)

type UserService interface {
	Create(telegramID int64, username string) (string, error)
	GetTelegramIDByID(id string) (int64, error)
	GetIDByTelegramID(telegramID int64) (string, error)
}

type userService struct {
	userRepo user.Repository
	uuider   user.UUIDer
}

func NewUserService(userRepo user.Repository, uuider user.UUIDer) *userService {
	return &userService{userRepo: userRepo, uuider: uuider}
}

func (s *userService) Create(telegramID int64, username string) (string, error) {
	user := user.New(s.uuider.Create(), telegramID, username)
	if err := s.userRepo.Save(context.TODO(), user); err != nil {
		return "", fmt.Errorf("userService.Create: %w", err)
	}

	return user.ID, nil
}

func (s *userService) GetTelegramIDByID(id string) (int64, error) {
	telegramID, err := s.userRepo.GetTelegramIDByID(context.TODO(), id)
	if err != nil {
		return telegramID, fmt.Errorf("userService.GetTelegramIDByID: %w", err)
	}

	return telegramID, nil
}

func (s *userService) GetIDByTelegramID(telegramID int64) (string, error) {
	id, err := s.userRepo.GetIDByTelegramID(context.TODO(), telegramID)
	if err != nil {
		return id, fmt.Errorf("userService.GetIDByTelegramID: %w", err)
	}

	return id, nil
}
