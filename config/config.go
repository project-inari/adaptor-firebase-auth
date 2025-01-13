// Package config provides configuration settings for the server
package config

import (
	"encoding/base64"
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var once sync.Once
var config *Config

// New loads the configuration from the .env file
func New() *Config {
	once.Do(func() {
		e := os.Getenv("APP_ENV_STAGE")
		if e == "" || e == "LOCAL" {
			if err := godotenv.Load(".env.generated"); err != nil {
				slog.Warn("[config.New] unable to load .env.generated file", slog.Any("error", err))
			}
		}

		cfg := &Config{}
		if err := env.Parse(cfg); err != nil {
			log.Panicf("error - [config.New] unable to parse config: %v", err)
		}
		config = cfg
		decodeBase64Field(&config.FirebaseAuthConfig.CredentialsJSON)
	})

	return config
}

// Config represents the configuration of the server
type Config struct {
	AppConfig          AppConfig
	LogConfig          LogConfig
	SentryConfig       SentryConfig
	FirebaseAuthConfig FirebaseAuthConfig
}

// AppConfig represents the configuration of the application
type AppConfig struct {
	Name     string `env:"APP_NAME,notEmpty"`
	Port     string `env:"APP_PORT,notEmpty"`
	EnvStage string `env:"APP_ENV_STAGE,notEmpty"`
}

// LogConfig represents the configuration of the logger
type LogConfig struct {
	Level             string `env:"LOG_LEVEL,notEmpty"`
	MaskSensitiveData bool   `env:"LOG_MASK_SENSITIVE_DATA,notEmpty"`
}

// SentryConfig represents the configuration of Sentry.io
type SentryConfig struct {
	SentryDSN string `env:"SENTRY_DSN"`
}

// FirebaseAuthConfig represents the configuration of Firebase Auth
type FirebaseAuthConfig struct {
	ProjectID       string `env:"FIREBASE_PROJECT_ID,notEmpty"`
	CredentialsJSON string `env:"FIREBASE_CREDENTIALS_JSON,notEmpty"`
}

func decodeBase64Field(fields ...*string) {
	for _, field := range fields {
		if field != nil {
			decoded, err := base64.StdEncoding.DecodeString(*field)
			if err == nil {
				*field = string(decoded)
			}
		}
	}
}
