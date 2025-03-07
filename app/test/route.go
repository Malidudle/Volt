package test

import (
	"net/http"
	"time"
)

// Method defines the HTTP method for this route
var Method = http.MethodGet

func Handler(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	response := map[string]interface{}{
		"message": "Hot reloading is working! " + time.Now().Format(time.RFC3339),
		"time":    time.Now().Format(time.RFC3339),
		"path":    r.URL.Path,
		"method":  r.Method,
	}

	return response, http.StatusOK, nil
}
