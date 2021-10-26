package backend

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

var TC *TeamsConfig
var TeamsIsUsed bool = false

type TeamsNotification struct {
	l *Limit
	e string
}

type TeamsConfig struct {
	Webhooks []config.Webhook `json:"hooks"`
}

func (c *TeamsConfig) UnmarshalConfig(input []byte) {
	if TeamsIsUsed {
		TC = new(TeamsConfig)
	} else {
		return
	}

	if err := json.Unmarshal(input, TC); err == nil {
		n := new(WebhookNotification)
		RegisterNotifier(n)
	} else {
		return
	}
}

func init() {
	n := new(TeamsNotification)
	config.RegisterConfig(n.Name(), DC, TeamsIsUsed, false)
}

func (n *TeamsNotification) Name() string {
	return "Teams"
}

func (n *TeamsNotification) Send(msg *Message) error {

	for _, hook := range TC.Webhooks {
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

func (n *TeamsNotification) GetLimit() *Limit {
	return n.l
}

func (n *TeamsNotification) Error() string {
	return n.e
}
