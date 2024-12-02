package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DBConfig struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
	Charset      string
	ParseTime    bool
	TimeZone     string
	MaxIdleConns int
	MaxOpenConns int
	SSLMode      string
	ConfigBuilt  bool
	DSN          string
}

func (c *DBConfig) BuildConfig() {
	// Set charset default
	if c.Charset == "" {
		c.Charset = "utf8mb4"
	}

	// Set ParseTime default
	if !c.ParseTime {
		c.ParseTime = true
	}

	// Set TimeZone default
	if c.TimeZone == "" {
		c.TimeZone = "Local"
	}

	// Set connection pool defaults
	if c.MaxIdleConns == 0 {
		c.MaxIdleConns = 10
	}

	if c.MaxOpenConns == 0 {
		c.MaxOpenConns = 100
	}

	// Set SSLMode default
	if c.SSLMode == "" {
		c.SSLMode = "disable"
	}

	c.ConfigBuilt = true
}

func (c *DBConfig) GetDSN() string {
	if c.ConfigBuilt != true {
		c.BuildConfig()
	}

	if c.DSN == "" {
		c.DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
			c.Username,
			c.Password,
			c.Host,
			c.Port,
			c.DatabaseName,
			c.Charset,
			c.ParseTime,
			c.TimeZone,
		)
	}

	return c.DSN
}

var db *gorm.DB
var GlobalDBConfig *DBConfig

func Init(config *DBConfig) error {
	var err error

	config.BuildConfig()
	SetGlobalDBConfig(config)

	db, err = gorm.Open(mysql.Open(config.GetDSN()), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func SetGlobalDBConfig(config *DBConfig) {
	GlobalDBConfig = config
}

func GetGlobalDBConfig() *DBConfig {
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
