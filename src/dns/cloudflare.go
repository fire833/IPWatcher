package dns

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/valyala/fasthttp"
)

var CC *CloudflareConfig
var ClouflareIsUsed bool = true

type CloudflareDNSUpdater struct {
}

type CloudflareConfig struct {
	ZoneID string `json:"zone_id"`
	ApiKey string `json:"api_key"`
	// Specify if you want updated entries to use Cloudflare IP proxying for the
	// entry of to leave unproxied.
	ProxyEntries bool `json:"enable_proxy"`
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

func init() {
	n := new(CloudflareDNSUpdater)
	config.RegisterConfig(n.Name(), CC, ClouflareIsUsed, false)
}

func (c *CloudflareConfig) UnmarshalConfig(input []byte) {
	if ClouflareIsUsed {
		CC = new(CloudflareConfig)
	} else {
		return
	}

	if err := json.Unmarshal(input, CC); err == nil {
		n := new(CloudflareDNSUpdater)
		RegisterDNSUpdater(n)
	} else {
		return
	}
}

func (d *CloudflareDNSUpdater) UpdateDNS(OldHost string, NewEntry *DNSEntry) error {

	recordid, err := d.getOldRecord(OldHost)
	if err != nil {
		return err
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("PUT")
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("User-Agent", fmt.Sprintf("IPWatcher v%s", config.Version))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", CC.ApiKey)) // Add API key token to request header for authentication.
	// req.Header.Add("X-Auth-Email", CC.UserEmail) // Add email because apprently they just changed their API or something.

	req.SetBody([]byte(fmt.Sprintf(`{type":"%s","name":"%s","content":"%s","ttl":3600,"proxied":%v}`,
		NewEntry.Type, NewEntry.Hostname, NewEntry.Content.String(), CC.ProxyEntries)))
	req.SetRequestURI(fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s", CC.ZoneID, recordid))

	if err := fasthttp.Do(req, resp); err != nil {
		log.Default().Printf("Error updating DNS with provider %s for host %s: %v\n", d.Name(), OldHost, err)
		return err
	}

	r := &CloudflareDNSResp{}
	_ = json.Unmarshal(resp.Body(), r)

	if !r.Success {
		log.Default().Printf("Error updating DNS with provider %s for host %s: %v\n", d.Name(), OldHost, r.Errors)
		return errors.New(fmt.Sprintf("Error updating DNS with provider %s for host %s.\n", d.Name(), OldHost))
	}

	return nil
}

func (d *CloudflareDNSUpdater) getOldRecord(OldHost string) (string, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("User-Agent", fmt.Sprintf("IPWatcher v%s", config.Version))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", CC.ApiKey)) // Add API key token to request header for authentication.
	// req.Header.Add("X-Auth-Email", CC.UserEmail) // Add email because apprently they just changed their API or something.

	req.SetRequestURI(fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records/?name=%s", CC.ZoneID, OldHost))

	if err := fasthttp.Do(req, resp); err != nil {
		log.Default().Printf("Error updating DNS with provider %s for host %s: %v\n", d.Name(), OldHost, err)
		return "", err
	}

	r := &CloudflareDNSResp{}
	if err := json.Unmarshal(resp.Body(), r); err != nil {
		log.Default().Printf("Error updating DNS with provider %s for host %s: %v\n", d.Name(), OldHost, err)
		return "", err
	}

	for _, result := range r.Result {
		if result.Name == OldHost {
			return result.ID, nil
		}
	}

	return "", errors.New("No DNS entry found for the desired hostname.")
}

func (d *CloudflareDNSUpdater) Name() string { return "Cloudflare" }
