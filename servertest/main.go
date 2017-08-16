package main

import (
	"net/http"
)

const (
	OauthPrefix string = "/oauth/v1/"
)

var (
	mux = http.NewServeMux()
)

func main1() {
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		Error(err)
	}
}
func main() {
	Debug(calulate(2, 3, div))
}

func calulate(n, m int, calu calType) int {
	return calu(n, m)
}

type calType func(int, int) int

func add(n, m int) int {
	return n + m
}
func div(n, m int) int {
	return n * m
}
