package db

import (
	"fmt"
	"os"
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
		conn := fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("ROTA_DBUSER"), os.Getenv("ROTA_DBPASS"), conf.Server.DBName)

		var err error
		RDBClient, err = gorm.Open("mysql", conn)
		if err != nil {
			log.Fatal("%v", err)
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
