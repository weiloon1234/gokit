package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DBConfig struct {
	DSN string
}

var db *gorm.DB

func Init(config DBConfig) error {
	var err error
	dsn := config.DSN
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error while closing DB: %v", err)
			return
		}
		err = sqlDB.Close()
		if err != nil {
			return
		}
		log.Println("Database connection closed.")
	}
}
