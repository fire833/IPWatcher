package backend

import (
	"fmt"
	"log"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

type SlackNotification struct {
	l *Limit
	e string
}

func (n *SlackNotification) Name() string {
	return "Slack"
}

func (n *SlackNotification) Send(msg *Message) error {

	for _, hook := range config.GlobalConfig.Slack.Webhooks {
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
