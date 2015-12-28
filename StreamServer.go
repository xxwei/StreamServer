package main

import (
	"fmt"
	"net"
	"strconv"
)

var localaddr string

func InitStreamServer(port int) {

	localaddr = ":" + strconv.Itoa(port)
	fmt.Println("Server Start At ", localaddr)
}

func StartStreamServer() {
	if tcpaddr, err := net.ResolveTCPAddr("tcp4", localaddr); err != nil {
		panic(err)
	} else {
		net.ListenTCP("tcp", tcpaddr)
	}
}

func CreateStreamChannel(uuid string) {

}
