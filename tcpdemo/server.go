package main

import (
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var nowCount int64
var lastCount int64
var count int
var mutex sync.RWMutex
var note int
var temp int64

func staticstic() {
	for {
		time.Sleep(time.Second)

		temp := atomic.LoadInt64(&nowCount)
		Debug("statistic this second :", temp-lastCount)
		lastCount = temp
	}
}

func main() {
	SetLogLevel("DEBUG")
	var tcpAddr *net.TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:7777")
	if err != nil {
		Error(err)
		return
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		Error(err)
		return
	}
	defer tcpListener.Close()
	go staticstic()
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			Error(err)
			continue
		}
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn *net.TCPConn) {
	defer func() {
		conn.Close()
	}()

	var msg string
	for i := 0; i < 2000; i++ {
		msg = msg + "a"
	}
	var rev string

	for i := 0; i < 2000; i++ {
		rev = rev + "1"
	}

	b := []byte(msg)
	lenth := 0
	wlen := 0
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		data := make([]byte, 2000)
		//		var err error
		lenth = 0
		for {
			templen, err := conn.Read(data[lenth:])
			lenth += templen
			if err != nil {
				Error(err)
				time.Sleep(time.Hour)
				return
			}
			if lenth == 2000 {
				break
			}
		}
		Info(string(data))
		if rev != string(data) {
			Error("not the same")
			time.Sleep(time.Hour)
		}
		wlen = 0
		for {
			templen, err := conn.Write(b[wlen:])
			wlen += templen
			if err != nil {
				Error(err)
				time.Sleep(time.Hour)
				return
			}
			if wlen == 2000 {
				break
			}
		}
		atomic.AddInt64(&nowCount, 1)
	}
}
