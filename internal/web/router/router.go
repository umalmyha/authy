package router

import (
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

type Middleware func(next Handler) Handler

type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

type PanicHandler = httptreemux.PanicHandler

type Router struct {
	mux  *httptreemux.ContextMux
	root *Group
}

func NewRouter(opts ...ConfigFunc) *Router {
	cfg := Config{
		errHandler:   nil,
		panicHandler: nil,
	}

	mux := httptreemux.NewContextMux()

	return &Router{
		mux: mux,
		root: &Group{
			group:      mux.ContextGroup,
			errHandler: nil,
		},
	}
}

func (r *Router) Use(mw ...Middleware) {
	r.root.Use(mw...)
}

func (r *Router) GET(path string, handler Handler, mw ...Middleware) {
	r.root.GET(path, handler, mw...)
}

func (r *Router) POST(path string, handler Handler, mw ...Middleware) {
	r.root.POST(path, handler, mw...)
}

func (r *Router) PUT(path string, handler Handler, mw ...Middleware) {
	r.root.PUT(path, handler, mw...)
}

func (r *Router) PATCH(path string, handler Handler, mw ...Middleware) {
	r.root.PATCH(path, handler, mw...)
}

func (r *Router) DELETE(path string, handler Handler, mw ...Middleware) {
	r.root.DELETE(path, handler, mw...)
}

func (r *Router) HEAD(path string, handler Handler, mw ...Middleware) {
	r.root.HEAD(path, handler, mw...)
}

func (r *Router) OPTIONS(path string, handler Handler, mw ...Middleware) {
	r.root.OPTIONS(path, handler, mw...)
}

func (r *Router) TRACE(path string, handler Handler, mw ...Middleware) {
	r.root.TRACE(path, handler, mw...)
}

func (r *Router) CONNECT(path string, handler Handler, mw ...Middleware) {
	r.root.CONNECT(path, handler, mw...)
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	r.mux.ServeHTTP(rw, rq)
}
