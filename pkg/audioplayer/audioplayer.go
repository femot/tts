package audioplayer

import (
	"os"
	"os/exec"
)

const vlcPath = `C:\Program Files\VideoLAN\VLC\vlc.exe`

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
