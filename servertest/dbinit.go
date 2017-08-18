package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DBPool struct {
	db *gorm.DB
}

func (dbpool *DBPool) newDB(dnsRaw, dnsStr, dbName string) {
	db, err := sql.Open("mysql", dnsRaw)
	if err != nil {
		Error(err)
		os.Exit(-1)
	}
	sqlStr := fmt.Sprintf(`create database if not exists %s default character set utf8`, dbName)
	if _, err := db.Exec(sqlStr); err != nil {
		Error(err)
		os.Exit(-1)
	}
	dbpool.db, err = gorm.Open("mysql", dnsStr)
	if err != nil {
		Error(err)
		os.Exit(-1)
	}
	dbpool.db.DB().SetMaxOpenConns(500)
	dbpool.db.DB().SetMaxIdleConns(100)
}

var (
	dbpool DBPool
)

func initDB() {
	dnsRaw := fmt.Sprintf(`%s:%s@tcp(%s)/?charset=utf8&parseTime=True&loc=Local`, ServerConfig.Mysql.Username, ServerConfig.Mysql.Password, ServerConfig.Mysql.Addr)
	dnsStr := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local`, ServerConfig.Mysql.Username, ServerConfig.Mysql.Password, ServerConfig.Mysql.Addr, ServerConfig.Mysql.DBName)
	dbpool.newDB(dnsRaw, dnsStr, ServerConfig.Mysql.DBName)
	dbpool.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Login{})
	dbpool.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Music{})
	dbpool.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Test{})
	dbpool.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Contact{}, &Addr{})

}
