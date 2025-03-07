package generator

import (
	"net/http"

	exampleRoute "volt/app/example"
	testRoute "volt/app/test"
)

// RegisteredRoute represents a registered API route with its handler and method
type RegisteredRoute struct {
	Path    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request) (interface{}, int, error)
}

// GetAllRoutes returns all registered API routes
func GetAllRoutes() []RegisteredRoute {
	return []RegisteredRoute{
		// /example route
		{
			Path:    "/example",
			Method:  exampleRoute.Method,
			Handler: exampleRoute.Handler,
		},
		// /test route
		{
			Path:    "/test",
			Method:  testRoute.Method,
			Handler: testRoute.Handler,
		},
	}
}
