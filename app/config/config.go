package config

import (
	"time"
	"github.com/vrischmann/envconfig"

	"battleship/app/errors"
)

type Config struct {
	Database struct {
		DSN             string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifetime time.Duration
	}
}

type envConfig struct {
	Dsn             string
	MaxIdleConns    int           `envconfig:"default=1"`
	MaxOpenConns    int           `envconfig:"default=50"`
	ConnMaxLifetime time.Duration `envconfig:"default=0"`
}

func Init() (*Config, error) {
	conf := &envConfig{}

	err := envconfig.InitWithPrefix(conf, "BATTLESHIP")
	if err != nil {
		return nil, errors.NewSystemError(errors.SystemConfigurationError, err)
	}

	appConf := &Config{
		Database: struct {
			DSN             string
			MaxIdleConns    int
			MaxOpenConns    int
			ConnMaxLifetime time.Duration
		}{
			DSN:             conf.Dsn,
			MaxIdleConns:    conf.MaxIdleConns,
			MaxOpenConns:    conf.MaxOpenConns,
			ConnMaxLifetime: conf.ConnMaxLifetime,
		},
	}
	return appConf, nil
}

func (c *Config) GetDSN() string {
	return c.Database.DSN
}

func (c *Config) GetConnectionsMaxIdle() int {
	return c.Database.MaxIdleConns
}

func (c *Config) GetConnectionsMaxOpen() int {
	return c.Database.MaxOpenConns
}

func (c *Config) GetConnectionsMaxLifetime() time.Duration {
	return c.Database.ConnMaxLifetime
}
