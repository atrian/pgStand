package config

import (
	"flag"

	"github.com/caarlos0/env/v6"

	"pgStand/pkg/logger"
)

const DefaultBuffSize = 1000

type Config struct {
	dsn                string `env:"DATABASE_DSN"`
	connectionType     string `env:"DB_CONNECTION_TYPE"`
	importFile         string `env:"DATASET_FILE"`
	entitiesBufferSize int    `env:"ENTITIES_BUFFER_SIZE"`
	logger             logger.Logger
}

func New() *Config {
	config := Config{}

	return &config
}

func (c *Config) loadFlags() {
	c.logger.Info("Read flags and defaults")

	dsn := flag.String("d", "", "DSN for SQL server")
	connectionType := flag.String("ct", "pool", "Single or Pool DB connection type")
	importFile := flag.String("f", "pool", "Single or Pool DB connection type")
	entitiesBufferSize := flag.Int("bs", DefaultBuffSize, "Flush temp buffer to database every N entities")

	flag.Parse()

	c.dsn = *dsn
	c.connectionType = *connectionType
	c.importFile = *importFile
	c.entitiesBufferSize = *entitiesBufferSize
}

func (c *Config) loadEnvParams() {
	c.logger.Info("Read env params")

	err := env.Parse(c)
	if err != nil {
		c.logger.Error("Can't parse Env params", err)
	}
}
