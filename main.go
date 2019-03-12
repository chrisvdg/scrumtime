package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/chrisvdg/scrumtime/config"
	"github.com/robfig/cron"
)

func init() {
	// Enable linenumbers in log messages
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	configPath := flag.String("c", "config.yaml", "YAML config file location")
	verbose := flag.Bool("v", false, "verbose output")
	flag.Parse()

	var err error
	cfg, err := config.NewAppFromYaml(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	if *verbose {
		fmt.Println("Verbose output set.")
	}
	c := cron.New()

	for name, scheduleCfg := range cfg.Schedules {
		job, err := NewScheduledMessage(name, scheduleCfg, cfg.Messengers, *verbose)
		if err != nil {
			log.Fatal(err)
		}
		c.AddJob(scheduleCfg.Schedule, job)
	}

	c.Start()
	fmt.Println("Messages are scheduled.\nPress ctrl + c to stop and exit.")
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	c.Stop()
}
