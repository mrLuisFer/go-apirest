package methods

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrLuisFer/go-apirest/src/utils"
)

func GetPhotos(context *gin.Context) {
	context.JSON(http.StatusOK, utils.GetData())
}
