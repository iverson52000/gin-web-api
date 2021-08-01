package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {

	// dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	var err error

	db, err = ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./db/sqliteDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("[db] db connected!")
	// log.Println("[db] db connected!")

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
