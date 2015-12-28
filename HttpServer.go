package main

import (
	"fmt"
	"net/http"
)

func InitHttpServer() {

	fmt.Println("Init HttpServer")
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
	}
	defer resp.Body.Close()
}
