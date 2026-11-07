// user-service: REST API trả hồ sơ admin. Chỉ chấp nhận token có scope users:read.
// Backoffice gọi service này bằng token của CHÍNH backoffice (client_credentials).
package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/phil/user-service/cmd/rest/routes"
	"github.com/phil/user-service/config"
	v1 "github.com/phil/user-service/internal/handler/rest/v1"
	"github.com/phil/user-service/internal/repository"
	"github.com/phil/user-service/internal/service/profile"
	auth "github.com/phil/user-service/pkg/authkit"
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

	// Repository -> Service -> Handler
	repo := repository.New()
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
