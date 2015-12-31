package main

import (
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
		CheckErrorExit(err)
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			uuid := Rand()
			ret, err := createTCPStreamChannel(conn, uuid.Hex())
			if ret {
				go handleClient(conn, uuid.Hex())
			}
		}
	}
}

func createTCPStreamChannel(c net.Conn, uuid string) (ret bool, err error) {
	fmt.Println("Remote ip addr ", c.RemoteAddr())
	ret = false
	return
}

func handleClient(c net.Conn, uuid string) {

}

func closeTCPStreamChannel(c net.Conn, uuid string) (ret bool, err error) {
	ret = true
	return
}

func AddStreamDestAddr(uuid string, addr string) (ret bool, err error) {
	ret = true
	return
}

func DelStreamDestAddr(uuid string, add string) (ret bool, err error) {
	ret = true
	return
}

func FindStream(uuid string) (ret bool, err error) {
	ret = true
	return
}
func ListStreamFormatJson() (ret string, err error) {
	ret = ""
	return
}
