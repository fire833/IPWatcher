package dns

import (
	"encoding/json"
	"net"
	"os"
	"testing"
)

type TestingConfig struct {
	Hostname string `json:"hostname"`
	UpdateIP string `json:"update"`
}

func TestCloudflareUpdater(t *testing.T) {
	file, _ := os.ReadFile("keyfile.json")

	CC = new(CloudflareConfig)
	lconf := &TestingConfig{}
	_ = json.Unmarshal(file, CC)
	_ = json.Unmarshal(file, lconf)

	cupdater := CloudflareDNSUpdater{}

	if cupdater.Name() != "Cloudflare" {
		t.Fail()
	}

	err := cupdater.UpdateDNS(lconf.Hostname, &DNSEntry{
		Type:     "A",
		Hostname: lconf.Hostname,
		Content:  net.ParseIP(lconf.UpdateIP),
		TTL:      3600,
	})

	if err != nil {
		t.Fail()
		t.Log("Error with updating IP with Cloudflare: " + err.Error())
	}

}
