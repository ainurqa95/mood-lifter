package config

type BotType int

const (
	TgBot BotType = iota
)

type Config struct {
	BotType BotType
	TgCfg   TgConfig
}

func NewConfig() Config {
	// TODO from env
	tgCfg := TgConfig{
		token: "6781544832:AAEHpWffFUg8Jp2dAlboJnzoFStmIvIQrk8",
	}
	botType := TgBot

	return Config{
		TgCfg:   tgCfg,
		BotType: botType,
	}
}
