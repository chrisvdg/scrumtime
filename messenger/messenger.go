package messenger

// Messenger defines the interface of an object that can send messages to a messages platform
type Messenger interface {
	SendMessage() error
	Platform() string
}
