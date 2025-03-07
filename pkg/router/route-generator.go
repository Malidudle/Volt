package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Route represents an API route
type Route struct {
	Path        string
	Method      string
	Handler     interface{}
	PackagePath string
	ImportName  string
}

// routesFileTemplate is the template for the generated routes.go file
const routesFileTemplate = `package generator

import (
	"net/http"
{{ range .Imports }}
	{{ .ImportName }} "volt/{{ .PackagePath }}"{{ end }}
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
{{ range .Routes }}		// {{ .Path }} route
		{
			Path:    "{{ .Path }}",
			Method:  {{ .ImportName }}.Method,
			Handler: {{ .ImportName }}.Handler,
		},
{{ end }}	}
}
`

// TemplateData holds data for the routes file template
type TemplateData struct {
	Imports []Route
	Routes  []Route
}

// GenerateRoutes scans the app directory for routes and generates a routes.go file
func GenerateRoutes() []Route {
	dir := "./app"
	var routes []Route

	// First, check if the root route.go exists
	rootRoutePath := filepath.Join(dir, "route.go")
	if _, err := os.Stat(rootRoutePath); err == nil {
		fmt.Printf("Found root route at %s -> API path: /api\n", rootRoutePath)

		routes = append(routes, Route{
			Path:        "/",
			ImportName:  "root",
			PackagePath: "/",
		})
	}

	// Now walk the directory to find other routes
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}

		// Skip the root directory itself and the root route.go file (already handled)
		if path == dir || path == rootRoutePath {
			return nil
		}

		// Skip directories themselves
		if info.IsDir() {
			return nil
		}

		// Look for route.go files
		if filepath.Base(path) == "route.go" {
			// Convert filesystem path to import path
			relPath, err := filepath.Rel(dir, path)
			if err != nil {
				fmt.Printf("Error getting relative path for %s: %v\n", path, err)
				return nil
			}

			// Remove the route.go part
			packagePath := filepath.Dir(relPath)

			// Convert to API URL path
			apiPath := "/" + strings.ReplaceAll(packagePath, string(os.PathSeparator), "/")

			fmt.Printf("Found route at %s -> API path: %s\n", path, apiPath)

			// Create import name (e.g., app/example -> exampleRoute)
			importName := filepath.Base(packagePath) + "Route"

			routes = append(routes, Route{
				Path:        apiPath,
				ImportName:  importName,
				PackagePath: filepath.Join("app", packagePath), // Full import path
			})
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory structure: %v\n", err)
	}

	// Generate the routes.go file
	if len(routes) > 0 {
		generateRoutesFile(routes)
	}

	return routes
}

// generateRoutesFile creates the routes.go file with all discovered routes
func generateRoutesFile(routes []Route) {
	templateData := TemplateData{
		Imports: routes,
		Routes:  routes,
	}

	// Create template
	tmpl, err := template.New("routes").Parse(routesFileTemplate)
	if err != nil {
		fmt.Printf("Error creating template: %v\n", err)
		return
	}

	// Create the output file
	outputFile := "./pkg/router/routes.go"
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", outputFile, err)
		return
	}
	defer file.Close()

	// Execute the template
	err = tmpl.Execute(file, templateData)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}

	fmt.Printf("Successfully generated routes file at %s\n", outputFile)
}
