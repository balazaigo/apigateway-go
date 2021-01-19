package main

import (
	"github.com/gin-gonic/gin"
	"main.go/routes"
)

func RequestHandler() {
	r := gin.Default()
	r.Any("/login", routes.Login)
	r.Run(":10000")
}
