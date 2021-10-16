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

func init() {
	p := new(MyIPParser)
	config.RegisterConfig(p.Name(), nil, MyIPIsUsed, false)
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
	req.Header.Add("User-Agent", fmt.Sprintf("IPWatcher v%s", config.Version))

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
