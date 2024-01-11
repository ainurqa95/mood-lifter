package config

type TgConfig struct {
	Token string `env:"BOT_TOKEN" envDefault:""`
}

func (t *TgConfig) GetToken() string {
	return t.Token
}
