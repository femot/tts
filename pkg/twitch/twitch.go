package twitch

import (
	"log"
	"strings"

	"github.com/femot/tts/pkg/playlist"
	t "github.com/gempir/go-twitch-irc/v3"
)

type baruuk struct {
	client *t.Client
	player *playlist.Playlist
}

func Connect() {
	client := t.NewAnonymousClient()

	b := baruuk{client: client, player: playlist.NewPlaylist(5)}

	client.OnPrivateMessage(func(msg t.PrivateMessage) {
		log.Println(msg.User.Name+":", msg.Message)
		b.parseMessage(msg)
	})
	client.Join("kitsuxiu")
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

func (b *baruuk) parseMessage(msg t.PrivateMessage) {
	// Message too short or not a command
	if len(msg.Message) < 2 || msg.Message[0] != '!' {
		return
	}

	// Split message to get actual command
	split := strings.Split(msg.Message, " ")
	command := split[0][1:]

	switch command {
	case "tts":
		if len(split) > 1 {
			b.tts(strings.Join(split[1:], " "))
		}
	default:
		log.Printf("unkown command: %s", command)

	}
}

func (b *baruuk) tts(text string) {
	if err := b.player.QueueTTS(text); err != nil {
		log.Println(err)
	}
}
