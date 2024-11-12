package helpers

import (
	"github.com/gin-gonic/gin"
)

type ResponseHelper struct {
	Code int
	Message string
}

func Response(c *gin.Context, data any, options ResponseHelper) {
	code := options.Code
	message := options.Message
	if code == 0 {
		code = 200
	}

	if message == "" {
		message = "success"
	}
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

