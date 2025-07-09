package messageprocessing

import (
	"fmt"
	"log/slog"
	"telegram-observer/internal/domain/message"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	maxButtonsPerRow = 3
)

type TelegramClient struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramClient(
	token string,
	webhookUrl string,
	logger *slog.Logger,
) (*TelegramClient, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		logger.Error("failed to create telegram bot", "error", err)
		return nil, fmt.Errorf("messageprocessing.NewTelegramClient: %w", err)
	}

	webhookCfg, err := tgbotapi.NewWebhook(webhookUrl)
	if err != nil {
		logger.Error("failed to create telegram bot webhook config", "error", err)
		return nil, fmt.Errorf("messageprocessing.NewTelegramClient: %w", err)
	}

	if _, err := bot.Request(webhookCfg); err != nil {
		logger.Error("failed to create telegram bot webhook", "error", err)
		return nil, fmt.Errorf("messageprocessing.NewTelegramClient: %w", err)
	}

	logger.Info("successful created telegram bot client")

	return &TelegramClient{bot: bot}, nil
}

func (c *TelegramClient) Send(msg *message.Message) error {
	tgMsg := tgbotapi.NewMessage(msg.ChatID, msg.Text)
	tgMsg.ParseMode = "MarkdownV2"

	if len(msg.Buttons) != 0 {
		tgMsg.ReplyMarkup = c.createReplyMarkup(msg.Buttons)
	}

	if _, err := c.bot.Send(tgMsg); err != nil {
		return fmt.Errorf("telegramClient.Send: %w", err)
	}

	return nil

}

func (c *TelegramClient) SendPhotoMessage(msg *message.Message) error {
	photo := tgbotapi.NewPhoto(msg.ChatID, tgbotapi.FilePath(msg.FilePath))
	photo.Caption = msg.Text
	photo.ParseMode = "MarkdownV2"

	if _, err := c.bot.Send(photo); err != nil {
		return fmt.Errorf("telegramClient.SendPhoto: %w", err)
	}

	return nil
}

func (c *TelegramClient) createReplyMarkup(buttons []message.Button) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0, 1)
	currentRow := make([]tgbotapi.InlineKeyboardButton, 0, 1)

	for i, button := range buttons {
		currentRow = append(
			currentRow,
			tgbotapi.NewInlineKeyboardButtonData(button.Text, button.Data),
		)

		if len(currentRow) == maxButtonsPerRow || i == len(buttons)-1 {
			rows = append(rows, currentRow)
			currentRow = nil
		}
	}

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}
