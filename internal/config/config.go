package config

import (
	"github.com/jmoiron/sqlx"
)

// Config ...
type Config struct {
	LoggerLevel string
	Database    *Database
}

var config *Config

// Load ...
func Load() error {
	db, err := NewDB()
	if err != nil {
		return err
	}

	cfg := Config{
		LoggerLevel: GetEnv("LOGGER_LEVEL", "info"),
		Database:    db,
	}

	config = &cfg
	return nil
}

// Get ...
func Get() *Config {
	return config
}

// GetDBConnection ...
func (c *Config) GetDBConnection() *sqlx.DB {
	return config.Database.GetConnection()
}

// Close ...
func (c *Config) Close() error {
	return c.Database.GetConnection().Close()
}
