package watcher

import (
	"net"
	"time"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/fire833/ipwatcher/src/watcher/parsers"
)

// Cache of all previously stored IP responses for comparison.
var LocalCache *ipCache = new(ipCache)

type iPResult struct {
	IP   net.IP
	Time time.Time
}

type ipCache struct {
	Cache []*iPResult
}

// Thread function that will reach out and acquire IP at the given timestamp and do that forever,
// triggers further evaluation/things depending on if IP changed or not.
func WatcherThread() {
	for {
		ip := &iPResult{
			IP:   AcquireIP(&parsers.MyIPParser{}),
			Time: time.Now(),
		}
		if len(LocalCache.Cache)+1 > config.GlobalConfig.CachedResponseBuffer {
			// If the buffer is full, then rmeove the first element,
			// shift all other elements down by one, and append the new value to the bottom.

			newCache := &ipCache{}

			for i, oldip := range LocalCache.Cache {
				if i == 0 {
					continue
				}
				// Append each but the first element to a new array so there is room for the newest element to be appended
				// and keep within the user-specified buffer size.
				newCache.Cache = append(newCache.Cache, oldip)
			}

			newCache.Cache = append(newCache.Cache, ip)
			// Set the current global buffer to the newly created/rotated buffer.
			LocalCache = newCache

		} else {
			LocalCache.Cache = append(LocalCache.Cache, ip)
		}

		// Sleep before polling again given the specific polling interval.
		time.Sleep(time.Duration(time.Second * time.Duration(config.GlobalConfig.PollingInterval)))
	}
}

func AcquireIP(p parsers.IpParserLocator) (Ip net.IP) {

	switch config.GlobalConfig.IPresolver {
	case "whatsmyip":
		{
			resolv := &parsers.WhatsMyIPAddrParser{}
			resolv.Get()
			return resolv.ParseIP()
		}
	default:
		{
			resolv := &parsers.MyIPParser{}
			resolv.Get()
			return resolv.ParseIP()
		}
	}
}
