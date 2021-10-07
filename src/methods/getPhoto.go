package methods

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPhoto(ctx *gin.Context) {
	// params := ctx.Request.URL.Query()
	params := ctx.Params

	ctx.JSON(http.StatusOK, gin.H{
		"params": params,
	})
}
