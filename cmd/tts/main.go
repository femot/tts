package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/femot/tts/pkg/twitch"
)

var (
	channel = flag.String("channel", "", "Enter Twitch channel name.")
	debug   = flag.Bool("debug", false, "Enable debug logs.")
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	flag.Parse()
	if *channel == "" {
		fmt.Println("Missing channel name. Please provide one with the --name flag.")
		return
	}
	log.Printf("Starting TTS bot for channel: %s\n", *channel)
	twitch.Connect(*channel, *debug)
}
