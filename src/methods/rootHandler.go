package methods

import (
	"github.com/gin-gonic/gin"
)

func GetRootHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"msg": "response",
	})

}
