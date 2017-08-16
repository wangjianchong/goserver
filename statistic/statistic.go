package main

import (
	"net"
	"strings"
	"time"

	"v1/mgo.v2/bson"
)

//一個tcp服務器,設備上傳數據,保存下來負責記錄
//标记
//协议  ip
//分布式
//获取数据的接口
var mongo Mongo

func main() {
	mongoInit()
	tcpListen()
}

func mongoInit() {
	err := mongo.Connection("")
	if err != nil {
		Error(err)
	}
}

func tcpListen() {
	l, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		Error(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			Error(err)
			return
		}
		go handleConn(c)
	}

}
func handleConn(c net.Conn) {
	var data []byte = make([]byte, 400)
	ret, err := c.Read(data)
	if err != nil {
		Error(err)
		goto here
	}
	Debug(data)
	if ret == int(data[0]) {
		strs := strings.Split(string(data[1:]), "_")
		istData(strs)
	}

here:
}

func intervalLog() {
	for {
		time.Sleep(60 * time.Second)
		Debug()
	}
}
func istData(strs []string) {
	Debug(strs)
	var mgoStu MgoStu
	mgoStu.DB = "statistic"
	mgoStu.C = "co"
	mgoStu.SelectBson = bson.M{"addr": strs[0], "server": strs[1], "key": strs[2]}
	switch strs[3] {
	case "add":
		mongo.Insert(mgoStu)
	}
}
