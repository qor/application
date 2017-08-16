package application

import "net/http"

// MiddlewareStack middleware stack
type MiddlewareStack struct {
}

// Middleware HTTP middleware
type Middleware func(http.Handler) http.Handler

// Use use middleware
func (MiddlewareStack) Use(name string, middleware Middleware) {
	// TODO
}

// Run generate a new handler wrapped with registered middlewares
func (MiddlewareStack) Run(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		handler(w, req)
	})
}
