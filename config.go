package main

import "github.com/jinzhu/configor"

type config struct {
	BaseURL  string `toml:"base_url" required:"true"`
	Username string `required:"true"`
	Password string `required:"true"`
}

func getConfig(path string) (config, error) {
	config := config{}

	err := configor.Load(&config, path)
	if err != nil {
		return config, err
	}

	return config, err
}
