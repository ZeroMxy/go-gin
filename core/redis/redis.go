package redis

import (
	"fmt"
	"go-gin/config"
	"sync"

	"github.com/go-redis/redis/v8"
)

type Redis struct{}

var (
	conn *redis.Client
	once sync.Once
)

func Conn () *redis.Client {
	
	once.Do(func () {
		new()
	})
	return conn
}

func new () {

	conn = redis.NewClient(&redis.Options {
		Addr:     fmt.Sprintf("%s:%d", config.Redis["host"], config.Redis["port"]),
		Password: config.Redis["password"].(string),
		DB:       config.Redis["dbname"].(int),
	})
	return
}
