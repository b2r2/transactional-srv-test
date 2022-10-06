package config

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// ConnectionConfig ...
type ConnectionConfig struct {
	Driver string
	Host   string
	Port   string
	User   string
	Pass   string
	Name   string
}

// Database ...
type Database struct {
	connectionConfig *ConnectionConfig
	connection       *sqlx.DB
}

// GetDSN ...
func (cc *ConnectionConfig) GetDSN() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		cc.Driver,
		cc.User,
		cc.Pass,
		cc.Host,
		cc.Port,
		cc.Name,
	)
}

// GetConnection ...
func (d *Database) GetConnection() *sqlx.DB {
	return d.connection
}

// NewDB ...
func NewDB() (*Database, error) {
	cfg := ConnectionConfig{
		Driver: GetEnv("DB_DRIVER", "postgres"),
		Host:   GetEnv("DB_HOST", "localhost"),
		Port:   GetEnv("DB_PORT", "5432"),
		User:   GetEnv("DB_USERNAME", "postgres"),
		Pass:   GetEnv("DB_PASSWORD", "sample"),
		Name:   GetEnv("DB_DATABASE", "postgres"),
	}

	conn, err := sqlx.Connect(cfg.Driver, cfg.GetDSN())
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5 * time.Minute)

	d := Database{
		connectionConfig: &cfg,
		connection:       conn,
	}

	return &d, nil
}
