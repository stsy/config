package config

import (
	"encoding/json"
	"io/ioutil"
)

// Open config file return a map
func Open(path string) (config map[string]string, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = map[string]string{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return
	}

	return
}
