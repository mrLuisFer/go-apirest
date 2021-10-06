package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrLuisFer/go-apirest/src/router"
)

func main() {

	routerGin := gin.Default()
	routerGin.LoadHTMLGlob("src/templates/*.tpml")

	router.Init(routerGin)
}
