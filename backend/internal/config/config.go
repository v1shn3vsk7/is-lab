package config

import "time"

type Config struct {
	HTTPAddr string `env:"HTTP_ADDR" envDefault:":8080"`

	MngDSN           string        `env:"MNG_DSN"`
	MngDBName        string        `env:"MNG_DB_NAME" envDefault:"admin"`
	MngPingInterval  time.Duration `env:"MNG_PING_INTERVAL" envDefault:"10s"`
	MngMinPoolSize   int           `env:"MNG_MIN_POOL_SIZE" envDefault:"400"`
	MngMaxPoolSize   int           `env:"MNG_MAX_POOL_SIZE" envDefault:"500"`
	MngMaxConnecting int           `env:"MNG_MAX_CONNECTING" envDefault:"30"`

	// key for hashing secrets
	SecretKey string `env:"SECRET_KEY" envDefault:""`
}
