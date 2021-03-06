package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

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
		t := time.Now()
		fmt.Printf("Current date time: %s\n\n", t.Format("02/01/2006 15:04"))
	}
	c := cron.New()

	fmt.Println("Scheduling messages:")
	for name, scheduleCfg := range cfg.Messages {
		if len(scheduleCfg.Messengers) == 0 {
			if *verbose {
				fmt.Printf("\t%s does not contain messengers, skipped scheduling the message\n", name)
			}
			continue
		}
		job, err := NewScheduledMessage(name, scheduleCfg, cfg.Bots, *verbose)
		if err != nil {
			log.Fatal(err)
		}
		c.AddJob(scheduleCfg.Schedule, job)
		fmt.Printf("\t- %s\n", name)
	}

	c.Start()
	fmt.Println("\nMessages are scheduled.\nPress ctrl + c to stop and exit.")
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	c.Stop()
}
