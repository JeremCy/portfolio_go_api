package routes

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(superRoute *gin.RouterGroup) {
	projectRoutes(superRoute)
	carrerRoutes(superRoute)
	userRoutes(superRoute)
	authRoutes(superRoute)
}
