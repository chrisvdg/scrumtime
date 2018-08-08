package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/chrisvdg/scrumtime/config"
	"github.com/nlopes/slack"
	"github.com/robfig/cron"
)

func init() {
	// Enable linenumbers in log messages
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var (
	cfg *config.App
	api *slack.Client
)

func main() {

	configPath := flag.String("c", "config.yaml", "YAML config file location")
	flag.Parse()

	var err error
	cfg, err = config.NewAppFromYaml(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	api = slack.New(cfg.APIKey)

	c := cron.New()
	c.AddFunc(cfg.Schedule, sendMessage)
	c.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func sendMessage() {
	params := slack.PostMessageParameters{}
	channelID, timestamp, err := api.PostMessage(cfg.Channel, cfg.Message, params)
	if err != nil {
		fmt.Printf("Something went wrong sending the message: %s\n", err)
	} else {
		fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
	}
}
