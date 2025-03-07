package main

import (
	"log"
	"os"

	"volt/pkg/handler"
	generator "volt/pkg/router"
	"volt/pkg/server"
)

func main() {
	routes := generator.GetAllRoutes()
	if err := validateRoutes(routes); err != nil {
		log.Fatalf("Invalid route configuration: %v", err)
	}
	log.Printf("Loaded %d routes", len(routes))

	port := os.Getenv("PORT")
	srv := server.New(port)
	apiHandler := handler.NewAPIHandler()

	if err := srv.Start(apiHandler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func validateRoutes(routes []generator.RegisteredRoute) error {
	for _, route := range routes {
		log.Printf("Route: %s %s", route.Method, route.Path)
	}
	return nil
}
