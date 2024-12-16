package database

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/weiloon1234/gokit/config"
)

var redisClient *redis.Client
var GlobalRedisConfig *config.RedisConfig

func InitRedis(config *config.RedisConfig) error {
	config.BuildConfig()

	// Configure TLS if enabled
	var tlsConfig *tls.Config
	if config.TLS {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: config.TLSSkipVerify,
		}
	}

	redisClient = redis.NewClient(&redis.Options{
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

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func SetGlobalRedisConfig(config *config.RedisConfig) {
	GlobalRedisConfig = config
}

func GetGlobalRedisConfig() *config.RedisConfig {
	return GlobalRedisConfig
}

func GetRedisClient() *redis.Client {
	return redisClient
}
