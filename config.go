package main

import (
	"strings"

	"github.com/jinzhu/configor"
)

type config struct {
	BaseURL  string `toml:"base_url" required:"true"`
	Username string `toml:"username" required:"true"`
	Password string `toml:"password" required:"true"`
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
