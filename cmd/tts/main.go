package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/femot/tts/pkg/twitch"
)

var (
	channel = flag.String("channel", "", "Enter Twitch channel name.")
	redeem  = flag.String("redeem", "", "Enter name of the redeem. If left empty, bot will work with !tts <message> command instead of point redeems.")
	token   = flag.String("token", "", "OAuth2 Token. Needed to work with redeems.")
	debug   = flag.Bool("debug", false, "Enable debug logs.")
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	flag.Parse()
	if *channel == "" {
		fmt.Println("Missing channel name. Please provide one with the --name flag.")
		return
	}

	opts := twitch.Options{
		Debug:        *debug,
		RedeemEnable: *redeem != "",
		RedeemName:   *redeem,
		Token:        *token,
	}

	log.Printf("Starting TTS bot for channel: %s\n", *channel)
	twitch.Connect(*channel, opts)
}
