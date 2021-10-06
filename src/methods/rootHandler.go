package methods

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRootHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tpml", gin.H{
		"title": "Main Website",
	})
}
