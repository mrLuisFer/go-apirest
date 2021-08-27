package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func GetPosts(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "example",
	})
}

func main() {
	router := gin.Default()

	router.GET("/posts", GetPosts)

	error := router.Run()

	if error != nil {
		fmt.Print(error.Error())
		os.Exit(1)
	}
}
