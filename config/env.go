package config

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

func Env (name string, defaultValue interface{}) interface{} {

	env, err := ioutil.ReadFile("env.yaml")
	if err != nil {
		panic("Description Failed to read the configuration file")
	}

	var config map[string]interface{}
	if yaml.Unmarshal(env, &config); err != nil {
		panic("Configuration file parsing failed Procedure")
	}

	value := getValue(name, config)
	if value == "" {
		value = defaultValue
	}

	return value
}

// 递归获取值
func getValue (name string, configs map[string]interface{}) interface{} {

	names := strings.Split(name, ".")
	length := len(names)
	data := configs[names[0]]

	for i := 1; i < length; i++ {

		datas, ok := data.(map[string]interface{})
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
