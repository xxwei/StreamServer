package util

import (
	"fmt"
	"os"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	//log"
	"net"
)

func init() {
	//fmt.Println("util.init")
}

func Welcome() (ret bool) {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Println("Welcome ")
	fmt.Println("Stream Server By Golang")
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
		return false
	} else {
		fmt.Println("The address is ", addr.String())
		return true
	}

	//fmt.Println(addr)
}

func CheckErrorExit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
