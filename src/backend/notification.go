package backend

import (
	"log"
	"time"
)

var GlobalNotifiers []Notifier

type Message struct {
	Title   string `json:"title" yaml:"title"`
	Message string `json:"message" yaml:"message"`

	Priority    int           `json:"priority" yaml:"priority"`
	URL         string        `json:"url" yaml:"url"`
	URLTitle    string        `json:"utl_title" yaml:"urlTitle"`
	Timestamp   int64         `json:"timestamp" yaml:"timstamp"`
	Retry       time.Duration `json:"retry" yaml:"retry"`
	Expire      time.Duration `json:"expire" yaml:"expire"`
	CallbackURL string        `json:"callback" yaml:"callback"`
	DeviceName  string        `json:"device" yaml:"device"`
	Sound       string        `json:"sound" yaml:"sound"`
}

type Limit struct {
	MessagesRemaining int `json:"remaining_total" yaml:"remainingTotal"`
	MessagesLeftWeek  int `json:"remaining_week" yaml:"remainingWeek"`
	MessagesLeftMonth int `json:"remaining_month" yaml:"remainingMonth"`
}

type Notifier interface {
	// Use this method to send a message using the specified backend.
	Send(msg *Message) error
	// Use this method to get the request limits left for that notification method, if any.
	GetLimit() *Limit
	// Return the name of the notifier backend
	Name() string
	// Return an error of the notifier backend
	Error() string
}

func RegisterNotifier(in Notifier) {
	GlobalNotifiers = append(GlobalNotifiers, in)
	log.Default().Printf("Successfully registered %s notifier with IPwatcher.", in.Name())
	return
}
