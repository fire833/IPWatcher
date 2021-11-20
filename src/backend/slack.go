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

var SC *SlackConfig
var SlackIsUsed bool = false

type SlackNotification struct {
	l *Limit
	e string
}

type SlackConfig struct {
	Webhooks []config.Webhook `json:"hooks"`
}

func (c *SlackConfig) UnmarshalConfig(input []byte) {
	if SlackIsUsed {
		SC = new(SlackConfig)
	} else {
		return
	}

	if err := json.Unmarshal(input, SC); err == nil {
		n := new(SlackNotification)
		RegisterNotifier(n)
	} else {
		return
	}
}

func init() {
	n := new(DiscordNotification)
	config.RegisterConfig(n.Name(), DC, DiscordIsUsed, false)
}

func (n *SlackNotification) Name() string {
	return "Slack"
}

func (n *SlackNotification) Send(msg *Message) error {

	for _, hook := range SC.Webhooks {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		// Setup the request
		req.SetRequestURI(string(hook))
		req.Header.SetMethod("POST")
		req.Header.SetContentType("application/json")
		req.Header.SetUserAgent(fmt.Sprintf("IPWatcher v%s", config.Version))

		if err := fasthttp.Do(req, resp); err != nil {
			log.Default().Printf("Error with sending %s notification: %v", n.Name(), err)
			return err
		}
	}

	return nil
}

func (n *SlackNotification) GetLimit() *Limit {
	return n.l
}

func (n *SlackNotification) Error() string {
	return n.e
}
