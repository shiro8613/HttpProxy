package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

func Dump(path string) error {
	config := Config{
		Listen: "127.0.0.1:8080",
		Location: map[string]Location{
			"example": {
				Path:      "/sample",
				ProxyPass: "http://127.0.0.1:3000",
			},
		},
	}

	b, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	file, err := os.Create(path)

	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(b)

	return nil
}