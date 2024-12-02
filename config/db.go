package config

import "fmt"

type DBConfig struct {
	Driver       string
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
	if c.Driver == "" {
		c.Driver = "mysql"
	}

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
	if c.DSN == "" {
		// Format the DSN differently if the password is empty
		if c.Password == "" {
			c.DSN = fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
				c.Username,
				c.Host,
				c.Port,
				c.DatabaseName,
				c.Charset,
				c.ParseTime,
				c.TimeZone,
			)
		} else {
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
	}

	return c.DSN
}
