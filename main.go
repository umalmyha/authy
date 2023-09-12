package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/umalmyha/kit/bootstrap"
)

func main() {

	mux := chi.NewRouter()

	srv := &http.Server{
		Addr:                         ":8080",
		Handler:                      router,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  0,
		ReadHeaderTimeout:            0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
	}

	start := func() error {
		return srv.ListenAndServe()
	}

	stop := func(ctx context.Context) error {
		return srv.Shutdown(ctx)
	}

	orc := bootstrap.New().Service(bootstrap.ToService(start, stop))
	if err := orc.Serve(); err != nil {

	}
}
