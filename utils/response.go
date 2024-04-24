package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

const (
	ERROR   = 500
	SUCCESS = 200
)

func newResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{code, message, data})
}
func Ok(c *gin.Context) {
	newResponse(c, SUCCESS, "Success", nil)
}
func OkWithData(c *gin.Context, data interface{}) {
	newResponse(c, SUCCESS, "Success", data)
}
func FailWithMsg(c *gin.Context, msg string) {
	newResponse(c, ERROR, msg, nil)
}
