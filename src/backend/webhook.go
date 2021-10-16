package backend

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

var WC *WebhookConfig
var WebhookIsUsed bool = false

type WebhookNotification struct {
	l *Limit
	e string
}

type WebhookConfig struct {
	Webhooks []config.Webhook `json:"hooks"`
}

func (c *WebhookConfig) UnmarshalConfig(input []byte) {

	json.Unmarshal(input, c)

}

func init() {

	WC = new(WebhookConfig)
	n := new(WebhookNotification)
	config.RegisterConfig(n.Name(), DC, WebhookIsUsed, false)

}

func (n *WebhookNotification) Name() string {
	return "Webhook"
}

func (n *WebhookNotification) Send(msg *Message) error {

	for _, hook := range WC.Webhooks {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		// Setup the request
		req.SetRequestURI(string(hook))
		req.Header.SetMethod("POST")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("User-Agent", fmt.Sprintf("IPWatcher v%s", config.Version))

		data, _ := json.Marshal(msg)
		req.SetBody(data)

		if err := fasthttp.Do(req, resp); err != nil {
			log.Default().Printf("Error with sending %s notification: %v", n.Name(), err)
			continue
		}
	}

	return nil
}

func (n *WebhookNotification) Limit() *Limit {
	return n.l
}

func (n *WebhookNotification) Error() string {
	return n.e
}
