package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=localhost user=admin password=123456 dbname=go-day-3-library-management port=54321 sslmode=disable TimeZone=Asia/Shanghai"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// d, err := gorm.Open("mysql", "root:rootPass/dbname")
	if err != nil {
		panic(err)

	}
	db = d

}

func GetDB() *gorm.DB {
	return db
}
