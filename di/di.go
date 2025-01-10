// Package di provides dependency injection for the server
package di

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"

	"github.com/project-inari/adaptor-firebase-auth/config"
	"github.com/project-inari/adaptor-firebase-auth/handler"
	"github.com/project-inari/adaptor-firebase-auth/repository"
	"github.com/project-inari/adaptor-firebase-auth/service"
)

// New injects the dependencies for the server
func New(c *config.Config) {
	ctx := context.Background()

	// Sentry initialization
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              c.SentryConfig.SentryDSN,
		Debug:            true,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		slog.Error("error - [di.New] sentry initialization failed", slog.Any("error", err))
	}

	// Echo server initialization
	e := echo.New()
	setupServer(ctx, e, c)

	// Firebase Auth initialization
	firebaseAuthClient, err := setupFirebaseAuth(ctx, c.FirebaseAuthConfig)
	if err != nil {
		log.Panicf("error - [di.New] unable to setup firebase auth: %v", err)
	}

	// Repository initialization
	firebaseAuthRepo := repository.NewFirebaseAuthRepository(repository.FirebaseAuthRepositoryDependencies{
		Client: firebaseAuthClient,
	})

	// Service initialization
	service := service.New(service.Dependencies{
		FirebaseAuthRepository: firebaseAuthRepo,
	})

	// Handler initialization
	handler.New(e, handler.Dependencies{
		Service: service,
	})

	// HTTP Listening
	if err := e.Start(":" + c.AppConfig.Port); err != nil && err != http.ErrServerClosed {
		log.Panicf("error - [main.New] unable to start server: %v", err)
	}
}
