package config

var Database = map[string]interface{} {

	// Database
	"drive":           Env("database.drive", "mysql"),
	"host":            Env("database.host", "localhost"),
	"port":            Env("database.port", "3306"),
	"dbname":          Env("database.dbname", "basic"),
	"username":        Env("database.username", "root"),
	"password":        Env("database.password", "root"),
	"maxIdleConns":    Env("database.maxIdleConns", ""),
	"maxOpenConns":    Env("database.maxOpenConns", ""),
	"connMaxLifetime": Env("database.password", ""),
}
