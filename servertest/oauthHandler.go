package main

import "net/http"

func init() {
	mux.HandleFunc(OauthPrefix+"login", loginHandle)
}

func loginHandle(res http.ResponseWriter, req *http.Request) {

}
