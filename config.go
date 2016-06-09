package main

import "github.com/jinzhu/configor"
import "strings"

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

	config.BaseURL = strings.TrimSuffix(config.BaseURL, `/`)

	return config, err
}
