package database

import (
	"github.com/weiloon1234/gokit/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var GlobalDBConfig *config.DBConfig

func Init(config *config.DBConfig) error {
	var err error

	SetGlobalDBConfig(config)

	db, err = gorm.Open(mysql.Open(config.GetDSN()), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func SetGlobalDBConfig(config *config.DBConfig) {
	GlobalDBConfig = config
}

func GetGlobalDBConfig() *config.DBConfig {
	return GlobalDBConfig
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
