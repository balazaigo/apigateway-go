package routes

import (
	"main.go/utilities"
	"main.go/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Login(c *gin.Context) {
	var finalUrl string
	requestedEndpoint := c.Request.URL.Path
	if val, ok := utilities.Urls[requestedEndpoint]; ok {
		finalUrl = val
		remote, err := url.Parse(finalUrl)
		if err != nil {
			panic(err)
		}
		director := func(req *http.Request) {
			req.Header.Add("X-Forwarded-Host", c.Request.Host)
			req.Header.Add("X-Origin-Host", remote.Host)
			req.Host = remote.Host
			req.URL = remote
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	} else {
		helper.RespWriter(c, "Sorry! Requested Endpoint is Not available", http.StatusNotFound)
	}
}
