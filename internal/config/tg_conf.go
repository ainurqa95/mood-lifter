package config

type TgConfig struct {
	token string
}

func (t *TgConfig) GetToken() string {
	return t.token
}
