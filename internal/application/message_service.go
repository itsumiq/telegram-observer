package application

import (
	"fmt"
	"telegram-observer/internal/domain/message"
)

type MessageService interface {
	SendStartResponse(telegramID int64) error
	SendProfileResponse(telegramID int64, userID string) error
	SendUserNotFoundResponse(telegramID int64) error
	SendUnknownCommandResponse(telegramID int64) error
}

type messageService struct {
	messageProcessing message.MessageProcessing
}

func NewMessageService(messageProcessing message.MessageProcessing) *messageService {
	return &messageService{messageProcessing: messageProcessing}
}

func (s *messageService) SendStartResponse(telegramID int64) error {
	msgText := "Привет! С помощью этого бота ты сможешь получать уведомления с камер видеонаблюдения при обнаружении человека\n" +
		"Команды:\n" +
		"/start - Показывает приветствие и основные команды\n" +
		"/profile - Показывает информацию о твоем профиле"

	msg := message.New(telegramID, msgText)
	if err := s.messageProcessing.Send(msg); err != nil {
		return fmt.Errorf("messageService.SendStartResponse: %w", err)
	}

	return nil
}

func (s *messageService) SendProfileResponse(telegramID int64, userID string) error {
	msgText := fmt.Sprintf("Твой id: %s", userID)
	msg := message.New(telegramID, msgText)
	if err := s.messageProcessing.Send(msg); err != nil {
		return fmt.Errorf("messageService.SendProfileResponse: %w", err)
	}

	return nil
}

func (s *messageService) SendUserNotFoundResponse(telegramID int64) error {
	msgText := "Чтобы использовать бота необходимо зарегистрироваться: /start"
	msg := message.New(telegramID, msgText)
	if err := s.messageProcessing.Send(msg); err != nil {
		return fmt.Errorf("messageService.SendUserNotFoundResponse: %w", err)
	}

	return nil
}

func (s *messageService) SendUnknownCommandResponse(telegramID int64) error {
	msgText := "Неизвестная команда"
	msg := message.New(telegramID, msgText)
	if err := s.messageProcessing.Send(msg); err != nil {
		return fmt.Errorf("messageService.SendUnknownCommandResponse: %w", err)
	}

	return nil
}
