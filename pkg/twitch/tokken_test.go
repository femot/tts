package twitch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	tok := Token("kouks")
	err := tok.Save()
	assert.NoError(t, err)
	assert.FileExists(t, "token")
}

func TestLoad(t *testing.T) {
	tok, err := LoadToken()
	assert.NoError(t, err)
	assert.Equal(t, Token("kouks"), tok)
}
