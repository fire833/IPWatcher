package backend

import (
	"log"

	"github.com/valyala/fasthttp"
)

type TelegramNotification struct {
}

func (n *TelegramNotification) Name() string {
	return "Telegram"
}

func (n *TelegramNotification) Send(msg *Message) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// Setup the request

	if err := fasthttp.Do(req, resp); err != nil {
		log.Default().Printf("Error with sending %s notification: %v", n.Name(), err)
		return err
	}

	return nil
}

func (n *TelegramNotification) Limit() *Limit {
	return nil
}
