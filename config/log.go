package config

var Log = map[string]interface{} {

	// Log
	"path": Env("log.logic", "storage/log/logic.log"),
}
