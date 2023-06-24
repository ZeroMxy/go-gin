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

	// Email
	"smtpHost":     env.Get("email.host", "smtp.example.com"),
	"smtpPort":     env.Get("email.port", "587"),
	"smtpUsername": env.Get("email.username", "your-email@example.com"),
	"smtpPassword": env.Get("email.password", "your-password"),
}
