package backend

import (
	"log"
	"os/exec"
	"runtime"

	"github.com/fire833/ipwatcher/src/config"
)

var OSIsUsed bool = false

type OSNotification struct {
	l *Limit
	e string
}

func init() {
	n := new(OSNotification)
	config.RegisterConfig(n.Name(), nil, OSIsUsed, false)
}

func (n *OSNotification) Name() string { return "OS-Notification" }

func (n *OSNotification) Send(msg *Message) error {
	switch runtime.GOOS {
	case "linux":
		{
			path, err := exec.LookPath("notify-send")
			if err != nil {
				log.Default().Printf("Error with sending %s, error finding 'notify-send': %v", n.Name(), err)
				return err
			}

			err1 := exec.Command(path, msg.DeviceName, msg.Message).Run()
			if err1 != nil {
				return err1
			}
		}
	case "windows":
		{
			// TODO get notifications working in windows
		}
	}

	return nil
}

func (n *OSNotification) GetLimit() *Limit {
	return n.l
}

func (n *OSNotification) Error() string {
	return n.e
}
