package dns

import (
	"fmt"
	"time"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

type CloudflareDNSUpdater struct {
}

type CloudflareDNSResp struct {
	Success  bool            `json:"success"`
	Errors   []interface{}   `json:"errors"`
	Messages []interface{}   `json:"messages"`
	Result   []CloudFlareDNS `json:"result"`
}

type CloudFlareDNS struct {
	ID         string    `json:"id"`
	Type       string    `json:"type"`
	Name       string    `json:"name"`
	Content    string    `json:"content"`
	Proxiable  bool      `json:"proxiable"`
	Proxied    bool      `json:"proxied"`
	TTL        int       `json:"ttl"`
	Locked     bool      `json:"locked"`
	ZoneID     string    `json:"zone_id"`
	ZoneName   string    `json:"zone_name"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
	Data       struct {
	} `json:"data"`
	Meta struct {
		AutoAdded bool   `json:"auto_added"`
		Source    string `json:"source"`
	} `json:"meta"`
}

func (d *CloudflareDNSUpdater) UpdateDNS(OldHost string, NewEntry *DNSEntry) error {

	recordid := d.getOldRecord(OldHost)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.Add("Content-Type", "application/json")
	req.Header.SetMethod("PUT")
	req.Header.Add("User-Agent", fmt.Sprintf("IPWatcher v%s", config.Version))
	req.Header.Add("Authentication", fmt.Sprintf("Bearer %s", "")) // Add API key token to request header for authentication.

	req.SetRequestURI(fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", "", recordid))

	return nil
}

func (d *CloudflareDNSUpdater) getOldRecord(OldHost string) string {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.Add("Content-Type", "application/json")
	req.Header.SetMethod("GET")
	req.Header.Add("User-Agent", fmt.Sprintf("IPWatcher v%s", config.Version))
	req.Header.Add("Authentication", fmt.Sprintf("Bearer %s", "")) // Add API key token to request header for authentication.

	req.SetRequestURI(fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/?name=%s", "", OldHost))

	return "" // TODO return the Entry ID for modification with the other request.
}

func (d *CloudflareDNSUpdater) Name() string { return "Cloudflare" }
