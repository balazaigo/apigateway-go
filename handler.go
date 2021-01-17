package main

import (
	"net/http/httputil"
	"net/url"
	"main.go/utilities"
	"net/http"
	"github.com/gin-gonic/gin"
	"sync"
	"sync/atomic"
)

type Data struct {
	message string
	status int
}

//type Backend struct {
//	URL          *url.URL
//	Alive        bool
//	mux          sync.RWMutex
//	ReverseProxy *httputil.ReverseProxy
//}
//
//type ServerPool struct {
//	backends []*Backend
//	current  uint64
//}
//
//func (s *ServerPool) NextIndex() int {
//	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
//}

func RequestHandler() {
	var finalUrl string
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
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
		}else{
			respWriter(c,"Sorry! Requested Endpoint is Not available",http.StatusNotFound)
		}
	})

	//http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request) {
	//	requestedEndpoint := path.Base(r.URL.Path)
	//	if val, ok := utilities.Urls[requestedEndpoint]; ok {
	//		finalUrl = val
	//	}else{
	//		errResp(w, "Sorry! Requested Endpoint is Not available",http.StatusNotFound)
	//		return
	//	}
	//	//finalUrl = utilities.Prefix + "auth/login"
	//	if r.Method == "GET" {
	//		finalUrl = utilities.GetMethodUrl("auth/login",r)
	//		return
	//	}
	//	body, err := ioutil.ReadAll(r.Body)
	//	w.Header().Set("Content-Type","application/json")
	//	if err != nil {
	//		errResp(w,"Error while parsing your inputs, please try again later",http.StatusNoContent)
	//		fmt.Println("I am calling from first err")
	//		return
	//	}
	//	// you can reassign the body if you need to parse it as multipart
	//	r.Body = ioutil.NopCloser(bytes.NewReader(body))
	//
	//	proxyReq, err := http.NewRequest(r.Method,finalUrl,bytes.NewReader(body))
	//
	//	proxyReq.Header = make(http.Header)
	//	for h, val := range r.Header {
	//		proxyReq.Header[h] = val
	//	}
	//
	//	httpClient := http.Client{}
	//	resp, err := httpClient.Do(proxyReq)
	//	if err != nil {
	//		errResp(w,"Sorry! First Server is down",http.StatusBadGateway)
	//		fmt.Println("I am calling from second err",err.Error())
	//		return
	//	}
	//	if err == nil {
	//		body, _ := ioutil.ReadAll(resp.Body)
	//		w.Write([]byte(body))
	//		return
	//	}
	//
	//	if err != nil {
	//		panic(err.Error())
	//		return
	//	}
	//	//json.NewEncoder(w).Encode(body)
	//})
	//return "not called"
	//c.JSON(400,gin.H{
	//	"message": "Default return",
	//	"status": "Nothing",
	//})
	r.Run(":10000")
}

func respWriter (c *gin.Context, message string,statusCode int) {
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

