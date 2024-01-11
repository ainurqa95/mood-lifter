package config

import "fmt"

type DbConfig struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	Username string `env:"DB_USERNAME" envDefault:"default"`
	Password string `env:"DB_PASSWORD" envDefault:"secret"`
	Dbname   string `env:"DB_NAME" envDefault:"mood_lifter"`
}

func (d *DbConfig) GetSource() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Dbname,
	)
}
