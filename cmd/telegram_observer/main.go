package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"telegram-observer/internal/app"
	"telegram-observer/internal/infrastructure/config"
	"telegram-observer/internal/infrastructure/logger"
)

func main() {
	mode := os.Getenv("MODE")

	logger := logger.New(logger.LevelInfo)

	cfg, err := config.New(mode, logger)
	if err != nil {
		log.Fatal(err)
	}

	app, err := app.New(cfg, logger)
	if err != nil {
		log.Fatal(err)
	}

	go app.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	if err := app.Shutdown(); err != nil {
		log.Fatal(err)
	}

}
