package route

import (
	controller "job-portal-project/api/controllers"
	UserController "job-portal-project/api/controllers/user"
	middlewares "job-portal-project/api/middlewares"
	"net/http"

	// _ "job-portal-project/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

/* Master */

func JobRouter(
	JobController controller.JobController,
) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)

	router.With(func(next http.Handler) http.Handler {
		return middlewares.JWTAndRBACMiddleware([]string{"employer", "admin", "talent"}, next)
	}).Get("/", JobController.GetJobList)

	router.With(func(next http.Handler) http.Handler {
		return middlewares.JWTAndRBACMiddleware([]string{"employer", "admin", "talent"}, next)
	}).Get("/{job_id}", JobController.GetJobById)

	router.With(func(next http.Handler) http.Handler {
		return middlewares.JWTAndRBACMiddleware([]string{"employer", "admin"}, next)
	}).Post("/", JobController.SaveJob)

	router.With(func(next http.Handler) http.Handler {
		return middlewares.JWTAndRBACMiddleware([]string{"employer", "admin"}, next)
	}).Patch("/{job_id}", JobController.ChangeStatusJob)

	return router
}

func UserRouter(
	UserController UserController.UserController,
) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)

	// router.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		// Check if the route is for GetUsernameByUserID
	// 		if r.URL.Path == "/{user_id}" && r.Method == http.MethodGet {
	// 			// If it is, apply the middleware only to this route
	// 			middlewares.JWTAndRBACMiddleware([]string{"admin", "employer"}, next).ServeHTTP(w, r)
	// 			return
	// 		}
	// 		next.ServeHTTP(w, r)
	// 	})
	// })

	// User routes
	router.Post("/register", UserController.Register)
	router.Post("/login", UserController.Login)
	// router.Get("/{user_id}", UserController.GetUsernameByUserID)

	router.With(func(next http.Handler) http.Handler {
		return middlewares.JWTAndRBACMiddleware([]string{"employer", "admin"}, next)
	}).Get("/{user_id}", UserController.GetUsernameByUserID)

	return router
}

func RoleRouter(
	RoleController controller.RoleController,
) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)

	router.With(func(next http.Handler) http.Handler {
		return middlewares.JWTAndRBACMiddleware([]string{"employer", "admin", "talent"}, next)
	}).Get("/", RoleController.GetRoleList)

	router.With(func(next http.Handler) http.Handler {
		return middlewares.JWTAndRBACMiddleware([]string{"admin"}, next)
	}).Get("/{role_id}", RoleController.GetRoleById)

	router.With(func(next http.Handler) http.Handler {
		return middlewares.JWTAndRBACMiddleware([]string{"admin"}, next)
	}).Post("/", RoleController.SaveRole)

	router.With(func(next http.Handler) http.Handler {
		return middlewares.JWTAndRBACMiddleware([]string{"admin"}, next)
	}).Patch("/{role_id}", RoleController.ChangeStatusRole)

	return router
}

func JobApplicationRouter(
	JobApplicationController controller.JobApplicationController,
) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)

	router.With(func(next http.Handler) http.Handler {
		return middlewares.JWTAndRBACMiddleware([]string{"employer", "admin"}, next)
	}).Get("/{job_id}", JobApplicationController.GetJobApplicationListByJobId)

	router.With(func(next http.Handler) http.Handler {
		return middlewares.JWTAndRBACMiddleware([]string{"employer", "admin", "talent"}, next)
	}).Get("/{job_application_id}", JobApplicationController.GetJobApplicationById)

	router.With(func(next http.Handler) http.Handler {
		return middlewares.JWTAndRBACMiddleware([]string{"admin", "talent"}, next)
	}).Post("/", JobApplicationController.SaveJobApplication)

	router.With(func(next http.Handler) http.Handler {
		return middlewares.JWTAndRBACMiddleware([]string{"admin"}, next)
	}).Patch("/{job_application_id}", JobApplicationController.ChangeStatusJobApplication)

	router.With(func(next http.Handler) http.Handler {
		return middlewares.JWTAndRBACMiddleware([]string{"employer", "admin"}, next)
	}).Put("/update/{job_application_id}", JobApplicationController.UpdateJobApplication)

	return router
}
