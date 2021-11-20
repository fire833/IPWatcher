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
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

var MyIPIsUsed bool = false

type MyIpResp struct {
	Success bool   `json:"success"`
	IP      string `json:"ip"`
	Type    string `json:"type"`
}

type MyIPParser struct {
	addr   string
	body   []byte
	isV4   bool
	parsed *MyIpResp
	addrb  net.IP
}

func (p *MyIPParser) Name() string {
	return "api.my-ip.io"
}

func (p *MyIPParser) ParseIP() net.IP {
	return p.addrb
}

func (p *MyIPParser) GetStringIP() string {
	return p.addr
}

func (p *MyIPParser) Get() error {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI("https://api.my-ip.io/ip.json")
	req.Header.SetMethod("GET")
	req.Header.Add("Content-Type", "application/json")
	req.Header.SetContentType("application/json")
	req.Header.SetUserAgent(fmt.Sprintf("IPWatcher v%s", config.Version))

	if err := fasthttp.Do(req, resp); err != nil {
		log.Default().Printf("Error with acquiring public IP from %s, error is: %v", p.Name(), err)
		return err
	}

	p.body = resp.Body()

	parse := &MyIpResp{}

	if err1 := json.Unmarshal(resp.Body(), parse); err1 != nil {
		log.Default().Printf("Error with unmarshalling public IP info from %s, error is: %v", p.Name(), err1)
		return err1
	}

	p.parsed = parse
	p.addr = p.parsed.IP
	p.addrb = net.ParseIP(p.parsed.IP)
	if len(p.addrb) == 4 {
		p.isV4 = true
	} else {
		p.isV4 = false
	}

	return nil
}

func (p *MyIPParser) IsV4() bool {
	return p.isV4
}

func (p *MyIPParser) Body() []byte {
	return p.body
}
