package audioplayer

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type AudioPlayer2 struct {
	streamer beep.StreamSeekCloser
	done     chan struct{}
}

func StartPlayer2(path string) (*AudioPlayer2, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open audio file: %v", err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("failed to decode audio file: %v", err)
	}

	done := make(chan struct{})

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/2))
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(done)
		streamer.Close()
	})))

	return &AudioPlayer2{done: done, streamer: streamer}, nil
}

func (a *AudioPlayer2) Stop() error {
	return a.streamer.Seek(a.streamer.Len())
}

func (a *AudioPlayer2) Done() <-chan struct{} {
	return a.done
}
