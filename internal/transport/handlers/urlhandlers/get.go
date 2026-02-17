package urlhandlers

import (
	"log/slog"
	"net/http"

	resp "github.com/fed-605/url-shortener-go/internal/lib/api/response"
	"github.com/fed-605/url-shortener-go/internal/lib/logger/sl"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

/*
Contract
	pattern:/url/
	method:GET
*/

func (app *Application) getAllRecords(w http.ResponseWriter, r *http.Request) {
	const op = "handlers.urlhandlers.getAllRecords"

	log := app.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	recs, err := app.store.GetAllRecords()
	if err != nil {
		log.Error("failed to retrieve all urls from storage", sl.Err(err))
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, resp.Error("internal error"))
		return
	}
	log.Info("urls retrieved from storage")

	render.Status(r, http.StatusOK)

	render.JSON(w, r, recs)

}
