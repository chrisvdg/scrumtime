package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

// App represents the app's configuration
type App struct {
	Schedules map[string]*Schedule `yaml:"schedules"`
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

func truncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "..."
	}
	return bnoden
}

// Repr resturns a string representation of the app configuration
func (a *App) Repr() string {
	var s string

	for k, v := range a.Schedules {
		v := addIndent(v.Repr())
		s += fmt.Sprintf("%s:\n\t%s\v\n", k, v)
	}

	return s
}

// addIndent adds an indentation after each new line in the provided string
func addIndent(s string) string {
	return strings.Replace(s, "\n", "\n\t", -1)
}

// Schedule reprents the configuration of a single schedule
type Schedule struct {
	APIKey   string `yaml:"api_key"`
	Channel  string `yaml:"channel"`
	Schedule string `yaml:"schedule"`
	Message  string `yaml:"message"`
}

// Repr resturns a string representation of the Schedule config
func (s *Schedule) Repr() string {
	apikey := truncateString(s.APIKey, 20)
	return fmt.Sprintf("API Key: %s\nChannel: %s\nSchedule: %s\nMessage: %s",
		apikey,
		s.Channel,
		s.Schedule,
		s.Message)
}
