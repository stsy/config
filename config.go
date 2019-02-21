package config

import (
	"encoding/json"
	"io/ioutil"
)

// Open config file return a map
func Open(path string) (config map[string]interface{}, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return
	}

	return
}
