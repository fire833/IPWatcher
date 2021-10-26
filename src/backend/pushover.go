package backend

import (
	"log"
	"time"

	"github.com/gregdel/pushover"
)

var PushoverIsUsed bool = false
var PLC *PushoverConfig

type PushoverNotification struct {
	// Limit
	l *Limit
	e string
}

// Configuration for the pushover notification backend
type PushoverConfig struct {
	ApiKey string   `json:"api_key" yaml:"apiKey"`
	Users  []string `json:"users" yaml:"users"`
}

func init() {

}

func (n *PushoverNotification) Name() string {
	return "Pushover"
}

func (n *PushoverNotification) Send(msg *Message) error {
	// Move from generic message to the pushover specific message.
	pmsg := &pushover.Message{
		Message:     msg.Message,
		Title:       msg.Title,
		Priority:    msg.Priority,
		URL:         msg.URL,
		URLTitle:    msg.URLTitle,
		Timestamp:   msg.Timestamp,
		Retry:       msg.Retry,
		Expire:      msg.Expire,
		CallbackURL: msg.CallbackURL,
		DeviceName:  msg.DeviceName,
		Sound:       msg.Sound,
	}

	app := pushover.New(PLC.ApiKey) // TODO add support for getting tokens from file.

	for i, user := range PLC.Users {
		if resp, err := app.SendMessage(pmsg, pushover.NewRecipient(user)); err != nil {
			log.Default().Printf("Error with sending %s notification: %v", n.Name(), err)
		} else {

			if i == len(PLC.Users)-1 {
				lim := &Limit{
					MessagesRemaining: resp.Limit.Remaining,
					MessagesLeftWeek:  resp.Limit.Remaining,
					MessagesLeftMonth: resp.Limit.Remaining,
				}

				n.l = lim
			}

			// Wait 1 second per request to be friendly to the API.
			time.Sleep(time.Second)
		}
	}

	return nil
}

func (n *PushoverNotification) GetLimit() *Limit {
	return n.l
}

func (n *PushoverNotification) Error() string {
	return n.e
}
