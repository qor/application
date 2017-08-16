package application

import "net/http"

// MiddlewareStack middleware stack
type MiddlewareStack struct {
	middlewares []*MiddlewareHandler
}

// Middleware HTTP middleware
type Middleware func(http.Handler) http.Handler

// MiddlewareHandler middleware handler
type MiddlewareHandler struct {
	stack   *MiddlewareStack
	before  string
	after   string
	Name    string
	Handler Middleware
}

// Use use middleware
func (stack *MiddlewareStack) Use(name string, middleware Middleware) {
	stack.middlewares = append(stack.middlewares, &MiddlewareHandler{
		stack:   stack,
		Name:    name,
		Handler: middleware,
	})
}

// Before insert middleware before name
func (stack *MiddlewareStack) Before(name string) MiddlewareHandler {
	return MiddlewareHandler{
		stack:  stack,
		before: name,
	}
}

// After insert middleware after name
func (stack *MiddlewareStack) After(name string) MiddlewareHandler {
	return MiddlewareHandler{
		stack: stack,
		after: name,
	}
}

// Use use middleware
func (handler MiddlewareHandler) Use(name string, middleware Middleware) {
	handler.Name = name
	handler.Handler = middleware
	if handler.stack != nil {
		handler.stack.middlewares = append(handler.stack.middlewares, &handler)
	}
}

// Run generate a new handler wrapped with registered middlewares
func (MiddlewareStack) Run(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		handler.ServeHTTP(w, req)
	})
}
