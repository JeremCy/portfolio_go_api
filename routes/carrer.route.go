package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jcyrille972/go_api_test/controllers"
)

func carrerRoutes(superRoute *gin.RouterGroup) {
	carrerRoutes := superRoute.Group("/carrer")
	{
		carrerRoutes.POST("/", (&controllers.ProjectController{}).Create)
		carrerRoutes.GET("/", (&controllers.ProjectController{}).FetchAll)
		carrerRoutes.GET("/:id", (&controllers.ProjectController{}).FetchOne)
		carrerRoutes.PUT("/:id", (&controllers.ProjectController{}).Update)
		carrerRoutes.DELETE("/:id", (&controllers.ProjectController{}).Delete)
	}
}
