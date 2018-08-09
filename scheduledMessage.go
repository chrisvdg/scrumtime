package main

import (
	"fmt"

	"github.com/chrisvdg/scrumtime/config"
	"github.com/nlopes/slack"
)

// NewSchedulesMessage returns a schedules message that implements cron.Job
func NewSchedulesMessage(name string, cfg *config.Schedule) (ScheduledMessage, error) {
	var sm ScheduledMessage
	sm.cfg = cfg
	sm.api = slack.New(cfg.APIKey)
	sm.name = name

	return sm, nil
}

// ScheduledMessage represents a message that needs to be send on cron schedule
type ScheduledMessage struct {
	name string
	cfg  *config.Schedule
	api  *slack.Client
}

// Run implements cron.Job.Run()
func (s ScheduledMessage) Run() {
	channelID, timestamp, err := s.api.PostMessage(
		s.cfg.Channel,
		s.cfg.Message,
		slack.PostMessageParameters{})

	if err != nil {
		fmt.Printf("Job: %s\nSomething went wrong sending the message: %s\n", s.name, err)
	} else {
		fmt.Printf("Job: %s\nMessage successfully sent to channel %s at %s\n", s.name, channelID, timestamp)
	}
}
