package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/spf13/pflag"
)

const (
	defaultAccrualDatabaseURI  = "postgres://postgres:changeme@localhost:5432/accrual?sslmode=disable"
	defaultAccrualAddressStart = "127.0.0.1:8081"
)

type AccrualConfig struct {
	Address     string `env:"RUN_ADDRESS" mapstructure:"address"`
	DatabaseDSN string `env:"DATABASE_URI" mapstructure:"databasedsn"`
}

func NewAccrualConfig() (*AccrualConfig, error) {
	cfg := new(AccrualConfig)

	defaultAccrualFlag(cfg)

	initAccrualFlag(cfg)

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func defaultAccrualFlag(cfg *AccrualConfig) {
	cfg.Address = defaultAccrualAddressStart
	cfg.DatabaseDSN = defaultAccrualDatabaseURI
}

func initAccrualFlag(cfg *AccrualConfig) {
	pflag.StringVarP(&cfg.Address, "address", "a", "127.0.0.1:8081", "server address")
	pflag.StringVarP(&cfg.DatabaseDSN, "databasedsn", "d", "", "connect postgresql")
	pflag.Parse()
}
