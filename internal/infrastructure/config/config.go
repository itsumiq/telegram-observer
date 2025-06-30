package config

import (
	"fmt"
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

type Mode string

const (
	ModeDevelopment Mode = "dev"
	ModeProduction  Mode = "prod"
)

type server struct {
	Addr        string `yaml:"addr"`
	Port        int    `yaml:"port"`
	MaxFileSize int64  `yaml:"max-file-size"`
	FilePath    string `yaml:"file-path"`
}

type telegram struct {
	Token string `yaml:"token" env:"TELEGRAM_TOKEN" env-required:"true"`
}

type sqlite struct {
	DriverName string `yaml:"driver-name"`
	Path       string `yaml:"path"        env:"SQLITE_PATH" env-required:"true"`
}

type Config struct {
	Server   server
	Telegram telegram
	Sqlite   sqlite
}

func New(mode string, log *slog.Logger) (*Config, error) {
	config := &Config{}

	configPath := "configs/config-dev.yaml"
	if Mode(mode) == ModeProduction {
		configPath = "config/config-prod.yaml"
	}

	if err := cleanenv.ReadConfig(configPath, config); err != nil {
		log.Error("failed to read config", "error", err, "config_path", configPath)
		return nil, fmt.Errorf("config.New: %w", err)
	}

	log.Info("config successful readed")

	return config, nil
}
