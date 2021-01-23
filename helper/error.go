package helper

import (
	"github.com/gin-gonic/gin"
)

func RespWriter (c *gin.Context, message string,statusCode int) {
	status := "failure"
	if statusCode == 200 {
		status = "success"
	}
	c.JSON(statusCode,gin.H{
		"message": message,
		"status": status,
	})
	return
}