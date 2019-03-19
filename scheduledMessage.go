package main

import (
	"fmt"
	"strings"

	"github.com/chrisvdg/scrumtime/config"
	"github.com/chrisvdg/scrumtime/messenger"
)

// NewScheduledMessage returns a schedules message that implements cron.Job
func NewScheduledMessage(name string, cfg *config.Schedule, msgrs map[string]*config.Messenger, verbose bool) (ScheduledMessage, error) {
	var sm ScheduledMessage
	sm.cfg = cfg
	sm.name = name
	sm.verbose = verbose
	messengers := make([]messenger.Messenger, 0)

	for _, m := range cfg.Messengers {
		msgr, ok := msgrs[m]
		if !ok {
			return sm, fmt.Errorf("messenger %s not found", m)
		}

		switch strings.ToLower(msgr.Platform) {
		case "slack":
			slackMsgr, err := messenger.NewSlackMessenger(msgr.ChatID, cfg.Message, msgr.APIKey, verbose)
			if err != nil {
				return sm, err
			}
			messengers = append(messengers, slackMsgr)
		case "telegram":
			telegramMsgr, err := messenger.NewTelegramMessenger(msgr.ChatID, cfg.Message, msgr.APIKey, verbose)
			if err != nil {
				return sm, err
			}
			messengers = append(messengers, telegramMsgr)
		default:
			return sm, fmt.Errorf("unrecognized platform: %s", msgr.Platform)
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
		fmt.Printf("Message %s triggered\n", s.name)
	}

	for _, m := range s.messengers {
		if s.verbose {
			fmt.Printf("Sending message %s on %s\n", s.name, m.Platform())
		}
		err := m.SendMessage()
		if err != nil {
			fmt.Printf("Something went wrong sending message %s on %s: %s\n", s.name, m.Platform(), err)
		}
	}

	if s.verbose {
		fmt.Printf("Message %s completed\n", s.name)
	}
}
