// user-service: REST API trả hồ sơ admin. Chỉ chấp nhận token có scope users:read.
// Backoffice gọi service này bằng token của CHÍNH backoffice (client_credentials).
package main

import (
	"context"
	"errors"
	"github.com/phillong97254/go-lib/database"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	auth "github.com/phillong97254/go-lib/authkit"
	"github.com/phillong97254/user-service/cmd/rest/routes"
	"github.com/phillong97254/user-service/config"
	v1 "github.com/phillong97254/user-service/internal/handler/rest/v1"
	"github.com/phillong97254/user-service/internal/repository"
	"github.com/phillong97254/user-service/internal/service/profile"
	"github.com/phillong97254/user-service/migration"
)

func main() {
	// Load config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)).With("server", cfg.RestServer.Name, "env", cfg.API.Env)
	slog.SetDefault(logger)

	// Authorize: verify JWT bằng JWKS của auth-service
	verifier := auth.NewVerifier(cfg.OAuth.JWKSURL, cfg.OAuth.Issuer, cfg.OAuth.Audience)

	// Init Database (+ migration)
	db, err := database.New(database.Config{
		URL:                   cfg.DB.URL,
		MaxOpenConnections:    cfg.DB.MaxOpenConnections,
		MaxIdleConnections:    cfg.DB.MaxIdleConnections,
		ConnectionMaxLifetime: cfg.DB.ConnectionMaxLifetime,
		ConnectionMaxIdleTime: cfg.DB.ConnectionMaxIdleTime,
		LogSQL:                cfg.DB.LogSQL,
	})
	if err != nil {
		log.Fatal(err)
	}
	if cfg.DB.AutoMigrate {
		if err := database.Migrate(cfg.DB.URL, migration.FS, migration.Dir); err != nil {
			log.Fatal(err)
		}
	}

	// Repository -> Service -> Handler
	repo := repository.New(db)
	profileSvc := profile.New(logger, repo)
	handler := v1.New(logger, profileSvc)

	// Register Routes
	mux := http.NewServeMux()
	routes.V1Routes(mux, verifier, handler)

	// Start Server
	srv := &http.Server{Addr: cfg.RestServer.Addr, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	go func() {
		logger.Info("user-service (REST) listening", "addr", cfg.RestServer.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.RestServer.ShutdownTimeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("shutdown", "error", err)
	}
	logger.Info("user-service stopped")
}
