package main

/*
*************************************************************************

Author: Babulal Tamang
Purpose: Auth Server
Model Name:
Date: 23rd Jan 2025
Additional Notes:

****************************************************************************
*/

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/authserver/internal/adapter/auth/paseto"
	"github.com/sugaml/authserver/internal/adapter/config"
	http "github.com/sugaml/authserver/internal/adapter/handler"
	"github.com/sugaml/authserver/internal/adapter/storage/postgres"
	"github.com/sugaml/authserver/internal/adapter/storage/postgres/migrations"
	"github.com/sugaml/authserver/internal/adapter/storage/postgres/repository"
	"github.com/sugaml/authserver/internal/core/service"
)

// @title						Auth-Server API
// @version						1.0
// @description					This is a simple RESTful Service API written in Go using Gin web framework
// @securityDefinitions.apikey 	BearerAuth
// @in 							Header
// @name 						Authorization
func main() {
	// Load environment variables
	config, err := config.New()
	if err != nil {
		logrus.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	logrus.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Init database
	ctx := context.Background()
	db, err := postgres.New(ctx, config.DB)
	if err != nil {
		logrus.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	logrus.Info("Successfully connected to the database", "db", config.DB.Connection)

	// // Migrate database
	err = migrations.Migrate(db)
	if err != nil {
		logrus.Error("Error migrating database", "error", err)
		os.Exit(1)
	}

	logrus.Info("Successfully migrated the database")

	logrus.Info("Successfully connected to the cache server")

	// Init token service
	token, err := paseto.New(config.Token)
	if err != nil {
		logrus.Error("Error initializing token service", "error", err)
		os.Exit(1)
	}
	// Init data layer
	repo := repository.NewRepository(db)

	//oauth server
	srv := service.GetOauthServer(repo)

	// Init service
	svc := service.NewService(repo)
	// Init handler
	handler := http.NewHandler(config.HTTP, svc, token, srv)

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	logrus.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = handler.Serve(listenAddr)
	if err != nil {
		logrus.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
