package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func InitHttpServer() {

	fmt.Println("Init HttpServer")
	r := httprouter.New()

	r.GET("/", homeHandler)

	r.GET("/api", apiHandler)

	//http.Handle("/",  http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", r)

	/*
		resp, err := http.Get("http://www.baidu.com")
		if err != nil {
		}
		defer resp.Body.Close()
	*/
}
func homeHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

}
func apiHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

}
