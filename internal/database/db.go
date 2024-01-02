package database

import (
	"ezyevent-api/internal/util"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDatabase() {
	var err error
	var dbString string = "host=" + util.GetVariableWith("DB_HOST") + " user=" + util.GetVariableWith("DB_USER") + " dbname=" + util.GetVariableWith("DB") + " password=" + util.GetVariableWith("DB_PWD") + "  port=" + util.GetVariableWith("DB_PORT")
	DBConn, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dbString,
	}))
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}
