package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/femot/tts/pkg/playlist"
)

type api struct {
	list *playlist.Playlist
}

func (a api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "skip") {
		a.list.Skip()
		w.WriteHeader(http.StatusOK)
		return
	}

	query := r.URL.Query()
	if text, ok := query["text"]; ok {
		log.Printf("TTS request received for: %s", text[0])

		err := a.list.QueueTTS(text[0])
		if err != nil {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, err)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func Start() {
	a := api{list: playlist.NewPlaylist(6)}
	defer a.list.Stop()
	http.ListenAndServe("127.0.0.1:7777", a)
}
