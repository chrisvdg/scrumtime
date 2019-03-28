package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

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

	if err != nil {
		return nil, err
	}

	err = app.Validate()
	if err != nil {
		return nil, err
	}

	return app, nil
}

// App represents the app's configuration
type App struct {
	Bots     map[string]*Bot     `yaml:"bots"`
	Messages map[string]*Message `yaml:"messages"`
}

// Validate validates an app config
func (a *App) Validate() error {
	if len(a.Messages) == 0 {
		return fmt.Errorf("config file doesn't contain Messages")
	}

	if len(a.Bots) == 0 {
		return fmt.Errorf("config file doesn't contain messengers")
	}

	for msgName, msg := range a.Messages {
		for _, msgr := range msg.Messengers {
			if _, ok := a.Bots[msgr.Bot]; !ok {
				return fmt.Errorf("bot %s not defined", msgr)
			}
		}

		if msg.Body == "" {
			return fmt.Errorf("message %s does not contain a body", msgName)
		}
	}

	return nil
}

// Message represents the configuration of a single Messaged message
type Message struct {
	Messengers []Messenger `yaml:"messengers"`
	Schedule   string      `yaml:"schedule"`
	Body       string      `yaml:"body"`
}

// Messenger represents a messenger of a message
type Messenger struct {
	Bot     string   `yaml:"bot"`
	ChatIDs []string `yaml:"chat_ids"`
}

// Bot represents a bot config
type Bot struct {
	Platform string `yaml:"platform"`
	APIKey   string `yaml:"api_key"`
}
