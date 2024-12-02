package config

import "time"

type RedisConfig struct {
	Host          string        // Redis server host
	Port          string        // Redis server port
	Password      string        // Redis server password
	DB            int           // Redis database index
	DialTimeout   time.Duration // Connection dial timeout
	ReadTimeout   time.Duration // Read timeout for commands
	WriteTimeout  time.Duration // Write timeout for commands
	PoolSize      int           // Connection pool size
	MinIdleConns  int           // Minimum idle connections in the pool
	MaxRetries    int           // Number of retries for commands
	TLS           bool          // Use TLS for Redis connection
	TLSCertFile   string        // Path to TLS certificate file
	TLSKeyFile    string        // Path to TLS key file
	TLSSkipVerify bool          // Skip TLS certificate verification
	Prefix        string
}

// BuildConfig applies default values for RedisConfig if fields are not set
func (c *RedisConfig) BuildConfig() {
	if c.Host == "" {
		c.Host = "localhost"
	}
	if c.Port == "" {
		c.Port = "6379"
	}
	if c.DialTimeout == 0 {
		c.DialTimeout = 5 * time.Second
	}
	if c.ReadTimeout == 0 {
		c.ReadTimeout = 3 * time.Second
	}
	if c.WriteTimeout == 0 {
		c.WriteTimeout = 3 * time.Second
	}
	if c.PoolSize == 0 {
		c.PoolSize = 10
	}
	if c.MinIdleConns == 0 {
		c.MinIdleConns = 2
	}
	if c.MaxRetries == 0 {
		c.MaxRetries = 3
	}
}
