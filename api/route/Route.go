package route

import (
	"job-portal-project/api/config"
	"job-portal-project/api/helper"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

	controller "job-portal-project/api/controllers"
	repositoryimpl "job-portal-project/api/repositories/repositories-impl"
	serviceimpl "job-portal-project/api/services/service-impl"
)

func StartRouting(db *gorm.DB) {
	// Initialize Redis client
	// rdb := config.InitRedis()

	JobRepository := repositoryimpl.StartJobRepositoryImpl()
	JobService := serviceimpl.StartJobService(JobRepository, db)
	JobController := controller.NewJobController(JobService)

	/* Master */
	JobRouter := JobRouter(JobController)
	r := chi.NewRouter()
	// Route untuk setiap versi API
	r.Route("/v1", func(r chi.Router) {
		// Tambahkan routing untuk v1 versi di sini
		/* Master */
		r.Mount("/job", JobRouter)

	})

	server := http.Server{
		Addr:    config.EnvConfigs.ClientOrigin,
		Handler: r,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
