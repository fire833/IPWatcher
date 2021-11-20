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

package parsers

import (
	"fmt"
	"log"
	"net"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

var WhatsMyIPIsUsed bool = true

type WhatsMyIPAddrParser struct {
	body  []byte
	isV4  bool
	addr  string
	addrb net.IP
}

func (p *WhatsMyIPAddrParser) Name() string {
	return "bot.whatismyipaddress.com"
}

func (p *WhatsMyIPAddrParser) ParseIP() net.IP {
	return p.addrb
}

func (p *WhatsMyIPAddrParser) GetStringIP() string {
	return p.addr
}

func (p *WhatsMyIPAddrParser) Get() error {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI("https://bot.whatismyipaddress.com")
	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")
	req.Header.SetUserAgent(fmt.Sprintf("IPWatcher v%s", config.Version))

	if err := fasthttp.Do(req, resp); err != nil {
		log.Default().Printf("Error with acquiring public IP from %s, error is: %v", p.Name(), err)
		return err
	}

	p.body = resp.Body()
	p.addr = string(resp.Body())
	p.addrb = net.ParseIP(string(resp.Body()))
	if len(p.addrb) == 4 {
		p.isV4 = true
	} else {
		p.isV4 = false
	}

	return nil

}

func (p *WhatsMyIPAddrParser) IsV4() bool {
	return p.isV4
}

func (p *WhatsMyIPAddrParser) Body() []byte {
	return p.body
}
