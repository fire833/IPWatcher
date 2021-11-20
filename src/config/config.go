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

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var Version string = "unknown"
var Commit string = "unknown"
var Go string = runtime.Version()
var Os string = runtime.GOOS
var Arch string = runtime.GOARCH

var GlobalConfig *IpWatcherConfig = new(IpWatcherConfig)

var BackendConfigs map[string]BackendConfig = map[string]BackendConfig{}

type BackendConfig interface {
	UnmarshalConfig(input []byte)
}

type IpWatcherConfig struct {
	// Define the time between polling for your current IP address.
	PollingInterval int `json:"polling_interval" yaml:"pollingInterval"`
	// Define what resolver you want used to find your public IP. can either be my-ip or whatsmyip.
	IPresolver string `json:"resolver" yaml:"resolver"`
	// Define what info-gatherer you want to use to gather information about your new IP. currently only uses ipinfo.
	IPInfoGatherer string `json:"info_gatherer" yaml:"info_gatherer"`
}

func (c *IpWatcherConfig) UnmarshalConfig(input []byte) {
	if err := json.Unmarshal(input, c); err != nil {
		fmt.Println("Unable to load configuration.")
		os.Exit(1)
	}
}

type Webhook string

type WebhookConfig struct {
	Webhooks []Webhook `json:"hooks" yaml:"hooks"`
}

func init() {
	GlobalConfig = new(IpWatcherConfig)
	RegisterConfig("ipwatcher", GlobalConfig, true, true)
}

func LoadConfig() {

	data, err := os.ReadFile(ConfigFile)
	if err != nil {
		// Default just panic if the config can't be found for now.
		panic(err)
	}

	ext := filepath.Ext(ConfigFile)

	switch {
	case ext == ".json":
		{
			var c map[string]json.RawMessage

			err1 := json.Unmarshal(data, &c)
			if err1 != nil {
				panic(err1)
			}

			GlobalConfig.UnmarshalConfig(c["ipwatcher"])

			// Unmarshal all configurations with the respective backends so they can
			// then register with the daemon as being eligible/configured to perform their task.
			for name, conf := range BackendConfigs {
				conf.UnmarshalConfig(c[name])
			}

		}
	// case ext == ".yml" || ext == ".yaml":
	// 	{
	// 		var c map[string]yaml.RawMessage

	// 		err1 := yaml.Unmarshal(data, &c)
	// 		if err1 != nil {
	// 			panic(err1)
	// 		}

	// 		GlobalConfig = c
	// 	}
	default:
		{
			fmt.Println("Unsupported configuration file extension.")
			os.Exit(1)
		}
	}
}

func RegisterConfig(name string, conf BackendConfig, isUsed bool, isDefaultOn bool) {
	BackendConfigs[name] = conf
	Globalflags.Bool(&isUsed, fmt.Sprintf("%s", name), fmt.Sprintf("backend.%s", name), fmt.Sprintf("Call this flag to automatically enable the backend %s. (Default on?: %v)", name, isDefaultOn))
	return
}
