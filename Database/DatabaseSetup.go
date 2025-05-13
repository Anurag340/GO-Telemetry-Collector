package Database

import (
	"fmt"
	"github.com/Anurag340/telemetryCollector/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect() {
	dsn := "root@tcp(127.0.0.1:3306)/logdb?charset=utf8mb4&parseTime=True&loc=Local"


	db,err:=gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic("Failed to connect to database")
	}

	if err := db.AutoMigrate(&Models.LogEntry{}); err != nil {
		panic("Failed to migrate database")
	}

	DB = db
	fmt.Println("Connected to database")
}