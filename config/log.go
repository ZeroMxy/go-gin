package config

import "go-gin/core/env"

var Log = map[string]string{

	// Log
	"path": env.Get("log.logic", "storage/log/logic.log"),
}
