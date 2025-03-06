package routes

import (
	"net/http"

	"api-server/internal/handlers"
)

// RegisterRoutes registers all the application routes.
func RegisterRoutes() {
	http.HandleFunc("/healthz", handlers.HealthCheckHandler)
}
