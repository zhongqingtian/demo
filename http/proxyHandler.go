package main

import (
	"net/http"
	"strings"
)

func ProxyHandler(wr http.ResponseWriter, req *http.Request) {
	//认证
	user, password, ok := req.BasicAuth()
	if !ok {
		http.Error(wr, "basic auth required", http.StatusForbidden)
		return
	}

	if user != "[YOUR-USER]" || password != "[YOUR-PASSWORD]" {
		http.Error(wr, "basic auth failed", http.StatusForbidden)
		return
	}

	//墙外包
	if strings.HasPrefix(req.URL.RequestURI(), "cloud.google.com") {
	//	http.FileServer("[PrivateModulePath]").ServeHTTP(wr, req)
		return
	}

	//私有包
	if strings.HasPrefix(req.URL.RequestURI(), "your.company.com") {
	//	http.FileServer("[PrivateModulePath]").ServeHTTP(wr, req)
		return
	}

	//404
	http.NotFound(wr, req)
}
