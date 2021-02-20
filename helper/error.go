package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

func GetJwtKey () string {
	viper.SetConfigFile("/var/www/html/zaiserve-api-saas/.env")
	viper.ReadInConfig()
	return viper.GetString("JWT_SECRET")
}