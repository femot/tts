package main

import (
	"log"

	"github.com/femot/tts/pkg/twitch"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("Starting Shoe's personal TTS bot")
	twitch.Connect()
}
