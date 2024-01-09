package config

type BotType int

const (
	TgBot BotType = iota
)

type Config struct {
	BotType  BotType
	TgCfg    TgConfig
	DbConfig DbConfig
}

func NewConfig() (*Config, error) {
	// TODO from env
	tgCfg := TgConfig{
		token: "6781544832:AAEHpWffFUg8Jp2dAlboJnzoFStmIvIQrk8",
	}
	botType := TgBot
	baseDb := DbConfig{
		host:     "localhost",
		port:     "5432",
		username: "default",
		password: "secret",
		dbname:   "mood_lifter",
	}

	return &Config{
		TgCfg:    tgCfg,
		BotType:  botType,
		DbConfig: baseDb,
	}, nil
}
