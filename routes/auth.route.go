package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jcyrille972/go_api_test/controllers"
)

func authRoutes(superRoute *gin.RouterGroup) {
	authRoutes := superRoute.Group("/auth")
	{
		authRoutes.POST("/login", (&controllers.AuthController{}).Login)
		authRoutes.POST("/register", (&controllers.AuthController{}).Register)
	}
}
