package main

import (
	"net/http"
	"todo-app/configs"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	configs.ConnectDB()
	r.SetTrustedProxies(nil)
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8080")
}
