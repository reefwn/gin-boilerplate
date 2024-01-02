package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func toResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func ReturnResponse(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, toResponse(code, http.StatusText(code), data))
}
