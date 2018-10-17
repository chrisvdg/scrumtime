package main

import (
	"fmt"
	"strings"

	"github.com/chrisvdg/scrumtime/config"
	"github.com/chrisvdg/scrumtime/messenger"
)

// NewScheduledMessage returns a schedules message that implements cron.Job
func NewScheduledMessage(name string, cfg *config.Schedule, verbose bool) (ScheduledMessage, error) {
	var sm ScheduledMessage
	sm.cfg = cfg
	sm.name = name
	sm.verbose = verbose
	messengers := make([]messenger.Messenger, 0)

	for _, m := range cfg.Messengers {
		switch strings.ToLower(m.Platform) {
		case "slack":
			slackMsgr, err := messenger.NewSlackMessenger(m.ChatID, cfg.Message, m.APIKey)
			if err != nil {
				return sm, err
			}
			messengers = append(messengers, slackMsgr)

		case "telegram":
			telegramMsgr, err := messenger.NewTelegramMessenger(m.ChatID, cfg.Message, m.APIKey)
			if err != nil {
				return sm, err
			}
			messengers = append(messengers, telegramMsgr)

		default:
			return sm, fmt.Errorf("unrecognized platform: %s", m.Platform)

		}
	}

	sm.messengers = messengers

	return sm, nil
}

// ScheduledMessage represents a message that needs to be send on cron schedule
type ScheduledMessage struct {
	name       string
	cfg        *config.Schedule
	messengers []messenger.Messenger
	verbose    bool
}

// Run implements cron.Job.Run
func (s ScheduledMessage) Run() {
	if s.verbose {
		fmt.Printf("Job %s triggered\n", s.name)
	}

	for _, m := range s.messengers {
		err := m.SendMessage()
		if err != nil {
			fmt.Printf("Job: %s\nSomething went wrong sending the message: %s\n", s.name, err)
		}
	}

	if s.verbose {
		fmt.Printf("Job %s completed\n", s.name)
	}
}
