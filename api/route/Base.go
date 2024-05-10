package route

import (
	controller "job-portal-project/api/controllers"
	"job-portal-project/api/middlewares"

	// _ "job-portal-project/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

/* Master */

func JobRouter(
	JobController controller.JobController,
) chi.Router {
	router := chi.NewRouter()

	// Apply the CORS middleware to all routes
	router.Use(middlewares.SetupCorsMiddleware)
	router.Use(middleware.Recoverer)
	router.Use(middlewares.MetricsMiddleware)

	router.Get("/", JobController.GetJobList)
	router.Post("/", JobController.SaveJob)
	router.Get("/{job_id}", JobController.GetJobById)
	router.Patch("/{job_id}", JobController.ChangeStatusJob)

	return router
}
