package messenger

import (
	"fmt"

	"github.com/nlopes/slack"
)

// NewSlackMessenger returns a new Slack messenger
func NewSlackMessenger(channel, message, apikey string, verbose bool) (*SlackMessenger, error) {
	if apikey == "" {
		return nil, fmt.Errorf("no api key provided")
	}
	if channel == "" {
		return nil, fmt.Errorf("no channel (chat_id) provided")
	}
	sm := new(SlackMessenger)
	sm.Message = message
	sm.Channel = channel
	sm.client = slack.New(apikey)
	sm.verbose = verbose

	return sm, nil
}

// SlackMessenger represents a messenger for Slack
type SlackMessenger struct {
	Channel string
	Message string
	client  *slack.Client
	verbose bool
}

// SendMessage implements messenger.SendMessage
func (s *SlackMessenger) SendMessage() error {
	channelID, timestamp, err := s.client.PostMessage(
		s.Channel,
		s.Message,
		slack.PostMessageParameters{})

	if err != nil {
		err = fmt.Errorf("Slack messenger: Something went wrong sending a message (%s at %s): %s", channelID, timestamp, err)
	}

	return err
}

// Platform implements messenger.Platform
func (*SlackMessenger) Platform() string {
	return "Slack"
}
