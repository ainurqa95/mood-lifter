package config

import (
	"github.com/caarlos0/env/v8"
)

type BotType int

const (
	TgBot BotType = iota
)

type Config struct {
	BotType      BotType `env:"BOT_TYPE" envDefault:"0"`
	TgCfg        TgConfig
	DbConfig     DbConfig
	CronSchedule string `env:"SCHEDULE" envDefault:"0 8-18/2 * * *"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
