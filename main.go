package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"log"
	"net"
	"net/http"
	"runtime"
	"time"
	"flag"
)

func main() {
	fmt.Println("Stream Server By Golang")
	fmt.Println("Stream Server is Starting at ",time.Now())
	time.Sleep(time.Second)
	var port = flag.Int("port", 9080, "the stream  login auth port")
	var nbCpus = flag.Int("cpus", runtime.NumCPU()-1, "the cpus to use")
	var tcpNoDalay = flag.Bool("nodelay", false, "set TCP_NODELAY for connection")
	flag.Usage = func () {
		fmt.Println("Usage:[--port=int]")
		fmt.Println("		port,the listen port .default 9080")
	}
	flag.Parse()
	runtime.GOMAXPROCS(*nbCpus)
}

func init() {
	fmt.Println("CPU number is", runtime.NumCPU())

	session, err := mgo.Dial("10.192.165.195:20000,10.192.165.196:20000")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer session.Close()
	session.

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
