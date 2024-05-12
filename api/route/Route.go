package route

import (
	"job-portal-project/api/config"
	"job-portal-project/api/helper"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	controller "job-portal-project/api/controllers"
	usercontroller "job-portal-project/api/controllers/user"
	repositoryimpl "job-portal-project/api/repositories/repositories-impl"
	userrepositoryimpl "job-portal-project/api/repositories/user/user-repositories-impl"
	serviceimpl "job-portal-project/api/services/service-impl"
	userserviceimpl "job-portal-project/api/services/user/user-service-impl"
)

func StartRouting(db *gorm.DB) {
	// Initialize Redis client
	// rdb := config.InitRedis()

	JobRepository := repositoryimpl.StartJobRepositoryImpl()
	JobService := serviceimpl.StartJobService(JobRepository, db)
	JobController := controller.NewJobController(JobService)

	UserRepository := userrepositoryimpl.NewUserRepository()
	UserService := userserviceimpl.NewUserService(UserRepository, db)
	UserController := usercontroller.NewUserController(UserService)

	RoleRepository := repositoryimpl.StartRoleRepositoryImpl()
	RoleService := serviceimpl.StartRoleService(RoleRepository, db)
	RoleController := controller.NewRoleController(RoleService)

	JobApplicationRepository := repositoryimpl.StartJobApplicationRepositoryImpl()
	JobApplicationService := serviceimpl.StartJobApplicationService(JobApplicationRepository, db)
	JobApplicationController := controller.NewJobApplicationController(JobApplicationService)

	/* Master */
	JobRouter := JobRouter(JobController)
	UserRouter := UserRouter(UserController)
	JobApplicationRouter := JobApplicationRouter(JobApplicationController)
	RoleRouter := RoleRouter(RoleController)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		// Tambahkan routing untuk v1 versi di sini
		/* Master */
		r.Mount("/job", JobRouter)
		r.Mount("/user", UserRouter)
		r.Mount("/job-application", JobApplicationRouter)
		r.Mount("/role", RoleRouter)

	})

	server := http.Server{
		Addr:    config.EnvConfigs.ClientOrigin,
		Handler: r,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
