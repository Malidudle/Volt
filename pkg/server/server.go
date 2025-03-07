package server

import (
	"log"
	"net/http"
)

// Server represents the HTTP server
type Server struct {
	Port string
}

// New creates a new server instance
func New(port string) *Server {
	// Use provided port or default to 8080
	if port == "" {
		port = "8080"
	}
	return &Server{Port: port}
}

// Start starts the HTTP server
func (s *Server) Start(handler http.Handler) error {
	log.Printf("Server starting on port %s", s.Port)
	return http.ListenAndServe(":"+s.Port, handler)
}
