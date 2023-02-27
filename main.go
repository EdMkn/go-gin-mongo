package main

import (
	"gin-mongo-api/configs" //add this

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	router.Run("localhost:8080")
}
