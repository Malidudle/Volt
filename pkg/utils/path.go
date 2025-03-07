package utils

import "strings"

// NormalizePath removes trailing slash for path comparison
func NormalizePath(path string) string {
	if path != "/" && strings.HasSuffix(path, "/") {
		return strings.TrimSuffix(path, "/")
	}
	return path
}
