package linuxlens

import (
	"encoding/json"
	"os"
)

func LoadConfig(filename string) (*ServerConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var config ServerConfig
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
