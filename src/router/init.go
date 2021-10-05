package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrLuisFer/go-apirest/src/methods"
)

func Init(router *gin.Engine) {

	router.GET("/", methods.GetRootHandler)
	router.GET("/posts", methods.GetPosts)

	error := router.Run()

	if error != nil {
		fmt.Print(error.Error())
		os.Exit(1)
	}	
}
