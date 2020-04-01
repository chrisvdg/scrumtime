package messenger

import (
	"fmt"

	"github.com/chrisvdg/scrumtime/config"
	"github.com/nlopes/slack"
)

// NewSlackMessenger returns a new Slack messenger
func NewSlackMessenger(channel []string, message *config.Message, apikey string, verbose bool) (*SlackMessenger, error) {
	if apikey == "" {
		return nil, fmt.Errorf("no api key provided")
	}
	if len(channel) == 0 {
		return nil, fmt.Errorf("no channels (chat_ids) provided")
	}
	sm := new(SlackMessenger)
	sm.Message = message
	sm.Channels = channel
	sm.client = slack.New(apikey)
	sm.verbose = verbose

	if message.ExpireTime != "" {
		fmt.Println("Warning: Expire time is not supported by the Slack messenger.")
	}

	return sm, nil
}

// SlackMessenger represents a messenger for Slack
type SlackMessenger struct {
	Channels []string
	Message  *config.Message
	client   *slack.Client
	verbose  bool
}

// SendMessage implements messenger.SendMessage
func (s *SlackMessenger) SendMessage() error {
	for _, channel := range s.Channels {

		channelID, timestamp, err := s.client.PostMessage(
			channel,
			s.Message.Body,
			slack.PostMessageParameters{
				UnfurlLinks: s.Message.DisableLinkPreview,
				UnfurlMedia: s.Message.DisableLinkPreview,
			})

		if err != nil {
			err = fmt.Errorf("Slack messenger: Something went wrong sending a message (%s at %s): %s", channelID, timestamp, err)
			return err
		}
	}

	return nil
}

// Platform implements messenger.Platform
func (*SlackMessenger) Platform() string {
	return "Slack"
}
