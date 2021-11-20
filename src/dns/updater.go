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
	"errors"
	"log"
	"net"
)

var DNSUpdaters []DNSUpdater

// Interface that all DNS Updater types need to match in order to be considered proper
// DNS updaters that can be called if needed.
type DNSUpdater interface {
	// Call to update the DNS of the current hostname with a new entry.
	UpdateDNS(OldHost string, NewEntry *DNSEntry) error
	// Returns the name of the backend updater.
	Name() string
}

var (
	TypeError     = errors.New("Incorrect DNS record type provided.")
	HostnameError = errors.New("Incorrect hostname provided.")
	TTLError      = errors.New("Invalid TTL parameter provided for this DNS entry.")
)

func RegisterDNSUpdater(d DNSUpdater) {
	DNSUpdaters = append(DNSUpdaters, d)
	log.Default().Printf("Successfully registered %s as a DNS updater.", d.Name())
	return
}

type DNSEntry struct {
	Type     string
	Hostname string
	Content  net.IP
	TTL      int
}

func (d *DNSEntry) setType(t string) error {
	if t == "A" || t == "AAAA" || t == "CNAME" || t == "TXT" {
		d.Type = t
		return nil
	} else {
		return TypeError
	}
}

func (d *DNSEntry) setTTL(ttl int) error {
	if ttl == 1 || ttl > 60 {
		d.TTL = ttl
		return nil
	} else {
		return TTLError
	}
}

func NewEntry(IP net.IP, Type string, Hostname string, TTL int) (*DNSEntry, error) {
	entry := &DNSEntry{
		Content:  IP,
		Hostname: Hostname,
	}

	if err := entry.setType(Type); err != nil {
		return entry, err
	}

	if err := entry.setTTL(TTL); err != nil {
		return entry, err
	}

	return entry, nil
}
