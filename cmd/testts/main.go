package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/femot/tts/pkg/playlist"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	p := playlist.NewPlaylist(2)
	log.SetFlags(0)

	fmt.Println("Enter TTS to test:")
	for {
		fmt.Print("> ")
		foo, _ := reader.ReadString('\n')
		p.Skip()
		fmt.Printf("Length: %d\n", len(foo))
		p.QueueTTS(foo)
	}
}
