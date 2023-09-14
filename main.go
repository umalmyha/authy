package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/umalmyha/kit/bootstrap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger, err := zapLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	mux := chi.NewRouter()
	mux.Post("/login", func(writer http.ResponseWriter, request *http.Request) {

	})

	srv := &http.Server{
		Addr:                         ":8080",
		Handler:                      mux,
		DisableGeneralOptionsHandler: false,
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

func zapLogger() (*zap.SugaredLogger, error) {
	cfg := zap.NewProductionConfig()
	cfg.DisableStacktrace = true
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.CallerKey = "src"
	cfg.InitialFields = map[string]any{"service": "authy"}

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
