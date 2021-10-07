package backend

import (
	"github.com/gregdel/pushover"
)

type PushoverMsg struct {
}

func SendPushoverMsg(msg *pushover.Message) (resp *pushover.Response, err error) {
	app := pushover.New("")                                // TODO add support for getting tokens from file.
	return app.SendMessage(msg, pushover.NewRecipient("")) // Add users to send to.
}
