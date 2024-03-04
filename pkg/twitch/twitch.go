package twitch

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

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

type Options struct {
	RedeemEnable bool
	RedeemName   string
	Token        string
	Debug        bool
}

func Connect(channel string, opts Options) {
	twitch.PubSub()
	client := t.NewAnonymousClient()
	mh := &messageHandler{client: client, player: playlist.NewPlaylist(5)}
	if !opts.RedeemEnable {
		client.OnPrivateMessage(func(msg t.PrivateMessage) {
			mh.parseMessage(msg)
		})
	}
	client.OnRoomStateMessage(func(msg t.RoomStateMessage) {
		mh.mu.Lock()
		defer mh.mu.Unlock()
		if mh.pubsub != nil {
			// Probably can ignore this.
			log.Println("room state changed after pubsub initialized, ignoring message")
			return
		}
		topic := fmt.Sprintf("channel-points-channel-v1.%s", msg.RoomID)
		mh.pubsub = twitch.PubSub()
		if opts.RedeemEnable {
			mh.pubsub.OnShardMessage(func(i int, s string, b []byte) {
				var r redeem
				if err := json.Unmarshal(b, &r); err != nil {
					log.Println(err)
					return
				}
				if r.Data.Redemption.Reward.Title == opts.RedeemName {
					log.Printf("TTS redeemed (%s): [%s] %s", r.Data.Redemption.Reward.Title, r.Data.Redemption.User.DisplayName, r.Data.Redemption.UserInput)
					mh.player.QueueTTS(r.Data.Redemption.UserInput)
				} else if opts.Debug {
					log.Printf("Redeem received: %s; Does not match TTS redeem (%s)", r.Data.Redemption.Reward.Title, opts.RedeemName)
				}
			})
		}
		if err := mh.pubsub.ListenWithAuth(opts.Token, topic); err != nil {
			log.Printf("failed to listen to topic (%s): %s\n", topic, err)
			mh.pubsub.Close()
			mh.pubsub = nil
		}
	})

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

type redeem struct {
	Type string `json:"type"`
	Data struct {
		Timestamp  time.Time `json:"timestamp"`
		Redemption struct {
			ID   string `json:"id"`
			User struct {
				ID          string `json:"id"`
				Login       string `json:"login"`
				DisplayName string `json:"display_name"`
			} `json:"user"`
			ChannelID  string    `json:"channel_id"`
			RedeemedAt time.Time `json:"redeemed_at"`
			Reward     struct {
				ID                  string `json:"id"`
				ChannelID           string `json:"channel_id"`
				Title               string `json:"title"`
				Prompt              string `json:"prompt"`
				Cost                int    `json:"cost"`
				IsUserInputRequired bool   `json:"is_user_input_required"`
				IsSubOnly           bool   `json:"is_sub_only"`
				Image               any    `json:"image"`
				DefaultImage        struct {
					URL1X string `json:"url_1x"`
					URL2X string `json:"url_2x"`
					URL4X string `json:"url_4x"`
				} `json:"default_image"`
				BackgroundColor string `json:"background_color"`
				IsEnabled       bool   `json:"is_enabled"`
				IsPaused        bool   `json:"is_paused"`
				IsInStock       bool   `json:"is_in_stock"`
				MaxPerStream    struct {
					IsEnabled    bool `json:"is_enabled"`
					MaxPerStream int  `json:"max_per_stream"`
				} `json:"max_per_stream"`
				ShouldRedemptionsSkipRequestQueue bool      `json:"should_redemptions_skip_request_queue"`
				TemplateID                        any       `json:"template_id"`
				UpdatedForIndicatorAt             time.Time `json:"updated_for_indicator_at"`
				MaxPerUserPerStream               struct {
					IsEnabled           bool `json:"is_enabled"`
					MaxPerUserPerStream int  `json:"max_per_user_per_stream"`
				} `json:"max_per_user_per_stream"`
				GlobalCooldown struct {
					IsEnabled             bool `json:"is_enabled"`
					GlobalCooldownSeconds int  `json:"global_cooldown_seconds"`
				} `json:"global_cooldown"`
				RedemptionsRedeemedCurrentStream any `json:"redemptions_redeemed_current_stream"`
				CooldownExpiresAt                any `json:"cooldown_expires_at"`
			} `json:"reward"`
			UserInput string `json:"user_input"`
			Status    string `json:"status"`
		} `json:"redemption"`
	} `json:"data"`
}
