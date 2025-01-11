// Package main is the main entry point for adaptor-firebase-auth service
package main

import (
	"log/slog"
	"os"
	"runtime"
	"time"
	_ "time/tzdata"

	"gitlab.com/greyxor/slogor"

	"github.com/project-inari/adaptor-firebase-auth/config"
	"github.com/project-inari/adaptor-firebase-auth/di"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	// Initialize logger
	env := os.Getenv("APP_ENV_STAGE")
	var logger *slog.Logger
	if env == "" || env == "LOCAL" {
		logger = slog.New(slogor.NewHandler(os.Stdout, slogor.SetTimeFormat(time.Stamp), slogor.ShowSource()))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		}))
	}
	slog.SetDefault(logger)

	// Initiaize config
	cfg := config.New(env)

	// Initialize dependency injection
	di.New(cfg)
}
