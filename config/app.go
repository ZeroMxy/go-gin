package config

import "go-gin/core/env"

var App = map[string]string{

	// Application
	"name": env.Get("app.name", "go-gin"),
	"host": env.Get("app.host", "localhost:3000"),
	"mode": env.Get("app.mode", "debug"),

	// Snowflake
	"workerId":  env.Get("snowflake.workerId", "0"),
	"startTime": env.Get("snowflake.startTime", "0"),
}
