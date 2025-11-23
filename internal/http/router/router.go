// Package router
package router

import "net/http"

type Router struct {
	mux *http.ServeMux
}

func New() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) Handler() http.Handler {
	return r.mux
}
