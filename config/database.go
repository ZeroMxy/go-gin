package config

import "go-gin/core/env"

var Database = map[string]string{

	// Database
	"drive":           env.Get("database.drive", "mysql"),
	"host":            env.Get("database.host", "localhost"),
	"port":            env.Get("database.port", "3306"),
	"dbname":          env.Get("database.dbname", "basic"),
	"username":        env.Get("database.username", "root"),
	"password":        env.Get("database.password", "root"),
	"maxIdleConns":    env.Get("database.maxIdleConns", ""),
	"maxOpenConns":    env.Get("database.maxOpenConns", ""),
	"connMaxLifetime": env.Get("database.connMaxLifetime", ""),
}
