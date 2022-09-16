package twitch

import (
	"fmt"

	t "github.com/gempir/go-twitch-irc/v3"
)

func Foo() {
	tok, err := LoadToken()
	if err != nil {
		panic(err)
	}

	client := t.NewClient("streamot", string(tok))
	client.OnPrivateMessage(func(msg t.PrivateMessage) {
		fmt.Println(msg.Message)
	})
	client.Join("ifvta")

	select {}
}
