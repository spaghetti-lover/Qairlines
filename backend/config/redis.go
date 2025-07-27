package config

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr     string
	Username string
	Password string
	DB       int
}

func NewRedisClient() *redis.Client {
	config, err := LoadConfig(".")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}
	redisDB, err := strconv.ParseInt(config.RedisDB, 10, 64)
	if err != nil {
		panic("failed to parse Redis DB: " + err.Error())
	}
	cfg := RedisConfig{
		Addr:     config.RedisAddress,
		Username: config.RedisUsername,
		Password: config.RedisPassword,
		DB:       int(redisDB),
	}

	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Username:     cfg.Username,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     10,
		MinIdleConns: 5,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully")
	return client
}
