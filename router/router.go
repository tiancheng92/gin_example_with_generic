package router

import (
	"gin_example_with_generic/controller/api/universal"
	v1 "gin_example_with_generic/controller/api/v1"
	crossDomain "gin_example_with_generic/pkg/http/middleware/cross_domain"
	"gin_example_with_generic/pkg/http/middleware/handle_error"
	"gin_example_with_generic/pkg/http/middleware/logging"
	"gin_example_with_generic/store"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// InitRouter initialize route
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.LoggerWithConfig(logging.AccessLog))

	pprof.Register(r, "/debug/pprof")

	r.Use(gin.Recovery())

	r.Use(handle_error.HandleError)

	r.Use(crossDomain.CrossDomain())

	r.NoRoute(universal.NoRoute)

	r.GET("/healthz", universal.HealthCheck)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	api := r.Group("/api")
	{
		V1 := api.Group("/v1")
		{
			db := store.GetDefaultDB()

			countryController := v1.NewCountryController(db)
			country := V1.Group("/country")
			{
				country.GET("/:pk", countryController.Get)
				country.GET("", countryController.List)
				country.POST("", countryController.Create)
				country.PUT("/:pk", countryController.Update)
				country.DELETE("/:pk", countryController.Delete)
			}

			userController := v1.NewUserController(db)
			user := V1.Group("/user")
			{
				user.GET("/:pk", userController.Get)
				user.GET("", userController.List)
				user.GET("/name/:name", userController.ListByName)
				user.POST("", userController.Create)
				user.PUT("/:pk", userController.Update)
				user.DELETE("/:pk", userController.Delete)
			}
		}
	}
	return r
}
