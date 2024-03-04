package tts

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	voice  = "Brian"
	apiURL = "https://api.streamelements.com/kappa/v2/speech"
)

type speakResponse struct {
	Success  bool   `json:"success"`
	SpeakURL string `json:"speak_url"`
}

func SpeakFile(s string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s?voice=%s&text=%s", apiURL, voice, url.QueryEscape(s)))
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP response: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
