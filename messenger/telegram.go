package messenger

import (
	"fmt"
	"strconv"
	"time"

	"github.com/chrisvdg/scrumtime/config"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// NewTelegramMessenger returns a new Telegram messenger
func NewTelegramMessenger(chatIDs []string, message *config.Message, apikey string, verbose bool) (*TelegramMessenger, error) {
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
	Message *config.Message
	client  *tgbotapi.BotAPI
	verbose bool
}

// SendMessage implements messenger.SendMessage
func (t *TelegramMessenger) SendMessage() error {
	for _, id := range t.ChatIDs {

		msg := tgbotapi.NewMessage(id, t.Message.Body)
		msg.ParseMode = "markdown"
		message, err := t.client.Send(msg)
		if err != nil {
			return fmt.Errorf("Telegram messenger: Something went wrong sending message to %d: %s", id, err)
		}
		if t.Message.ExpireTime > 0 {
			go t.deleteMessage(t.Message.ExpireTime, id, message.MessageID)
		}
	}

	return nil
}

// deleteMessage delete message from a given chat
func (t *TelegramMessenger) deleteMessage(delay int, chatID int64, messageID int) error {
	time.Sleep(time.Duration(delay) * time.Second)
	msg := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := t.client.Send(msg)
	if err != nil {
		return fmt.Errorf("Telegram messenger: Something went wrong sending message to %d: %s", chatID, err)
	}

	return nil
}

// Platform implements messenger.Platform
func (*TelegramMessenger) Platform() string {
	return "Telegram"
}
