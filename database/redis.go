package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var client *redis.Client

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	Prefix   string
}

func InitRedis(config RedisConfig) error {
	client = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
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

func GetClient() *redis.Client {
	return client
}
