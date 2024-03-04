package tts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	voice  = "Brian"
	apiURL = "https://streamlabs.com/polly/speak"
)

type speakResponse struct {
	Success  bool   `json:"success"`
	SpeakURL string `json:"speak_url"`
}

func SpeakFile(s string) ([]byte, error) {
	resp, err := http.PostForm(apiURL, url.Values{
		"voice": {voice},
		"text":  {s},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP response: %d", resp.StatusCode)
	}

	var sr speakResponse
	err = json.NewDecoder(resp.Body).Decode(&sr)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve speak file: %w", err)
	}

	resp, err = http.Get(sr.SpeakURL)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	return b, err
}
