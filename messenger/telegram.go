package messenger

import (
	"fmt"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// NewTelegramMessenger returns a new Telegram messenger
func NewTelegramMessenger(chatIDs []string, message, apikey string, verbose bool) (*TelegramMessenger, error) {
	client, err := tgbotapi.NewBotAPI(apikey)
	if err != nil {
		return nil, err
	}

	tm := new(TelegramMessenger)

	for _, chatID := range chatIDs {
		id, err := strconv.ParseInt(chatID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Telegram chat id %s is not valid", chatID)
		}
		tm.ChatIDs = append(tm.ChatIDs, id)
	}

	tm.Message = message
	tm.client = client
	tm.verbose = verbose

	return tm, nil
}

// TelegramMessenger represents a telegram messenger
type TelegramMessenger struct {
	ChatIDs []int64
	Message string
	client  *tgbotapi.BotAPI
	verbose bool
}

// SendMessage implements messenger.SendMessage
func (t *TelegramMessenger) SendMessage() error {
	for _, id := range t.ChatIDs {

		msg := tgbotapi.NewMessage(id, t.Message)
		msg.ParseMode = "markdown"
		_, err := t.client.Send(msg)
		if err != nil {
			return fmt.Errorf("Telegram messenger: Something went wrong sending message to %d: %s", id, err)
		}
	}

	return nil
}

// Platform implements messenger.Platform
func (*TelegramMessenger) Platform() string {
	return "Telegram"
}
