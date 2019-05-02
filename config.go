package main

import (
	"net/url"
	"strings"

	"github.com/kovetskiy/ko"
)

type Config struct {
	BaseURL  string `toml:"base_url" required:"true"`
	Username string `toml:"username" required:"true"`
	Password string `toml:"password" required:"true"`
}

func getConfig(path string) (Config, error) {
	config := Config{}

	err := ko.Load(path, &config)
	if err != nil {
		return config, err
	}

	config.BaseURL = strings.TrimSuffix(config.BaseURL, `/`)

	return config, err
}

func getConfigFromURI(rawuri string) (Config, error) {
	uri, err := url.Parse(rawuri)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	config.BaseURL = uri.Scheme + "://" + uri.Host + uri.Path

	auth := uri.User
	if auth != nil {
		config.Username = auth.Username()
		config.Password, _ = auth.Password()
	}

	return config, nil
}
