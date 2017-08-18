package main

import (
	"regexp"
)

const (
	PHONEPATTERN = `^1[3|5|7|8|][\d]{9}$`
)

//正则表达式
var (
	phoneReg *regexp.Regexp = regexp.MustCompile(PHONEPATTERN)
)

var ServerConfig struct {
	Mysql MysqlConn
}

type MysqlConn struct {
	Username string
	Password string
	Addr     string
	DBName   string
}
