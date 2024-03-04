package twitch

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/pubsub"
	"github.com/femot/tts/pkg/playlist"
	t "github.com/gempir/go-twitch-irc/v3"
)

type messageHandler struct {
	mu     sync.Mutex
	client *t.Client
	pubsub *pubsub.Client
	player *playlist.Playlist
}

func Connect(channel string, debug bool) {
	twitch.PubSub()
	client := t.NewAnonymousClient()
	b := &messageHandler{client: client, player: playlist.NewPlaylist(5)}
	client.OnPrivateMessage(func(msg t.PrivateMessage) {
		if debug {
			log.Printf("OnPrivateMessage: %s", msg.Message)
		}
		b.parseMessage(msg)
	})
	client.OnRoomStateMessage(func(msg t.RoomStateMessage) {
		b.mu.Lock()
		defer b.mu.Unlock()
		if b.pubsub != nil {
			// Probably can ignore this.
			log.Println("room state changed after pubsub initialized, ignoring message")
			return
		}
		topic := fmt.Sprintf("channel-points-channel-v1.%s", msg.RoomID)
		b.pubsub = twitch.PubSub()
		b.pubsub.OnShardMessage(func(i int, s string, b []byte) {
			log.Printf("i: %d; s: %s; len(b): %d", i, s, len(b))
		})
		if err := b.pubsub.Listen(topic); err != nil {
			log.Printf("failed to listen to topic (%s): %s\n", topic, err)
			b.pubsub = nil
		}

	})

	// if debug {
	// 	debugCallbacks(client)
	// }

	client.Join(channel)
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

func debugCallbacks(client *t.Client) {
	client.OnClearChatMessage(func(msg t.ClearChatMessage) { log.Printf("OnClearChatMessage: %+v", msg) })
	client.OnClearMessage(func(msg t.ClearMessage) { log.Printf("ClearMessage: %+v", msg) })
	client.OnGlobalUserStateMessage(func(msg t.GlobalUserStateMessage) { log.Printf("GlobalUserStateMessage: %+v", msg) })
	client.OnNamesMessage(func(msg t.NamesMessage) { log.Printf("NamesMessage: %+v", msg) })
	client.OnPingMessage(func(msg t.PingMessage) { log.Printf("PingMessage: %+v", msg) })
	client.OnPongMessage(func(msg t.PongMessage) { log.Printf("PongMessage: %+v", msg) })
	client.OnReconnectMessage(func(msg t.ReconnectMessage) { log.Printf("ReconnectMessage: %+v", msg) })
	client.OnRoomStateMessage(func(msg t.RoomStateMessage) { log.Printf("RoomStateMessage: %+v", msg) })
	client.OnUnsetMessage(func(msg t.RawMessage) { log.Printf("OnUnsetMessage: %+v", msg) })
	client.OnUserJoinMessage(func(msg t.UserJoinMessage) { log.Printf("OnUserJoinMessage: %+v", msg) })
	client.OnUserNoticeMessage(func(msg t.UserNoticeMessage) { log.Printf("OnUserNoticeMessage: %+v", msg) })
	client.OnUserPartMessage(func(msg t.UserPartMessage) { log.Printf("OnUserPartMessage: %+v", msg) })
	client.OnUserStateMessage(func(msg t.UserStateMessage) { log.Printf("OnUserStateMessage: %+v", msg) })
}

func (b *messageHandler) parseMessage(msg t.PrivateMessage) {
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
			b.player.QueueTTS(strings.Join(split[1:], " "))
		}
	case "skiptts":
		if msg.User.Badges["moderator"] == 1 || msg.User.Badges["broadcaster"] == 1 {
			b.player.Skip()
		}
	default:
		log.Printf("unkown command: %s", command)
	}
}
