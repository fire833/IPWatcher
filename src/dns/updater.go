package dns

import (
	"errors"
	"net"
)

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
