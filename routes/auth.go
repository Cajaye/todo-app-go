package routes

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	routes := r.Group("/api/v1/auth")
	{
		routes.POST("/register", controllers.RegisterUser())
	}
}
