package controllers

import (
	"context"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		userPass := c.Param("userPassword")
		var dbuser models.User
		defer cancel()

		err := userCollection.FindOne(ctx, bson.M{"id": userId}).Decode(&dbuser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "identifiant introuvable", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		dbPass := []byte(dbuser.Password)
		passErr := bcrypt.CompareHashAndPassword(dbPass, []byte(userPass))
		if passErr != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "Mot de passe errone", Data: map[string]interface{}{"data": passErr.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": dbuser}})

	}
}
