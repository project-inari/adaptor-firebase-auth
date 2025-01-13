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
	"github.com/project-inari/adaptor-firebase-auth/middleware"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	// Initiaize config
	cfg := config.New()

	// Initialize logger
	var logger *slog.Logger
	if cfg.LogConfig.Level == "DEBUG" {
		logger = slog.New(slogor.NewHandler(os.Stdout, slogor.SetTimeFormat(time.Stamp), slogor.ShowSource()))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
				if cfg.LogConfig.MaskSensitiveData {
					value := middleware.MaskSensitiveData(a.Key, a.Value.Any())
					return slog.Any(a.Key, value)
				}

				return a
			},
		}))
	}
	slog.SetDefault(logger)

	// Initialize dependency injection
	di.New(cfg)
}
