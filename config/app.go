package config

var App = map[string]interface{} {

	// Application
	"name": Env("app.name", "go-gin"),
	"host": Env("app.host", "localhost:3000"),
	"mode": Env("app.mode", "debug"),

	// Snowflake
	"workerId":  Env("snowflake.workerId", "0"),
	"startTime": Env("snowflake.startTime", "0"),
}
