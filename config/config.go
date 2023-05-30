package config

import (
	"fmt"
	"os"
)

func LoadAndCreate(path string) (Config, error) {
	config, err := Load(path)
	if os.IsNotExist(err) {
		fmt.Println("FileNotFound")
		err := Dump(path)
		if err != nil {
			return Config{}, err
		}
	} else if err != nil {
		return Config{}, err
	}

	return config, err
}