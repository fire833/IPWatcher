/*
*	Copyright (C) 2021  Kendall Tauser
*
*	This program is free software; you can redistribute it and/or modify
*	it under the terms of the GNU General Public License as published by
*	the Free Software Foundation; either version 2 of the License, or
*	(at your option) any later version.
*
*	This program is distributed in the hope that it will be useful,
*	but WITHOUT ANY WARRANTY; without even the implied warranty of
*	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*	GNU General Public License for more details.
*
*	You should have received a copy of the GNU General Public License along
*	with this program; if not, write to the Free Software Foundation, Inc.,
*	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package backend

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

var DC *DiscordConfig
var DiscordIsUsed bool = false

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

type DiscordConfig struct {
	Webhooks []config.Webhook `json:"hooks"`
}

func (c *DiscordConfig) UnmarshalConfig(input []byte) {

	json.Unmarshal(input, c)

}

func init() {

	DC = new(DiscordConfig)
	n := new(DiscordNotification)
	config.RegisterConfig(n.Name(), DC, DiscordIsUsed, false)

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

	for _, hook := range DC.Webhooks {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		// Setup the request

		req.SetRequestURI(string(hook))
		req.Header.SetMethod("POST")
		req.Header.SetContentType("application/json")
		req.Header.SetUserAgent(fmt.Sprintf("IPWatcher v%s", config.Version))

		data, _ := json.Marshal(dmsg)
		req.SetBody(data)

		if err := fasthttp.Do(req, resp); err != nil {
			log.Default().Printf("Error with sending %s notification: %v", n.Name(), err)
			return err
		}
	}

	return nil
}

func (n *DiscordNotification) GetLimit() *Limit {
	return n.l
}

func (n *DiscordNotification) Error() string {
	return n.e
}
