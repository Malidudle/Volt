package example

import (
	"net/http"
	"time"
)

// Method defines the HTTP method for this route
var Method = http.MethodGet

// Handler handles requests to this route
func Handler(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	response := map[string]interface{}{
		"message": "Hot reloading working on port 8080! Updated at " + time.Now().Format(time.RFC3339),
		"time":    time.Now().Format(time.RFC3339),
		"path":    r.URL.Path,
		"method":  r.Method,
	}

	return response, http.StatusOK, nil
}
