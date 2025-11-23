package router

import "net/http"

func methodHandler(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
