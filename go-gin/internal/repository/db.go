package repository

import (
	"fmt"
	"go-api/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func InitDb() {
	dns := "root:root@tcp(localhost:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("MySQL DB 연결 실패")
	}
	if err := db.AutoMigrate(&model.Board{}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("MySQL DB 연결 성공")

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = db
}
