package config

import (
	"errors"
	"os"
)

const (
	dsnEnvName = "PG_DSN"
)

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

// Create postgres config
func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("PG dsn not found")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

// Get DNS string
func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
