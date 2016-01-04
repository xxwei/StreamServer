package main

import (
	"fmt"
	"os"
	"runtime"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"log"
	"net"
)

func init() {
	//fmt.Println("util.init")
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()

	log.SetFlags(log.Llongfile)
	log.Print("")
	/*
		var testslice [100]int

		var slice []int = testslice[3:20]

		slice1 := make([]int, 10, 50)

		slice1 = slice1[0 : len(slice1)+5] //切片扩容

		slice2 := copy(slice, slice1)

		//slice2 = append(slice, 4, 5, 6) //追加新元素
	*/
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
