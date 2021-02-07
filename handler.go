package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main.go/helper"
	"main.go/utilities"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const jwtKey = "7Bn6dTXzPIVYFTyPHSCGZH9RhKY1JyebTbCeazM82wB0xWNvNA94FSH3zAbiFvca"

func RequestHandler() {
	r := gin.Default()
	r.Any("/api/*any", func(c *gin.Context) {
		finalUrl := utilities.Host + c.Request.URL.String()
		if strings.Contains(c.Request.URL.String(), "/auth/login") {
			proxyRequest(finalUrl, c)
			return
		} else {
			tokenFromRequest := getSubString(c.Request.Header.Get("Authorization"), "Bearer ")
			tkn, err := jwt.Parse(tokenFromRequest, func(token *jwt.Token) (i interface{}, err error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(jwtKey), nil
			})
			if err != nil {
				helper.RespWriter(c, "something went wrong "+err.Error(), http.StatusBadGateway)
			}

			if _, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
			//if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
				//claims -> will retrieve the claims from the token, in case if we need to validate additionally we need to access this map object and validate it.
				proxyRequest(finalUrl, c)
				return
			}
		}
	})
	r.Run(":10000")
}

func getSubString(token string, splitstring string) (s string) {
	splitString := strings.Split(token, splitstring)
	s = splitString[1]
	return
}

func proxyRequest(rawurl string, c *gin.Context) {
	remote, err := url.Parse(rawurl)
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
	return
}
