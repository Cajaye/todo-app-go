package main

import (
	"todo-app/configs"
	"todo-app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	configs.ConnectDB()
	r.SetTrustedProxies(nil)
	r.Use(cors.Default())
	routes.AuthRoutes(r)
	r.Run("localhost:8080")
}
