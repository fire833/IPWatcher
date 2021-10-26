package backend

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

var SC *SlackConfig
var SlackIsUsed bool = false

type SlackNotification struct {
	l *Limit
	e string
}

type SlackConfig struct {
	Webhooks []config.Webhook `json:"hooks"`
}

func (c *SlackConfig) UnmarshalConfig(input []byte) {
	if SlackIsUsed {
		SC = new(SlackConfig)
	} else {
		return
	}

	if err := json.Unmarshal(input, SC); err == nil {
		n := new(SlackNotification)
		RegisterNotifier(n)
	} else {
		return
	}
}

func init() {
	n := new(DiscordNotification)
	config.RegisterConfig(n.Name(), DC, DiscordIsUsed, false)
}

func (n *SlackNotification) Name() string {
	return "Slack"
}

func (n *SlackNotification) Send(msg *Message) error {

	for _, hook := range SC.Webhooks {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		// Setup the request
		req.SetRequestURI(string(hook))
		req.Header.SetMethod("POST")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("User-Agent", fmt.Sprintf("IPWatcher v%s", config.Version))

		if err := fasthttp.Do(req, resp); err != nil {
			log.Default().Printf("Error with sending %s notification: %v", n.Name(), err)
			return err
		}
	}

	return nil
}

func (n *SlackNotification) GetLimit() *Limit {
	return n.l
}

func (n *SlackNotification) Error() string {
	return n.e
}
