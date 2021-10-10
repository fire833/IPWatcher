package parsers

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

type IPInfo struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

type IPInfoParser struct {
	// body   []byte
	parsed *IPInfo
}

func (p *IPInfoParser) Name() string {
	return "ipinfo.io"
}

func (p *IPInfoParser) GetASN() string {
	asn := strings.SplitN(p.parsed.Org, " ", 2)
	return asn[0]
}

func (p *IPInfoParser) GetHostname() string {
	return p.parsed.Hostname
}

func (p *IPInfoParser) GetLocality() []string {
	return []string{p.parsed.City, p.parsed.Region, p.parsed.Country}
}

func (p *IPInfoParser) GetOrg() string {
	asn := strings.SplitN(p.parsed.Org, " ", 2)
	return asn[1]
}

func (p *IPInfoParser) GetLocation() []float64 {
	loc := strings.SplitN(p.parsed.Location, ",", 2)

	lat, _ := strconv.ParseFloat(loc[0], 32)
	long, _ := strconv.ParseFloat(loc[1], 32)

	return []float64{lat, long}
}

func (p *IPInfoParser) Get(ip net.IP) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.Add("Content-Type", "application/json")
	req.Header.SetMethod("GET")
	req.Header.Add("User-Agent", fmt.Sprintf("IPWatcher v%s", config.Version))

	if len(ip) == net.IPv4len {
		req.SetRequestURI(fmt.Sprintf("https://ipinfo.io/%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3]))

		if err := fasthttp.Do(req, resp); err != nil {
			log.Default().Printf("Error with acquiring public IP information from %s, error is: %v", p.Name(), err)
		}

		if err1 := json.Unmarshal(resp.Body(), p.parsed); err1 != nil {
			log.Default().Printf("Error with unmarshalling public IP info from %s, error is: %v", p.Name(), err1)
		}

	} else if len(ip) == net.IPv6len {
		req.SetRequestURI(fmt.Sprintf("https://ipinfo.io/%x%x:%x%x:%x%x:%x%x:%x%x:%x%x:%x%x:%x%x", ip[0], ip[1], ip[2], ip[3], ip[4], ip[5], ip[6], ip[7], ip[8], ip[9], ip[10], ip[11], ip[12], ip[13], ip[14], ip[15]))

		if err := fasthttp.Do(req, resp); err != nil {
			log.Default().Printf("Error with acquiring public IP information from %s, error is: %v", p.Name(), err)
		}

		if err1 := json.Unmarshal(resp.Body(), p.parsed); err1 != nil {
			log.Default().Printf("Error with unmarshalling public IP info from %s, error is: %v", p.Name(), err1)
		}
	}
}
