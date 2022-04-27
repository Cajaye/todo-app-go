package controllers

import (
	"net/http"
	"todo-app/configs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollecttion *mongo.Collection   = configs.Collection(configs.DB, "users")
	validate        *validator.Validate = validator.New()
)

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"register": "user"})
	}
}
