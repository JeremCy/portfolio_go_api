package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jcyrille972/go_api_test/controllers"
)

func userRoutes(superRoute *gin.RouterGroup) {
	userRoutes := superRoute.Group("/user")
	{
		userRoutes.GET("/", (&controllers.UserController{}).FetchAll)
		userRoutes.GET("/:id", (&controllers.UserController{}).FetchOne)
		userRoutes.PUT("/:id", (&controllers.UserController{}).Update)
		userRoutes.DELETE("/:id", (&controllers.UserController{}).Delete)
	}
}
