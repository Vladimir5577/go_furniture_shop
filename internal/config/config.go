package config

import (
	"database/sql"
	"fmt"

	"github.com/goloop/env"
)

type EnvConfig struct {
	DBHost     string `env:"DB_HOST" def:"localhost"`
	DbPort     int    `env:"DB_PORT" def:"5432"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME"`
}

func NewEnvConfig() (*EnvConfig, error) {
	// Load .env file.
	if err := env.Load(".env"); err != nil {
		return nil, err
	}

	// Parse environment into struct.
	var envConfig EnvConfig
	if err := env.Unmarshal("", &envConfig); err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", envConfig)
	fmt.Println(envConfig.DBHost)
	return &envConfig, nil
}

func ConnectDB(conf *EnvConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost,
		conf.DbPort,
		conf.DbUser,
		conf.DbPassword,
		conf.DbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
