package main

import (
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var chanNote chan int
var nowCount int64
var lastCount int64
var mutex sync.RWMutex
var numCount int64

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		Error(err)
	}

	addr := "10.10.2.149:0101_"
	server := "testserver_"
	key := "testkey_add"

	str := addr + server + key
	lenth := len([]byte(str)) + 1
	var sendBtyte []byte = make([]byte, 0)
	sendBtyte = append(sendBtyte, byte(lenth))
	for i := 0; i < lenth-1; i++ {
		sendBtyte = append(sendBtyte, []byte(str)[i])
	}
	ret, err := conn.Write(sendBtyte)
	if err != nil {
		Error(err)
	}
	Debug(ret)
	Debug(len(sendBtyte))
	Debug(lenth)
}

func main1() {
	SetLogLevel("ERROR")
	a := 3000
	chanNote = make(chan int, a)
	for i := 0; i < a; i++ {
		go sendMsg(i)
	}

	go staticstic()
	for i := 0; i < a; i++ {
		<-chanNote
	}
}
func staticstic() {
	var temp int64
	for {
		time.Sleep(time.Second)
		temp = atomic.LoadInt64(&nowCount)
		Debug("seond statistic: ", temp-lastCount)
		Error(numCount)
		lastCount = temp
	}
}
func sendMsg(num int) {
	defer func() {
		chanNote <- 1
	}()
	var tcpAddr *net.TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:7777")
	if err != nil {
		Error(err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		Error(err)
		return
	}
	defer conn.Close()
	var msg string
	for i := 0; i < 2000; i++ {
		msg = msg + "1"
	}
	b := []byte(msg)

	var rev string
	for i := 0; i < 2000; i++ {
		rev = rev + "a"
	}
	lenth := 0
	wlen := 0
	for i := 0; i < 10000; i++ {
		conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		wlen = 0
		for {
			templen, err := conn.Write(b[wlen:])
			wlen += templen
			if err != nil {
				Error(err)
				atomic.AddInt64(&numCount, 1)
				time.Sleep(time.Hour)
				return
			}
			if wlen == 2000 {
				break
			}
		}
		data := make([]byte, 2000)
		lenth = 0
		for {
			templen, err := conn.Read(data[lenth:])
			lenth += templen
			if err != nil {
				Error(err)
				atomic.AddInt64(&numCount, 1)
				time.Sleep(time.Hour)
				return
			}
			if lenth == 2000 {
				break
			}
		}
		Info("received:", string(data))
		if string(data) != rev {
			Error("not the same", string(data))
			time.Sleep(time.Hour)
		}

		//		_, err = conn.Write(b)
		//		mutex.Lock()
		atomic.AddInt64(&nowCount, 1)
		//		mutex.Unlock()
		//		if err != nil {
		//			Error(err)
		//			return
		//		}
	}
	Debug("finish received")

}
