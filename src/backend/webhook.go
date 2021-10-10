package backend

import (
	"encoding/json"
	"log"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

type WebhookNotification struct {
}

func (n *WebhookNotification) Name() string {
	return "Webhook"
}

func (n *WebhookNotification) Send(msg *Message) error {

	for _, hook := range config.GlobalConfig.Webhook.Webhooks {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		// Setup the request
		req.SetRequestURI(string(hook))
		req.Header.SetMethod("POST")

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
	return nil
}
