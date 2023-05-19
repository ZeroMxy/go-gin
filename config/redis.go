package config

import "go-gin/core/env"

var Redis = map[string]string{

	// Redis
	"host":     env.Get("redis.host", "localhost"),
	"port":     env.Get("redis.port", "6379"),
	"password": env.Get("redis.password", ""),
	"dbname":   env.Get("redis.dbname", "0"),
}
