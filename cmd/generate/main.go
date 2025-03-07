package main

import (
	"fmt"
	"os"

	generator "backend-framework/pkg/router"
)

func main() {
	fmt.Println("Generating API routes...")
	routes := generator.GenerateRoutes()
	fmt.Printf("Found %d routes\n", len(routes))

	for _, route := range routes {
		fmt.Printf("Route: %s\n", route.Path)
	}

	fmt.Println("Route generation complete!")
	fmt.Println("Generated routes file at ./pkg/router/routes.go")

	os.Exit(0)
}
