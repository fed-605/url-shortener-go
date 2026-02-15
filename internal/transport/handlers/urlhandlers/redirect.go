package urlhandlers

import (
	"errors"
	"log/slog"
	"net/http"

	resp "github.com/fed-605/url-shortener-go/internal/lib/api/response"
	"github.com/fed-605/url-shortener-go/internal/lib/logger/sl"
	"github.com/fed-605/url-shortener-go/internal/storage"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/redis/go-redis/v9"
)

/*
Contract

	pattern:/{alias}
	method:GET
	info:query
*/
func (app *Application) redirect(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.urlhandlers.redirect"

	log := app.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	alias := chi.URLParam(r, "alias")

	log.Info("alias was retrieved from the params", slog.String("alias", alias))

	if err := app.validator.Var(alias, "required"); err != nil {
		log.Error("alias is empty", sl.Err(err))
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, resp.Error("invalid request"))
		return
	}

	resUrl, err := app.cache.RetrieveUrl(alias)
	if err != nil {
		if err != redis.Nil {
			log.Warn("failed to get url from redis(warning)", sl.Err(err))
		}

		storeUrl, err := app.store.GetUrl(alias)
		if err != nil {
			if errors.Is(err, storage.ErrURLNotFound) {
				log.Info("url not found", slog.String("alias", alias))
				render.Status(r, http.StatusNotFound)
				render.JSON(w, r, resp.Error("not found"))
				return
			}
			log.Error("failed to get url", sl.Err(err))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("internal error"))
			return
		}
		log.Info("got url from the storage", slog.String("url", storeUrl))

		if err := app.cache.SaveUrlMapping(storeUrl, alias); err != nil {
			log.Warn("failed to save url in cache", sl.Err(err))
		}

		// redirect to found url
		http.Redirect(w, r, storeUrl, http.StatusFound)
		return
	}

	log.Info("got url from the cache", slog.String("url", resUrl))

	http.Redirect(w, r, resUrl, http.StatusFound)
}
