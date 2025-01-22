package responses

import "github.com/gin-gonic/gin"

type ResponseStorage struct {
	URL string `json:"url" example:"https://example.com/123.jpg"`
}

func NewResponseStorage(context *gin.Context, statusCode int, url string) {
	Response(context, statusCode, ResponseStorage{
		URL: url,
	})
}
