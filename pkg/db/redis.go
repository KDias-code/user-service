package db

import "github.com/redis/go-redis/v9"

func ConnectRedis(host, port, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})

	return rdb
}
