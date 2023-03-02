package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//All routes related to users comes here
	router.POST("/add/users", controllers.CreateUser())
	router.POST("/login/:userId/:userPassword", controllers.Login())
	router.GET("/user/:id", controllers.GetAUser())
	router.PUT("/user/:id", controllers.EditAUser())
	router.DELETE("/delete/user/:id", controllers.DeleteAUser())
	router.GET("/users/list", controllers.GetAllUsers())
}
