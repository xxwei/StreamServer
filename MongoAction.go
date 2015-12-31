package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

func InitMongo() {
	var err error
	fmt.Println("Init Mongo")
	session, err = mgo.Dial("10.192.165.195:20000,10.192.165.196:20000") //链接服务器
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	//defer session.Close()
}

func CheckUser(name string, pass string) (ret bool, err error) {
	session.DB("SStream").C("User").Find(bson.M{"name": name, "pass": pass})
	ret = true
	err = nil
	return
}

func CheckAppKey(appkey string) (ret bool, err error) {
	ret = true
	err = nil
	return
}
