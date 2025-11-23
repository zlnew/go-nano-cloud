package router

import "net/http"

func (r *Router) Get(path string, h http.Handler) {
	r.mux.Handle(path, methodHandler(http.MethodGet, h))
}

func (r *Router) Post(path string, h http.Handler) {
	r.mux.Handle(path, methodHandler(http.MethodPost, h))
}

func (r *Router) Put(path string, h http.Handler) {
	r.mux.Handle(path, methodHandler(http.MethodPut, h))
}

func (r *Router) Patch(path string, h http.Handler) {
	r.mux.Handle(path, methodHandler(http.MethodPatch, h))
}

func (r *Router) Delete(path string, h http.Handler) {
	r.mux.Handle(path, methodHandler(http.MethodDelete, h))
}
