package parsers

import (
	"log"
	"net"

	"github.com/valyala/fasthttp"
)

type WhatsMyIPAddrParser struct {
	body  []byte
	isV4  bool
	addr  string
	addrb net.IP
}

func (p *WhatsMyIPAddrParser) ParseIP() net.IP {
	return p.addrb
}

func (p *WhatsMyIPAddrParser) GetStringIP() string {
	return p.addr
}

func (p *WhatsMyIPAddrParser) Get() {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI("https://bot.whatismyipaddress.com")

	if err := fasthttp.Do(req, resp); err != nil {
		log.Default().Printf("Error with acquiring public IP from %s, error is: %v", "bot.whatismyipaddress.com", err)
	}

	p.body = resp.Body()
	p.addr = string(resp.Body())
	p.addrb = net.ParseIP(string(resp.Body()))
	if len(p.addrb) == 4 {
		p.isV4 = true
	} else {
		p.isV4 = false
	}

}

func (p *WhatsMyIPAddrParser) IsV4() bool {
	return p.isV4
}

func (p *WhatsMyIPAddrParser) Body() []byte {
	return p.body
}
