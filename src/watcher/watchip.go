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

package watcher

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/fire833/ipwatcher/src/backend"
	"github.com/fire833/ipwatcher/src/config"
	"github.com/fire833/ipwatcher/src/watcher/parsers"
)

// Cache of all previously stored IP responses for comparison.
var RemoteCache *iPResult = new(iPResult)
var LocalCache *[]iPResult = new([]iPResult)

type iPResult struct {
	IP   net.IP
	Time time.Time
}

// Thread function that will reach out and acquire IP at the given timestamp and do that forever,
// triggers further evaluation/things depending on if IP changed or not.
func WatcherThread() {
	for {

		errc := 0
		resetc := 0

		ret, err := AcquireIP()
		if err != nil {
			time.Sleep(time.Duration(time.Second * time.Duration(config.GlobalConfig.PollingInterval)))
			errc++
			if errc >= 5 {
				log.Default().Fatalf("Error with getting IP address from resolver %s, persistent error is: %v", parsers.IpResolver.Name(), err)
				// os.Exit(1)
			}
			continue
		}

		ip := &iPResult{
			IP:   ret,
			Time: time.Now(),
		}

		if RemoteCache.IP == nil {
			RemoteCache = ip
		}

		for i, b := range ip.IP {
			// If any of the bytes are different, assume the IP has changed.
			if b != RemoteCache.IP[i] {
				// Trigger a public IP address change event.
				parsers.IpInfoGatherer = new(parsers.IPInfoParser)
				parsers.IpInfoGatherer.Get(ip.IP)
				locale := parsers.IpInfoGatherer.GetLocality()

				content := fmt.Sprintf(`Your public IP address has now changed!\n
				It used to be %s, but has switched to %s. This new IP address is part of %s, 
				and the controlling organization is %s, located in %s, %s, %s.\n
				The hostname of this IP is %s.`,
					RemoteCache.IP, ip.IP, parsers.IpInfoGatherer.GetASN(), parsers.IpInfoGatherer.GetOrg(), locale[0], locale[1], locale[2], parsers.IpInfoGatherer.GetHostname())

				msg := &backend.Message{
					Title:     "IPWatcher Notification",
					Message:   content,
					Timestamp: time.Now().UnixNano(),
					Priority:  1,
				}

				for _, notifier := range backend.GlobalNotifiers {
					if err := notifier.Send(msg); err != nil {
						continue
					}
				}

				log.Default().Println(content)

				break
			}
		}

		resetc++

		if resetc >= 10 {
			errc = 0
		}

		RemoteCache = ip

		// Sleep before polling again given the specific polling interval.
		time.Sleep(time.Duration(time.Second * time.Duration(config.GlobalConfig.PollingInterval)))
	}
}

// func LocalWatcherThread() {
// 	for {

// 		ifaces, err := net.Interfaces()
// 		if err != nil {
// 			// Handle permissions error with this.
// 		}

// 		for _, iface := range ifaces {
// 			addresses, err := iface.Addrs()
// 			if err != nil {
// 				continue
// 			}

// 			for _, addr := range addresses {
// 				net.ParseIP(addr.String())
// 			}

// 		}

// 	}
// }

func AcquireIP() (Ip net.IP, err error) {
	parsers.IpResolver = new(parsers.WhatsMyIPAddrParser)
	err = parsers.IpResolver.Get()
	return parsers.IpResolver.ParseIP(), err
}
