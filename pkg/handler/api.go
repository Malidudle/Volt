package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	generator "volt/pkg/router"
	"volt/pkg/utils"
)

// APIHandler handles API requests
type APIHandler struct{}

// NewAPIHandler creates a new API handler
func NewAPIHandler() *APIHandler {
	return &APIHandler{}
}

// ServeHTTP handles HTTP requests
func (h *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routes := generator.GetAllRoutes()

	if err := validateRoutes(routes); err != nil {
		sendError(w, http.StatusInternalServerError, "Server Configuration Error", err.Error())
		return
	}

	normalizedPath := utils.NormalizePath(r.URL.Path)

	var matchedRoute *generator.RegisteredRoute
	for _, route := range routes {
		normalizedRoutePath := utils.NormalizePath(route.Path)
		if normalizedRoutePath == normalizedPath && route.Method == r.Method {
			matchedRoute = &route
			break
		}
	}

	if matchedRoute == nil {
		handleNotFound(w, r, routes)
		return
	}

	response, statusCode, err := matchedRoute.Handler(w, r)
	if err != nil {
		log.Printf("Error in handler: %v", err)
		sendError(w, http.StatusInternalServerError, "API Processing Error", err.Error())
		return
	}

	sendJSON(w, statusCode, response)
}

func validateRoutes(routes []generator.RegisteredRoute) error {
	for _, route := range routes {
		if route.Method == nil {
			return fmt.Errorf("Path '%s' missing HTTP method", route.Path)
		}

		if route.Handler == nil {
			return fmt.Errorf("Path '%s' missing handler function", route.Path)
		}

		handlerType := reflect.TypeOf(route.Handler)
		if handlerType == nil || handlerType.Kind() != reflect.Func {
			return fmt.Errorf("Path '%s' has invalid handler", route.Path)
		}
	}
	return nil
}

func handleNotFound(w http.ResponseWriter, r *http.Request, routes []generator.RegisteredRoute) {
	var availableRoutes []string
	for _, route := range routes {
		availableRoutes = append(availableRoutes, fmt.Sprintf("%s %s", route.Method, route.Path))
	}

	data := map[string]interface{}{
		"error":            "Route Not Found",
		"message":          fmt.Sprintf("Endpoint '%s' with method '%s' doesn't exist", r.URL.Path, r.Method),
		"request_path":     r.URL.Path,
		"request_method":   r.Method,
		"available_routes": availableRoutes,
		"timestamp":        time.Now().Format(time.RFC3339),
	}

	sendJSON(w, http.StatusNotFound, data)
}

func sendError(w http.ResponseWriter, status int, title, details string) {
	data := map[string]interface{}{
		"error":     title,
		"details":   details,
		"timestamp": time.Now().Format(time.RFC3339),
	}
	sendJSON(w, status, data)
}

func sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
