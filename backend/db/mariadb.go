package db

import (
	"fmt"
	"rota/conf"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql orm
	"github.com/name5566/leaf/log"
)

// RDBClient is opened orm object
var RDBClient *gorm.DB = nil

// InitRDB init the orm object
func InitRDB(m ...interface{}) {
	if RDBClient == nil {
		conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.Server.DBUser, conf.Server.DBPass, conf.Server.DBAddr, conf.Server.DBName)

		var err error
		for {
			RDBClient, err = gorm.Open("mysql", conn)
			if err != nil {
				log.Error("%v", err)
				time.Sleep(5 * time.Second)
			} else {
				log.Release("MariaDB connected")
				break
			}
		}

		RDBClient.AutoMigrate(m...)

		RDBClient.DB().SetMaxIdleConns(0)
		RDBClient.DB().SetMaxOpenConns(256)
		RDBClient.DB().SetConnMaxLifetime(time.Second * 16)
	}
}

// CloseRDB close orm object
func CloseRDB() {
	if RDBClient != nil {
		err := RDBClient.Close()
		if err != nil {
			log.Fatal("%v", err)
		}
	}
}
