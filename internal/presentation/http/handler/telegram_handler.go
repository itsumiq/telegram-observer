package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"telegram-observer/internal/application"
	"telegram-observer/internal/infrastructure/storage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram interface {
	HandlePostWebhook(w http.ResponseWriter, r *http.Request)
}

type telegramHandler struct {
	userService    application.UserService
	messageService application.MessageService
}

func NewTelegramHandler(
	userService application.UserService,
	messageService application.MessageService,
) *telegramHandler {
	return &telegramHandler{userService: userService, messageService: messageService}
}

type updateRequest struct {
	tgbotapi.Update
}

func (h *telegramHandler) HandlePostWebhook(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value("logger").(*slog.Logger)

	updateReq := &updateRequest{}
	if err := json.NewDecoder(r.Body).Decode(updateReq); err != nil {
		log.Error("failed to decode telegram request", slog.String("error", err.Error()))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if updateReq.Message.IsCommand() {
		if err := h.handleCommand(updateReq, log); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (h *telegramHandler) handleCommand(updateReq *updateRequest, log *slog.Logger) error {
	switch updateReq.Message.Command() {
	case "/start":
		_, err := h.userService.Create(updateReq.Message.From.ID, updateReq.Message.From.UserName)
		if err != nil {
			if errors.Is(err, storage.ErrDublicateUser) {
				log.Warn("dublicate user", slog.String("error", err.Error()))
				if err := h.messageService.SendStartResponse(updateReq.Message.From.ID); err != nil {
					log.Error(
						"failed to send response to start command",
						slog.String("error", err.Error()),
					)
					return err
				}

				return nil
			}
			log.Error("failed to create user", slog.String("error", err.Error()))
			return err
		}

		if err := h.messageService.SendStartResponse(updateReq.Message.From.ID); err != nil {
			log.Error("failed to send response to start command", slog.String("error", err.Error()))
			return err
		}

		return nil
	case "/profile":
		userID, err := h.userService.GetIDByTelegramID(updateReq.Message.From.ID)
		if err != nil {
			if errors.Is(err, storage.ErrUserNotFound) {
				log.Warn("user not found", slog.String("error", err.Error()))
				if err := h.messageService.SendUserNotFoundResponse(updateReq.Message.From.ID); err != nil {
					log.Error(
						"failed to send response to user not found",
						slog.String("error", err.Error()),
					)
					return err
				}
				return nil
			}
			log.Error("failed to get userID", slog.String("error", err.Error()))
			return err
		}

		if err := h.messageService.SendProfileResponse(updateReq.Message.From.ID, userID); err != nil {
			log.Error(
				"failed to send response to profile command",
				slog.String("error", err.Error()),
			)
			return err
		}

		return nil
	default:
		if err := h.messageService.SendUnknownCommandResponse(updateReq.Message.From.ID); err != nil {
			log.Error(
				"failed to send response to unknown command",
				slog.String("error", err.Error()),
			)
			return err
		}

		return nil
	}
}
