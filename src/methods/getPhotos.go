package methods

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ID     string `json:"id"`
	AUTHOR string `json:"author"`
	WIDTH  int    `json:"width"`
	HEIGHT int    `json:"height"`
	URL    string `json:"url"`
}

// https://picsum.photos/v2/list
var data = []Response{
	{
		ID:     "1025",
		AUTHOR: "Matthew Wiebe",
		WIDTH:  4951,
		HEIGHT: 3301,
		URL:    "https://unsplash.com/photos/U5rMrSI7Pn4",
	},
}

func GetPhotos(context *gin.Context) {
	context.JSON(http.StatusOK, data)
}
