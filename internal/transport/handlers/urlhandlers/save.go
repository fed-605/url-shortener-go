package urlhandlers

import (
	"errors"
	"log/slog"
	"net/http"

	resp "github.com/fed-605/url-shortener-go/internal/lib/api/response"
	"github.com/fed-605/url-shortener-go/internal/lib/logger/sl"
	"github.com/fed-605/url-shortener-go/internal/lib/random"
	"github.com/fed-605/url-shortener-go/internal/storage"

	handlsers "github.com/fed-605/url-shortener-go/internal/transport/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

const (
	aliasSize = 4
)

/*
Contract
	pattern:/
	method:POST
	info:body
*/

func (app *Application) save(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.urlhandlers.save"

	log := app.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var req handlsers.SaveUrlRequest

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		log.Error("failed to decode request body", sl.Err(err))
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, resp.Error("failed to decode request body"))
		return
	}

	log.Info("request body decoded", slog.Any("request", req))

	//request validation
	if err := app.validator.Struct(req); err != nil {
		log.Error("invalid request", sl.Err(err))
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, resp.Error("invalid request"))
		return
	}

	if req.Alias == "" {
		req.Alias = random.NewRandomString(aliasSize)
	}
	//Add in storage
	if err := app.store.SaveUrl(req.URL, req.Alias); err != nil {
		if errors.Is(err, storage.ErrURLExists) {
			log.Info("url already exists", slog.String("url", req.URL))
			render.Status(r, http.StatusConflict)
			render.JSON(w, r, resp.Error("url already exists"))
			return
		}
		log.Error("failed to save url", sl.Err(err))
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, resp.Error("failed to add url"))
		return

	}

	log.Info("url added in storage")

	//Try to add in cache
	if err := app.cache.SaveUrlMapping(req.URL, req.Alias); err != nil {
		log.Warn("failed to save url in cache", sl.Err(err))
	}
	responseOK(w, r, req.Alias)
}

func responseOK(w http.ResponseWriter, r *http.Request, alias string) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, handlsers.Response{
		Response: resp.OK(),
		Alias:    alias,
	})
}
