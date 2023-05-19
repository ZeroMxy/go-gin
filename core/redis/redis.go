package redis

import (
	"go-gin/config"
	"strconv"
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

	dbname, err := strconv.Atoi(config.Redis["dbname"])
	if err != nil {
		dbname = 0
	}

	conn = redis.NewClient(&redis.Options {
		Addr:     config.Redis["host"] + ":" + config.Redis["port"],
		Password: config.Redis["password"],
		DB:       dbname,
	})
	return
}
