package main

import (
	"log/slog"
	"os"

	"github.com/fed-605/url-shortener-go/internal/config"
	"github.com/fed-605/url-shortener-go/internal/http_server/server"
	"github.com/fed-605/url-shortener-go/internal/lib/logger/sl"
	"github.com/fed-605/url-shortener-go/internal/storage/cache/redis"
	"github.com/fed-605/url-shortener-go/internal/storage/postgres"
	"github.com/fed-605/url-shortener-go/internal/transport/handlers/urlhandlers"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.LoadAllConfig()

	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	store, err := postgres.NewPostgres(cfg.PostgreDSN())
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	cacheStore, err := redis.NewRedisService(cfg.CacheStore.Addr, cfg.CacheStore.Timeout, cfg.CacheStore.DialTimeout)
	if err != nil {
		log.Error("failed to init cache service", sl.Err(err))
		os.Exit(1)
	}

	application := urlhandlers.NewApp(log, store, cacheStore)

	srv := server.New(cfg.Server.Address, application.Routes(cfg.Server.User, cfg.Server.Password), cfg.Server.Timeout, cfg.Server.Timeout, cfg.Server.IdleTimeout)
	if err := srv.Run(); err != nil {
		log.Error("failed to start server", sl.Err(err))
	}

	log.Error("server stopped")
}

// A new internal app logger
// logger depends on env
// because of different kinds of logging
// depend on environment(local->text,prod and dev -> json)
// dev -> debug level of logs
// prod -> > info level of logs
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	}
	return log
}
