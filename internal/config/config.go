package config

import (
	"github.com/joho/godotenv"
)

// Load
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
