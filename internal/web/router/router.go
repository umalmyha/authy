package router

import (
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/go-chi/chi/v5"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

type Middleware func(next HandlerFunc) HandlerFunc

type Router struct {
	mux *httptreemux.ContextMux
	mw  []Middleware
}

func NewRouter() *Router {
	return &Router{mux: httptreemux.NewContextMux()}
}

func (r *Router) Use(mw ...Middleware) {
	r.mw = append(r.mw, mw...)
}

func (r *Router) Handle(method string, path string, handler HandlerFunc, mw ...Middleware) {
	handler = r.wrap(handler, mw...)
	handler = r.wrap(handler, r.mw...)

	h := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {

		}
	}

	r.mux.Handle(method, path, h)
}

func (r *Router) Group() {
	c := chi.NewRouter()
	c.Group()
}

func (r *Router) wrap(handler HandlerFunc, middleware ...Middleware) HandlerFunc {
	for i := len(middleware) - 1; i >= 0; i-- {
		if mw := middleware[i]; mw != nil {
			handler = mw(handler)
		}
	}
	return handler
}
