package backend

import (
	"encoding/json"
	"log"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

type DiscordNotification struct {
	l *Limit
}

func (n *DiscordNotification) Name() string {
	return "Discord"
}

func (n *DiscordNotification) Send(msg *Message) error {

	for _, hook := range config.GlobalConfig.Discord.Webhooks {
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
			return err
		}
	}

	return nil
}

func (n *DiscordNotification) Limit() *Limit {
	return n.l
}
