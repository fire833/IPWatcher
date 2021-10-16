package backend

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

var SC *SlackConfig

type SlackNotification struct {
	l *Limit
	e string
}

type SlackConfig struct {
	Webhooks []config.Webhook `json:"hooks"`
}

func (c *SlackConfig) UnmarshalConfig(input []byte) {

	json.Unmarshal(input, c)

}

func init() {

	SC = new(SlackConfig)
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

func (n *SlackNotification) Limit() *Limit {
	return n.l
}

func (n *SlackNotification) Error() string {
	return n.e
}
