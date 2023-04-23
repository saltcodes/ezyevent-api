package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDatabase() {
	var err error
	DBConn, err = gorm.Open(postgres.New(postgres.Config{
		//DSN: "host=34.134.195.7 user=thanos dbname=ezyevents password=Ep0KmT1cvYU4xAAb  port=6432",
		DSN: "host=localhost user=johnson dbname=ezyevents password=c4dl3#54  port=5432",
	}))
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}
