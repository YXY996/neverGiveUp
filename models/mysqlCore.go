package models

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Println("failed to read")
		os.Exit(1)
	}
	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	user := config.Section("mysql").Key("user").String()
	password := config.Section("mysql").Key("password").String()
	database := config.Section("mysql").Key("database").String()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)
	//打印sql
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{QueryFields: true})
	if err != nil {
		fmt.Printf("gorm conn failed %v\n", err)
	}
}
