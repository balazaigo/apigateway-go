package utilities

import (
	"net/http"
	"net/url"
)

type EndPoint struct {
	Path string
	Url string
}

const Host = "http://devapi.zaicrm.com/"

const Prefix = Host + "api/"

var Urls = map[string]string{
	"/login": Prefix+"auth/login",
}


func GetMethodUrl (prefix string, r *http.Request) (finalUrl string){
	params := url.Values{}
	for key, val := range r.URL.Query() {
		params.Add(key,val[0])
	}
	finalUrl = Prefix+prefix+"?"+params.Encode()
	return
}