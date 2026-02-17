package urlhandlers

import (
	"log/slog"
	"net/http"

	"github.com/fed-605/url-shortener-go/internal/storage"
	"github.com/fed-605/url-shortener-go/internal/storage/cache"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

func (app *Application) Routes(user, pass string) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID) // unique id for every http request

	router.Use(middleware.Logger)

	router.Use(middleware.Recoverer)

	router.Route("/url", func(r chi.Router) {
		r.Use(middleware.BasicAuth("url-shortener", map[string]string{
			user: pass,
			// can list more users with creds just add a new map string
		}))
		r.Post("/", app.save)

		r.Delete("/{alias}", app.Delete)

		r.Get("/", app.getAllRecords)

	})

	router.Get("/{alias}", app.redirect)

	return router
}
