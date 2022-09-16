package audioplayer

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

const vlcPath = `C:\Program Files (x86)\VideoLAN\VLC\vlc.exe`

type AudioPlayer struct {
	cmd *exec.Cmd
}

func StartPlayer(path string) (*AudioPlayer, error) {
	p := AudioPlayer{}
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	p.cmd = exec.Command(vlcPath, "-I", "dummy", "--dummy-quiet", "--volume", "1", path, "vlc://quit")
	return &p, p.cmd.Start()
}

func (v AudioPlayer) Stop() error {
	if v.cmd.Process != nil {
		return v.cmd.Process.Kill()
	}
	return nil
}

func (v AudioPlayer) Done() <-chan struct{} {
	c := make(chan struct{})
	go func() {
		v.cmd.Wait()
		close(c)
	}()
	return c
}

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
