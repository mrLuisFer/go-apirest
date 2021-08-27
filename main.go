package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrLuisFer/go-apirest/src/methods"
)

func GetPosts(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "example",
	})
}

func main() {
	router := gin.Default()

	router.GET("/", methods.GetRootHandler)
	router.GET("/posts", GetPosts)

	error := router.Run()

	if error != nil {
		fmt.Print(error.Error())
		os.Exit(1)
	}
}
