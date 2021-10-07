package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrLuisFer/go-apirest/src/methods"
)

func Init(router *gin.Engine) {
	var port string = ":3000"

	router.GET("/", methods.GetRootHandler)
	router.GET("/photos", methods.GetPhotos)
	router.GET("/photos/:id", methods.GetPhoto)

	error := router.Run(port)

	if error != nil {
		fmt.Print(error.Error())
		os.Exit(1)
	}
}
