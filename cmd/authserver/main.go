package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/sugaml/authserver/docs"
	"github.com/sugaml/authserver/internal/adapter/auth/paseto"
	"github.com/sugaml/authserver/internal/adapter/config"
	http "github.com/sugaml/authserver/internal/adapter/handler"
	"github.com/sugaml/authserver/internal/adapter/logger"
	"github.com/sugaml/authserver/internal/adapter/storage/postgres"
	"github.com/sugaml/authserver/internal/adapter/storage/postgres/migrations"
	"github.com/sugaml/authserver/internal/adapter/storage/postgres/repository"
	"github.com/sugaml/authserver/internal/core/service"
)

// @title					AuthServer API
// @version					1.0
// @description				This is a simple RESTful PaAni Service API written in Go using Gin web framework, PostgreSQL database, and Redis cache.
//
// @contact.name			Babulal Tamang
// @contact.url				https://github.com/hamropani/pani-backend
// @contact.email			hamropani.office@hamropani.com
//
// @host					localhost:8080
// @BasePath				/v1
// @schemes					http https
//
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and the access token.
func main() {
	// Load environment variables
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	// Set logger
	logger.Set(config.App)

	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Init database
	ctx := context.Background()
	db, err := postgres.New(ctx, config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Successfully connected to the database", "db", config.DB.Connection)

	// // Migrate database
	err = migrations.Migrate(db)
	if err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully migrated the database")

	slog.Info("Successfully connected to the cache server")

	// Init token service
	token, err := paseto.New(config.Token)
	if err != nil {
		slog.Error("Error initializing token service", "error", err)
		os.Exit(1)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	// Init handler
	handler := http.NewHandler(config.HTTP, svc, token)

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = handler.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
