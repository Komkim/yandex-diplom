package config

import "github.com/caarlos0/env/v6"

type HTTP struct {
	Address string `env:"ADDRESS"`
}
type Server struct {
	PG_DSN string `env:"DSN"`
}

type Config struct {
	HTTP   HTTP
	Server Server
}

func NewConfig() (*Config, error) {
	cfg := new(Config)

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
