package utils

type Response struct {
	ID     string `json:"id"`
	AUTHOR string `json:"author"`
	WIDTH  int    `json:"width"`
	HEIGHT int    `json:"height"`
	URL    string `json:"url"`
}

func GetData() []Response {
	var data = []Response{
		{
			ID:     "1025",
			AUTHOR: "Matthew Wiebe",
			WIDTH:  4951,
			HEIGHT: 3301,
			URL:    "https://unsplash.com/photos/U5rMrSI7Pn4",
		},
	}

	return data
}
