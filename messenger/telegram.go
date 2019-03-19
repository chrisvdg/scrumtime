package messenger

import (
	"fmt"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// NewTelegramMessenger returns a new Telegram messenger
func NewTelegramMessenger(chatID, message, apikey string, verbose bool) (*TelegramMessenger, error) {
	client, err := tgbotapi.NewBotAPI(apikey)
	if err != nil {
		return nil, err
	}

	tm := new(TelegramMessenger)

	tm.ChatID, err = strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Telegram chat id is not valid")
	}
	tm.Message = message
	tm.client = client
	tm.verbose = verbose

	return tm, nil
}

// TelegramMessenger represents a telegram messenger
type TelegramMessenger struct {
	ChatID  int64
	Message string
	client  *tgbotapi.BotAPI
	verbose bool
}

// SendMessage implements messenger.SendMessage
func (t *TelegramMessenger) SendMessage() error {
	msg := tgbotapi.NewMessage(t.ChatID, t.Message)
	msg.ParseMode = "markdown"
	_, err := t.client.Send(msg)

	return err
}

// Platform implements messenger.Platform
func (*TelegramMessenger) Platform() string {
	return "Telegram"
}
