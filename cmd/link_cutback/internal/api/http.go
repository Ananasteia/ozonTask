package api

import (
	"AVITOtask/cmd/link_cutback/internal/app"
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

type application interface {
	HandlePost(ctx context.Context, ll string) (string, error)
	HandleGet(ctx context.Context, l app.Link) (*app.Link, error)
}
type api struct {
	app application
}

func New(a application) http.Handler {
	api := api{
		app: a,
	}
	r := mux.NewRouter()
	r.HandleFunc("/", api.giveShortLink).Methods(http.MethodPost)
	r.HandleFunc("/{link:[A-Za-z0-9_]+}", api.openLongLink).Methods(http.MethodGet)
	return r
}
