package main

import (
	//"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func InitHttpServer() {
	fmt.Println("Init HttpServer")
	r := httprouter.New()

	r.GET("/", homeHandler)

	r.GET("/api", apiHandler)
	r.POST("/api", apiHandler)

	r.GET("/api/login", loginHandler)
	r.POST("/api/login", loginHandler)

	go http.ListenAndServe(":8080", r)

	/*
		resp, err := http.Get("http://www.baidu.com")
		if err != nil {
		}
		defer resp.Body.Close()
	*/
}
func homeHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Hello Go Server")
}
func apiHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	rw.Header().Set("Content-Type", "application/json")

}

func loginHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	rw.Header().Set("Content-Type", "application/json")
	req.FormValue("user")
	req.FormValue("pass")
}
