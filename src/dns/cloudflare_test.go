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
