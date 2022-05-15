package controllers

import (
	"context"
	"net/http"
	"time"
	"todo-app/configs"
	"todo-app/models"
	"todo-app/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollecttion *mongo.Collection   = configs.Collection(configs.DB, "users")
	validate        *validator.Validate = validator.New()
)

const (
	badRequest      = http.StatusBadRequest
	acceptedRequest = http.StatusAccepted
	OKRequest       = http.StatusOK
)

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		var user models.UserModel
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.BuildResponse(http.StatusBadRequest, "Error", err.Error()))
			return
		}

		if validateErr := validate.Struct(&user); validateErr != nil {
			c.AbortWithStatusJSON(badRequest, responses.BuildResponse(badRequest, "err", validateErr.Error()))
			return
		}

		newUser := models.UserModel{
			Id:       primitive.NewObjectID(),
			Fullname: user.Fullname,
			Email:    user.Email,
		}

		result, err := userCollecttion.InsertOne(ctx, newUser)
		if err != nil {
			c.AbortWithStatusJSON(badRequest, responses.BuildResponse(badRequest, "Error while inserting", err.Error()))
			return
		}

		c.JSON(acceptedRequest, responses.BuildResponse(acceptedRequest, "Success", result))
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		id := c.Param("id")
		objectID, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			c.AbortWithStatusJSON(badRequest, responses.BuildResponse(badRequest, "err", "Invalid id"))
			return
		}

		result, err := userCollecttion.DeleteOne(ctx, bson.D{{"id", objectID}})
		if err != nil {
			c.AbortWithStatusJSON(badRequest, responses.BuildResponse(badRequest, "err", err.Error()))
			return
		}

		c.JSON(OKRequest, responses.BuildResponse(OKRequest, "Success", result))

	}
}

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, time.Second*20)
		var user models.UserModel
		defer cancel()

		id := c.Param("id")
		objectID, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			c.AbortWithStatusJSON(badRequest, responses.BuildResponse(badRequest, "err", "Invalid id"))
			return
		}

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.BuildResponse(http.StatusBadRequest, "Error", err.Error()))
			return
		}

		if validateErr := validate.Struct(&user); validateErr != nil {
			c.AbortWithStatusJSON(badRequest, responses.BuildResponse(badRequest, "err", validateErr.Error()))
			return
		}

		updatedUser := bson.M{"id": user.Id, "fullname": user.Fullname, "email": user.Email, "password": user.Password}

		result, err := userCollecttion.UpdateOne(ctx, bson.M{"id": objectID}, bson.M{"$set": updatedUser})
	}
}

func DeleteAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		result, err := userCollecttion.DeleteMany(ctx, bson.M{})
		if err != nil {
			c.AbortWithStatusJSON(badRequest, responses.BuildResponse(badRequest, "Error while deleting", err.Error()))
			return
		}

		if result.DeletedCount < 1 {
			c.AbortWithStatusJSON(badRequest, responses.BuildResponse(badRequest, "err", "No users"))
			return
		}

		c.JSON(OKRequest, responses.BuildResponse(OKRequest, "Success", result.DeletedCount))

	}
}
