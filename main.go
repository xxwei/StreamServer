package main

import (
	"fmt"
	"net"
	"net/http"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Stream Server By Golang")
	fmt.Println("Stream Server is Starting")
	time.Sleep(time.Second * 3)
	for {
		time.Sleep(time.Second * 3)
	}
}

func init() {
	fmt.Println("CPU number is %d", runtime.NumCPU())
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		conn.Close()
	}

}
func getconfig() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {

	}
	defer resp.Body.Close()
}
