package database

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var client *redis.Client
var GlobalRedisConfig *RedisConfig

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

func InitRedis(config *RedisConfig) error {
	config.BuildConfig()

	// Configure TLS if enabled
	var tlsConfig *tls.Config
	if config.TLS {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: config.TLSSkipVerify,
		}
	}

	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password:     config.Password,
		DB:           config.DB,
		DialTimeout:  config.DialTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
		TLSConfig:    tlsConfig,
	})

	// Test the connection with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func SetGlobalRedisConfig(config *RedisConfig) {
	GlobalRedisConfig = config
}

func GetGlobalRedisConfig() *RedisConfig {
	return GlobalRedisConfig
}

func GetClient() *redis.Client {
	return client
}
