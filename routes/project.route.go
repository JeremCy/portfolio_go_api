package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jcyrille972/go_api_test/controllers"
)

//create crud route for project
func projectRoutes(superRoute *gin.RouterGroup) {
	projectRoutes := superRoute.Group("/project")
	{
		projectRoutes.POST("/", (&controllers.ProjectController{}).Create)
		projectRoutes.GET("/", (&controllers.ProjectController{}).FetchAll)
		projectRoutes.GET("/:id", (&controllers.ProjectController{}).FetchOne)
		projectRoutes.PUT("/:id", (&controllers.ProjectController{}).Update)
		projectRoutes.DELETE("/:id", (&controllers.ProjectController{}).Delete)
	}
}
