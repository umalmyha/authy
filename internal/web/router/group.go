package router

import (
	"net/http"
	"slices"

	"github.com/dimfeld/httptreemux/v5"
)

type Group struct {
	group      *httptreemux.ContextGroup
	errHandler ErrorHandler
	mw         []Middleware
}

func (g *Group) Use(mw ...Middleware) {
	g.mw = append(g.mw, mw...)
}

func (g *Group) Group(path string, fn func(grp *Group)) {
	grp := &Group{
		mw:    slices.Clone(g.mw),
		group: g.group.NewGroup(path),
	}
	fn(grp)
}

func (g *Group) GET(path string, handler Handler, mw ...Middleware) {
	g.handle(http.MethodGet, path, handler, mw...)
}

func (g *Group) POST(path string, handler Handler, mw ...Middleware) {
	g.handle(http.MethodPost, path, handler, mw...)
}

func (g *Group) PUT(path string, handler Handler, mw ...Middleware) {
	g.handle(http.MethodPut, path, handler, mw...)
}

func (g *Group) PATCH(path string, handler Handler, mw ...Middleware) {
	g.handle(http.MethodPatch, path, handler, mw...)
}

func (g *Group) DELETE(path string, handler Handler, mw ...Middleware) {
	g.handle(http.MethodDelete, path, handler, mw...)
}

func (g *Group) HEAD(path string, handler Handler, mw ...Middleware) {
	g.handle(http.MethodHead, path, handler, mw...)
}

func (g *Group) OPTIONS(path string, handler Handler, mw ...Middleware) {
	g.handle(http.MethodOptions, path, handler, mw...)
}

func (g *Group) TRACE(path string, handler Handler, mw ...Middleware) {
	g.handle(http.MethodTrace, path, handler, mw...)
}

func (g *Group) CONNECT(path string, handler Handler, mw ...Middleware) {
	g.handle(http.MethodConnect, path, handler, mw...)
}

func (g *Group) handle(method, path string, handler Handler, mw ...Middleware) {
	handler = g.wrap(handler, mw...)
	handler = g.wrap(handler, g.mw...)

	h := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			g.errHandler(w, r, err)
		}
	}

	g.group.Handle(method, path, h)
}

func (g *Group) wrap(handler Handler, middleware ...Middleware) Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		if mw := middleware[i]; mw != nil {
			handler = mw(handler)
		}
	}
	return handler
}
