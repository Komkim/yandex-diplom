package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/spf13/pflag"
)

const (
	defaultServerAddress  = "127.0.0.1:8080"
	defaultDatabaseDsn    = "postgres://postgres:changeme@localhost:5432/yandex?sslmode=disable"
	defaultAccrualAddress = "127.0.0.1:8081"
)

type HTTP struct {
	Address string `env:"RUN_ADDRESS" mapstructure:"address"`
}
type Server struct {
	DatabaseDSN    string `env:"DATABASE_URI" mapstructure:"databasedsn"`
	AccrualAddress string `env:"ACCRUAL_SYSTEM_ADDRESS" mapstructure:"accrualaddress"`
}

type Config struct {
	HTTP   HTTP
	Server Server
}

func NewConfig() (*Config, error) {
	cfg := new(Config)

	defaultFlag(cfg)

	initFlag(cfg)

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func defaultFlag(cfg *Config) {
	cfg.HTTP.Address = defaultServerAddress

	cfg.Server.DatabaseDSN = defaultDatabaseDsn
	cfg.Server.AccrualAddress = defaultAccrualAddress
}

func initFlag(cfg *Config) {
	pflag.StringVarP(&cfg.HTTP.Address, "server.address", "a", "127.0.0.1:8080", "server address")
	pflag.StringVarP(&cfg.Server.AccrualAddress, "accrual.address", "r", "127.0.0.1:8081", "accrual address")
	pflag.StringVarP(&cfg.Server.DatabaseDSN, "server.databasedsn", "d", "", "connect postgresql")
	pflag.Parse()
}
