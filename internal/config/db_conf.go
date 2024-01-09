package config

import "fmt"

type DbConfig struct {
	host     string
	port     string
	username string
	password string
	dbname   string
}

func (d *DbConfig) GetSource() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?",
		d.username,
		d.password,
		d.host,
		d.port,
		d.dbname,
	)
}
