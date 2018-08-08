package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// App represents the app's configuration
type App struct {
	APIKey   string `yaml:"api_key"`
	Channel  string `yaml:"channel"`
	Schedule string `yaml:"schedule"`
	Message  string `yaml:"message"`
}

// NewAppFromYaml returns the app data from provided yaml file
func NewAppFromYaml(path string) (*App, error) {
	app := new(App)

	// Get data from provided yaml file
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		errstr := fmt.Sprintf("Error while reading config file: %v", err)
		return nil, errors.New(errstr)
	}
	err = yaml.Unmarshal(raw, app)

	return app, nil
}
