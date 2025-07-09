package app

import (
	"fmt"
	"log/slog"
	httpapp "telegram-observer/internal/app/http"
	"telegram-observer/internal/application"
	"telegram-observer/internal/infrastructure/config"
	messageprocessing "telegram-observer/internal/infrastructure/message_processing"
	"telegram-observer/internal/infrastructure/storage/sqlite"
	uuiderclient "telegram-observer/internal/infrastructure/uuider"
	"telegram-observer/internal/presentation/http/handler"
)

type App struct {
	httpApp *httpapp.App
	log     *slog.Logger
}

func New(cfg *config.Config, log *slog.Logger) (*App, error) {
	tgClient, err := messageprocessing.NewTelegramClient(cfg.Telegram.Token, cfg.Telegram.Url, log)
	if err != nil {
		return nil, fmt.Errorf("app.New: %w", err)
	}

	db, err := sqlite.Connect(cfg.Sqlite.DriverName, cfg.Sqlite.Path, cfg.Sqlite.MigrationPath, log)
	if err != nil {
		return nil, fmt.Errorf("app.New: %w", err)
	}

	userRepo := sqlite.NewUserRepository(db)
	uuider := uuiderclient.New()

	userService := application.NewUserService(userRepo, uuider)
	messageService := application.NewMessageService(tgClient)
	fileService := application.NewFileService(cfg.Server.FilePath, cfg.Server.MaxFileSize, log)

	tgHandler := handler.NewTelegramHandler(userService, messageService)
	userHandler := handler.NewUserHandler(
		userService,
		fileService,
		messageService,
		cfg.Server.MaxFileSize,
	)

	httpApp := httpapp.New(cfg.Server.Addr, cfg.Server.Port, tgHandler, userHandler, log)

	return &App{httpApp: httpApp, log: log}, nil
}

func (a *App) Run() {
	a.httpApp.Run()
}

func (a *App) Shutdown() error {
	if err := a.httpApp.Shutdown(); err != nil {
		return fmt.Errorf("app.Shutdown: %w", err)
	}
	return nil
}
