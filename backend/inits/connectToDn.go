package inits

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}