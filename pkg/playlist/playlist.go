package playlist

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/femot/tts/pkg/audioplayer"
	"github.com/femot/tts/pkg/tts"
)

type AudioPlayer interface {
	Stop() error
	Done() <-chan struct{}
}

type Playlist struct {
	queue   chan string
	mu      sync.Mutex
	current AudioPlayer
	stop    bool
}

func NewPlaylist(queueSize int) *Playlist {
	s := &Playlist{queue: make(chan string, queueSize)}
	go s.run()
	return s
}

func (p *Playlist) run() {
	for {
		if p.stop {
			return
		}
		select {
		case s := <-p.queue:
			buf, err := tts.SpeakFile(s)
			if err != nil {
				log.Printf("failed to get speak file: %s", err)
				continue
			}

			f, err := os.CreateTemp("", "speak-*.mp3")
			if err != nil {
				log.Printf("failed to create temp file: %s", err)
				continue
			}
			_, err = f.Write(buf)
			if err != nil {
				log.Printf("failed to write temp file: %s", err)
				f.Close()
				os.Remove(f.Name())
				continue
			}
			f.Close()

			player, err := audioplayer.StartPlayer2(f.Name())
			if err != nil {
				log.Printf("failed to start audio player: %s", err)
				continue
			}

			p.mu.Lock()
			p.current = player
			p.mu.Unlock()

			<-player.Done()
			if err := os.Remove(f.Name()); err != nil {
				log.Printf("failed to remove temp file: %s", err)
			}

			p.mu.Lock()
			p.current = nil
			p.mu.Unlock()
		}
	}
}

func (p *Playlist) QueueTTS(s string) error {
	select {
	case p.queue <- s:
		return nil
	default:
		log.Println("queue full")
		return errors.New("queue is full, try again later")
	}
}

func (p *Playlist) Skip() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.current != nil {
		log.Printf("skipping current TTS")
		p.current.Stop()
	}
}

func (p *Playlist) Stop() {
	fmt.Println("stopping playlist ...")
	p.stop = true
	p.Skip()
}
