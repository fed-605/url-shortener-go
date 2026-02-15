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
)

/*
Contract
	pattern:/{alias}
	method:DELETE
	info:query
*/

func (app *Application) Delete(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.urlhandlers.delete"

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
	// deleting from storage
	if err := app.store.DeleteUrl(alias); err != nil {
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("url not found", "alias", alias)
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, resp.Error("not found"))
			return
		}
		log.Error("failed to delete url from storage", sl.Err(err))
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, resp.Error("internal error"))
		return
	}
	log.Info("url deleted from storage")
	//deleting from cache
	if err := app.cache.DeleteUrl(alias); err != nil {
		log.Warn("failed to delete url from cache", sl.Err(err))
	}

	w.WriteHeader(http.StatusNoContent)

}
