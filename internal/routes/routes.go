package routes

import (
	"net/http"

	"api-server/internal/handlers"
	"api-server/internal/middleware"
)

// Register all the application routes
func RegisterRoutes() {
	// Public routes
	http.HandleFunc("/healthz", handlers.HealthCheckHandler) // GET /healthz
	http.HandleFunc("/v1/user", handlers.CreateUserHandler)  // POST /v1/user
	// http.HandleFunc("/v1/instructor/", handlers.GetInstructorHandler) // GET /v1/instructor/{instructor_id}
	// http.HandleFunc("/v1/course/", handlers.GetCourseHandler) // GET /v1/course/{course_id}

	// Private routes
	http.HandleFunc("/v1/user/", middleware.AuthMiddleware(handlers.UserHandler)) // GET & PUT /v1/user/{user_id}
	// examples for future routes
	// http.HandleFunc("/v1/instructor", middleware.AuthMiddleware(handlers.CreateInstructorHandler))
	// http.HandleFunc("/v1/course", middleware.AuthMiddleware(handlers.CreateCourseHandler))
}
