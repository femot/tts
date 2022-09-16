package twitch

import (
	"fmt"
	"io/ioutil"
)

const TOKEN_FILE = "token"

type Token string

func LoadToken() (Token, error) {
	p, err := ioutil.ReadFile(TOKEN_FILE)
	if err != nil {
		return "", fmt.Errorf("failed to load token: %w", err)
	}

	return Token(p), nil
}

func (t Token) Save() error {
	return ioutil.WriteFile(TOKEN_FILE, []byte(t), 0601)
}

func GetTokenFromTwitchYes() (Token, error) {
	// send idiot to twitch

	// get it back ???

	// ahhh ffs

	return "", nil
}
