package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrLuisFer/go-apirest/src/router"
)

func main() {

	routerGin := gin.Default()

	router.Init(routerGin)
}
