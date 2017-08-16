package main

import "net/http"

func init() {
	mux.HandleFunc("/server/test1", test1)
}
func test1(writer http.ResponseWriter, request *http.Request) {
	Debug(request)
	num, err := writer.Write([]byte("hello world"))
	Debug([]byte("hello world"))
	if err != nil {
		Error(err)
	}
	Debug(num)
}
