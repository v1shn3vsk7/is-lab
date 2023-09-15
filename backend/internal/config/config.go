package config

import "time"

type Config struct {
	HTTPAddr string `env:":8080"`

	MngDSN          string        `env:"MNG_DSN"`
	MngDBName       string        `env:"MNG_DB_NAME" envDefault:"admin"`
	MngPingInterval time.Duration `env:"MNG_PING_INTERVAL" envDefault:"1s"`
}
