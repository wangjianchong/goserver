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
	err = dbpool.db.Model(&Login{}).Where("username = ?", "www").Updates(map[string]interface{}{
		"username": "yyy",
		"pwd":      "yyy",
	}).Error
	if err != nil {
		Error(err)
	}
	var test []Login
	db := dbpool.db.Where("username = ?", "yyy")
	db.Find(&test)
	Debug(test)
	var test2 []Test
	db.Find(&test2)
	Debug(test2)

}
