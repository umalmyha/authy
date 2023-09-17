package router

import "net/http"

type ConfigFunc func(c *Config)

type Config struct {
	errHandler   ErrorHandler
	panicHandler PanicHandler
}

func WithErrorHandler(errHandler ErrorHandler) ConfigFunc {
	return func(c *Config) {
		if errHandler != nil {
			c.errHandler = errHandler
		}
	}
}

func WithPanicHandler(panicHandler PanicHandler) ConfigFunc {
	return func(c *Config) {
		if panicHandler != nil {
			c.panicHandler = panicHandler
		}
	}
}

func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {

}
