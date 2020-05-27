package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDb() {

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	// apps_port := os.Getenv("PORT")
	var err error
	db, err = gorm.Open("mysql", db_user+":"+db_pass+"@tcp("+db_host+":"+db_port+")/mydb")
	// db, err = gorm.Open("mysql", "root:A123b456c@tcp(localhost:3306)/mydb")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Customer_Info{})
}
