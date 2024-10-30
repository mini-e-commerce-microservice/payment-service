package handler

import (
	"fmt"
	whttp "github.com/SyaibanAhmadRamadhan/http-wrapper"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type handler struct {
	r        *chi.Mux
	httpOtel *whttp.Opentelemetry
}

func NewHandler(r *chi.Mux) {

	h := &handler{
		r: r,
		httpOtel: whttp.NewOtel(
			whttp.WithRecoverMode(true),
			whttp.WithPropagator(),
			whttp.WithValidator(nil, nil),
		),
	}
	h.route()
}

func (h *handler) route() {
	h.r.Post("/v1/callback-midtrans", h.httpOtel.Trace(
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Println("callback")
		}, whttp.WithLogRequestBody(false), whttp.WithLogResponseBody(false),
	))
}
