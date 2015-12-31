package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {

	Welcome()
	//InitMongo()
	InitHttpServer()

	time.Sleep(time.Second)

	var port = flag.Int("port", 9080, "the stream  login auth port")
	var nbCpus = flag.Int("cpus", runtime.NumCPU()-1, "the cpus to use")
	//var tcpNoDalay = flag.Bool("nodelay", false, "set TCP_NODELAY for connection")
	flag.Usage = func() {
		fmt.Println("Usage:[--port=int]")
		fmt.Println("		port,the listen port .default 9080")
	}
	flag.Parse()
	runtime.GOMAXPROCS(*nbCpus)
	InitStreamServer(*port)
	go StartStreamServer()
	beOver := 0
	for {
		time.Sleep(time.Second)
		if beOver > 3000 {
			break
		}
		fmt.Println(beOver, "was runing... ")
		beOver++
	}
}
