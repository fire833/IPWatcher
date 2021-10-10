package backend

import (
	"log"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

type TeamsNotification struct {
}

func (n *TeamsNotification) Name() string {
	return "Teams"
}

func (n *TeamsNotification) Send(msg *Message) error {

	for _, hook := range config.GlobalConfig.Teams.Webhooks {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		// Setup the request
		req.SetRequestURI(string(hook))
		req.Header.SetMethod("POST")

		if err := fasthttp.Do(req, resp); err != nil {
			log.Default().Printf("Error with sending %s notification: %v", n.Name(), err)
			return err
		}
	}

	return nil
}

func (n *TeamsNotification) Limit() *Limit {
	return nil
}
