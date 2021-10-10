package parsers

import (
	"fmt"
	"log"
	"net"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

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
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", fmt.Sprintf("IPWatcher v%s", config.Version))

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
