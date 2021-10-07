package backend

import "time"

type GlobalNotifiers *[]Notifier

type Message struct {
	Title   string `json:"title" yaml:"title"`
	Message string `json:"message" yaml:"message"`

	Priority    int
	URL         string
	URLTitle    string
	Timestamp   int64
	Retry       time.Duration
	Expire      time.Duration
	CallbackURL string
	DeviceName  string
	Sound       string
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
}
