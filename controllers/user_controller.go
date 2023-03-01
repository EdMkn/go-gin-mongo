package controllers

import (
	"context"
	"fmt"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()
var directorypath = "resources/"

func createFile(filename string, data string) {
	f, err := os.OpenFile(directorypath+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	} else {
		_, err = f.Write([]byte(data))
		if err != nil {
			fmt.Printf("error %v", err)
		}
	}

	f.Close()
}
func deleteFile(filename string) {
	// delete file
	var err0 = os.Chmod(directorypath+filename, 0770)
	var err = os.Remove(directorypath + filename)
	if err0 != nil {
		fmt.Printf("Unable to change file permission: %v", err0)
	}
	if err != nil {
		fmt.Printf("Unable to delete file: %v", err)
	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		var altuser models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&users); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		for _, user := range users {
			//use the validator library to validate required fields
			if validationErr := validate.Struct(user); validationErr != nil {
				c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error_struct", Data: map[string]interface{}{"data": validationErr.Error()}})
				continue
			}

			if uniqErr := userCollection.FindOne(ctx, bson.M{"id": user.Id}).Decode(&altuser); uniqErr == nil {
				c.JSON(http.StatusConflict, responses.UserResponse{Status: http.StatusConflict, Message: "error_uniq", Data: map[string]interface{}{"data": "utilisateur deja present"}})
				continue
			}
			hash := getHash([]byte(user.Password))
			user.Password = hash
			newUser := models.User{
				Id:         user.Id,
				Password:   user.Password,
				IsActive:   user.IsActive,
				Balance:    user.Balance,
				Age:        user.Age,
				Name:       user.Name,
				Gender:     user.Gender,
				Company:    user.Company,
				Email:      user.Email,
				Phone:      user.Phone,
				Address:    user.Address,
				About:      user.About,
				Registered: user.Registered,
				Latitude:   user.Latitude,
				Longitude:  user.Longitude,
				Tags:       user.Tags,
				Friends:    user.Friends,
				Data:       user.Data,
			}
			result, err := userCollection.InsertOne(ctx, newUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				continue
			}

			createFile(newUser.Id, newUser.Data)
			c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})

		}

	}
}

func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("id")
		var user models.User
		defer cancel()

		//objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"id": userId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}
func EditAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("id")
		var user models.User
		isDataChanged := false
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		/*if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error_struct", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		} */
		var update bson.D

		if len(user.Password) > 0 {
			update = append(update, bson.E{"password", getHash([]byte(user.Password))})
		}
		if reflect.ValueOf(user.IsActive).Kind() == reflect.Bool {
			update = append(update, bson.E{"isActive", user.Password})
		}
		if len(user.Balance) > 0 {
			update = append(update, bson.E{"balance", user.Balance})
		}
		if user.Age != 0 {
			update = append(update, bson.E{"age", user.Age})
		}

		if len(user.Name) > 0 {
			update = append(update, bson.E{"name", user.Name})
		}
		if len(user.Company) > 0 {
			update = append(update, bson.E{"company", user.Company})
		}
		if len(user.Email) > 0 {
			update = append(update, bson.E{"email", user.Email})
		}
		if len(user.Phone) > 0 {
			update = append(update, bson.E{"phone", user.Phone})
		}
		if len(user.Address) > 0 {
			update = append(update, bson.E{"address", user.Address})
		}
		if len(user.About) > 0 {
			update = append(update, bson.E{"about", user.About})
		}
		if len(user.Registered) > 0 {
			update = append(update, bson.E{"registered", user.Registered})
		}
		if len(fmt.Sprintf("%f", user.Latitude)) != 0 {
			update = append(update, bson.E{"latitude", user.Latitude})
		}
		if len(fmt.Sprintf("%f", user.Longitude)) != 0 {
			update = append(update, bson.E{"longitude", user.Longitude})
		}
		if user.Friends != nil {
			update = append(update, bson.E{"friends", user.Friends})
		}
		if len(user.Data) > 0 {
			update = append(update, bson.E{"data", user.Data})
			isDataChanged = true
		}

		/*
			update = bson.M{
				"id":         user.Id,
				"password":   user.Password,
				"isActive":   user.IsActive,
				"balance":    user.Balance,
				"age":        user.Age,
				"name":       user.Name,
				"gender":     user.Gender,
				"Company":    user.Company,
				"email":      user.Email,
				"phone":      user.Phone,
				"address":    user.Address,
				"about":      user.About,
				"registered": user.Registered,
				"latitude":   user.Latitude,
				"longitude":  user.Longitude,
				"tags":       user.Tags,
				"friends":    user.Friends,
				"data":       user.Data,
			}
		*/
		result, err := userCollection.UpdateOne(ctx, bson.M{"id": userId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated user details
		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"id": userId}).Decode(&updatedUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		if isDataChanged {
			deleteFile(userId)
			createFile(userId, user.Data)
		}
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedUser}})
	}
}
func DeleteAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("id")
		defer cancel()

		result, err := userCollection.DeleteOne(ctx, bson.M{"id": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
			)
			return
		}

		deleteFile(userId)
		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
		)
	}
}
func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			users = append(users, singleUser)
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
		)
	}
}
