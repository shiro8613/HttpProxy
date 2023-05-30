package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

func Load(path string) (Config, error){
	b, err := os.ReadFile(path);
	if err != nil {
		return Config{}, err
	}

	config := &Config{}

	yaml.Unmarshal(b, config)

	return *config, nil
}
