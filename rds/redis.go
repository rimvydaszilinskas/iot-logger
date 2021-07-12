package rds

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type RedisConnection struct {
	redis *redis.Client
}

func GetRedisConnection() (*RedisConnection, error) {
	port, err := getRedisPort()

	if err != nil {
		return nil, fmt.Errorf("error retrieving redis host - %s", err)
	}

	db, err := getRedisDB()

	if err != nil {
		return nil, fmt.Errorf("error retrieving redis db - %s", err)
	}

	host := getRedisHost()
	host = fmt.Sprintf("%s:%d", host, port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: getRedisPassword(),
		DB:       db,
	})
	return &RedisConnection{
		redis: rdb,
	}, nil
}

func getRedisHost() string {
	host := os.Getenv("REDIS_HOST")
	if len(host) == 0 {
		return "127.0.0.1"
	}
	return host
}

func getRedisPassword() string {
	return os.Getenv("REDIS_PASSWORD")
}

func getRedisPort() (int, error) {
	port := os.Getenv("REDIS_PORT")

	if len(port) == 0 {
		return 6379, nil
	}
	return strconv.Atoi(port)
}

func getRedisDB() (int, error) {
	port := os.Getenv("REDIS_DB")

	if len(port) == 0 {
		return 6379, nil
	}
	return strconv.Atoi(port)
}
