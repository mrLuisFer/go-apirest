package methods

import (
	"github.com/gin-gonic/gin"
)

func GetPosts(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "example",
	})
}
