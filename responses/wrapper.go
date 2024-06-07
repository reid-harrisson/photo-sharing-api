package responses

import (
	"github.com/gin-gonic/gin"
)

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(context *gin.Context, statusCode int, data interface{}) {
	context.IndentedJSON(statusCode, data)
}

func ErrorResponse(context *gin.Context, statusCode int, message string) {
	Response(context, statusCode, Error{
		Code:  statusCode,
		Error: message,
	})
}

func MessageResponse(context *gin.Context, statusCode int, message string) {
	Response(context, statusCode, Data{
		Code:    statusCode,
		Message: message,
	})
}
