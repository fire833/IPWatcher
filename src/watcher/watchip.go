package watcher

import (
	"encoding/json"
	"log"
	"net"
	"time"

	"github.com/fire833/ipwatcher/src/config"

	"github.com/valyala/fasthttp"
)

// Cache of all previously stored IP responses for comparison.
var LocalCache *ipCache = new(ipCache)

type MyIpResp struct {
	Success bool   `json:"success"`
	IP      string `json:"ip"`
	Type    string `json:"type"`
}

type iPResult struct {
	IP   *net.IPNet
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
			IP:   AcquireIP(),
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

func AcquireIP() (Ip *net.IPNet) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI("https://api.my-ip.io/ip.json")

	if err := fasthttp.Do(req, resp); err != nil {
		log.Default().Printf("Error with acquiring public IP from %s, error is: %v", "api.my-ip.io", err)
	}

	// var r *map[string]string
	r := &MyIpResp{}

	json.Unmarshal(resp.Body(), r)

	_, ipnet, err1 := net.ParseCIDR(r.IP)
	if err1 != nil {
		log.Default().Fatalf("Error with parsing CIDR from %s, error is: %v", "api.my-ip.io", err1)
	}

	return ipnet

}
