package backend

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

var TLC *TelegramConfig
var TelegramIsUsed bool = false

type TelegramNotification struct {
	l *Limit
	e string
}

type TelegramConfig struct {
	ApiKey string `json:"api_key"`
}

func (c *TelegramConfig) UnmarshalConfig(input []byte) {
	if TelegramIsUsed {
		TLC = new(TelegramConfig)
	} else {
		return
	}

	if err := json.Unmarshal(input, TLC); err == nil {
		n := new(TelegramNotification)
		RegisterNotifier(n)
	} else {
		return
	}
}

func init() {
	n := new(TelegramNotification)
	config.RegisterConfig(n.Name(), DC, TelegramIsUsed, false)
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
	req.Header.SetMethod("POST")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", fmt.Sprintf("IPWatcher v%s", config.Version))

	if err := fasthttp.Do(req, resp); err != nil {
		log.Default().Printf("Error with sending %s notification: %v", n.Name(), err)
		return err
	}

	return nil
}

func (n *TelegramNotification) GetLimit() *Limit {
	return n.l
}

func (n *TelegramNotification) Error() string {
	return n.e
}
