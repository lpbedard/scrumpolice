package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
	"github.com/scrumpolice/scrumpolice/bot"
	"github.com/scrumpolice/scrumpolice/scrum"
)

const header = "                                           _ _\n" +
	" ___  ___ _ __ _   _ _ __ ___  _ __   ___ | (_) ___ ___\n" +
	"/ __|/ __| '__| | | | '_ ` _ \\| '_ \\ / _ \\| | |/ __/ _ \\\n" +
	"\\__ \\ (__| |  | |_| | | | | | | |_) | (_) | | | (_|  __/\n" +
	"|___/\\___|_|   \\__,_|_| |_| |_| .__/ \\___/|_|_|\\___\\___|\n" +
	"                              |_|"

const Version = "0.7.1"

func main() {
	fmt.Println(header)
	fmt.Println("Version", Version)
	fmt.Println("")

	slackBotToken := os.Getenv("SCRUMPOLICE_SLACK_TOKEN")

	if slackBotToken == "" {
		log.Fatalln("slack bot token must be set in SCRUMPOLICE_SLACK_TOKEN")
	}

	configFile := "config.json"
	flag.StringVar(&configFile, "config", configFile, "The configuration file")
	flag.Parse()

	// Injection
	logger := log.New(os.Stderr, "", log.Lshortfile)
	configurationProvider := scrum.NewConfigWatcher(configFile)
	slackAPIClient := slack.New(slackBotToken)
	scrum := scrum.NewService(configurationProvider, slackAPIClient)

	// Create and run bot
	b := bot.New(slackAPIClient, logger, scrum)
	b.Run()
}
