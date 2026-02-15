package urlhandlers

import (
	"log/slog"
	"net/http"

	"github.com/fed-605/url-shortener-go/internal/storage"
	"github.com/fed-605/url-shortener-go/internal/storage/cache"
	"github.com/go-playground/validator/v10"
)

type Application struct {
	log       *slog.Logger
	store     storage.Storage
	cache     cache.Cache
	validator *validator.Validate
}

func NewApp(logger *slog.Logger, store storage.Storage, cache cache.Cache) *Application {
	return &Application{
		log:       logger,
		store:     store,
		cache:     cache,
		validator: validator.New(),
	}
}

func (app *Application) Routes() http.Handler {
	return nil
}
