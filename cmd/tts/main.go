package main

import (
	"log"

	"github.com/femot/tts/pkg/api"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	api.Start()
}
