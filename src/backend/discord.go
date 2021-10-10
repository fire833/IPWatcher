package backend

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

type DiscordMsg struct {
	Content   string `json:"content"`
	Username  string `json:"username,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Tts       bool   `json:"tts,omitempty"`
	// File      io.Reader `json:"file,omitempty"`
}

type DiscordNotification struct {
	l *Limit
	e string
}

func (n *DiscordNotification) Name() string {
	return "Discord"
}

func (n *DiscordNotification) Send(msg *Message) error {

	dmsg := &DiscordMsg{
		Content:   msg.Message,
		Username:  msg.DeviceName,
		AvatarUrl: msg.URL,
		Tts:       false,
	}

	for _, hook := range config.GlobalConfig.Discord.Webhooks {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		// Setup the request

		req.SetRequestURI(string(hook))
		req.Header.SetMethod("POST")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("User-Agent", fmt.Sprintf("IPWatcher v%s", config.Version))

		data, _ := json.Marshal(dmsg)
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

func (n *DiscordNotification) Error() string {
	return n.e
}
