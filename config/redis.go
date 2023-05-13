package config

var Redis = map[string]interface{} {

	// Redis
	"host":     Env("redis.host", "localhost"),
	"port":     Env("redis.port", "6379"),
	"password": Env("redis.password", ""),
	"dbname":   Env("redis.dbname", "0"),
}
