package main

import (
	"net/http"
	"os"
	"toml"
)

const (
	OauthPrefix string = "/oauth/v1/"
)

var (
	mux = http.NewServeMux()
)

func main() {
	loadConfig()
	//  fmt.Println(regexp.Match("H.* ", []byte("Hello World!")))
	initDB()
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		Error(err)
	}
}

func loadConfig() {
	var err error
	var tomlPath string = "./output/server.toml"
	if _, err = toml.DecodeFile(tomlPath, &ServerConfig); err != nil {
		Error(err)
		os.Exit(-1)
	}
	Debug(ServerConfig)
}
