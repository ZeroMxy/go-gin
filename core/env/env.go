package env

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

func Get (name, defaultValue string) string {
	
	return fmt.Sprint(env(name, defaultValue))
}

func env (name string, defaultValue interface {}) interface {} {

	env, err := ioutil.ReadFile("env.yaml")
	if err != nil {
		panic("Description Failed to read the configuration file")
	}

	var config map[string]interface {}
	if yaml.Unmarshal(env, &config); err != nil {
		panic("Configuration file parsing failed Procedure")
	}

	value := read(name, config)
	if value == "" {
		value = defaultValue
	}

	return value
}

func read (name string, configs map[string]interface {}) interface {} {

	names := strings.Split(name, ".")
	length := len(names)
	data := configs[names[0]]

	for i := 1; i < length; i++ {

		datas, ok := data.(map[string]interface {})
		if !ok {
			return ""
		}

		value, exist := datas[names[i]]
		if !exist {
			return ""
		}
		data = value
	}

	return data
}
