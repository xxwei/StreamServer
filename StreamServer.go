package main

import (
	"StreamServer/util"
	//"errors"
	"fmt"
	"net"
	//"os"
	"strconv"
)

var localaddr string

func InitStreamServer(port int) {

	localaddr = ":" + strconv.Itoa(port)
	fmt.Println("Server Start At ", localaddr)
}

func StartStreamServer() {
	// net.ResolveTCPAddr is create a TCPAddr
	if tcpaddr, err := net.ResolveTCPAddr("tcp4", localaddr); err != nil {
		panic(err)
	} else {
		listener, err := net.ListenTCP("tcp", tcpaddr)
		util.CheckErrorExit(err)
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			uuid := util.Rand()
			ret, err := createStreamChannel(conn, uuid.Hex())
			if ret {
				go handleClient(conn, uuid.Hex())
			}
		}
	}
}

func createStreamChannel(c net.Conn, uuid string) (ret bool, err error) {
	ret = true
	return
}

func handleClient(c net.Conn, uuid string) {

}

func closeStreamChannel(c net.Conn, uuid string) (ret bool, err error) {
	ret = true
	return
}

func AddStreamAddr(uuid string, addr string) (ret bool, err error) {
	ret = true
	return
}
