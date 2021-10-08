package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/fire833/ipwatcher/src/flag"
)

var GlobalConfig *IpWatcherConfig = new(IpWatcherConfig)

type IpWatcherConfig struct {
	// Define the time between polling for your current IP address.
	PollingInterval int `json:"polling_interval" yaml:"pollingInterval"`
	// Specify number of previous responses you want to keep cached in memory over time.
	CachedResponseBuffer int `json:"cache_response" yaml:"cacheResponse"`
	// Define what resolver you want used to find your public IP. can either be my-ip or whatsmyip.
	IPresolver string `json:"resolver" yaml:"resolver"`
	// Define the location you want to poll for your current IP address.
	// IPMirrorURL string `json:"ip_url" yaml:"ipUrl"`
}

func LoadConfig() {

	data, err := os.ReadFile(flag.ConfigFile)
	if err != nil {
		// Default just panic if the config can't be found for now.
		panic(err)
	}

	c := &IpWatcherConfig{}

	ext := filepath.Ext(flag.ConfigFile)

	switch {
	case ext == ".json":
		{
			err1 := json.Unmarshal(data, c)
			if err1 != nil {
				panic(err1)
			}

			GlobalConfig = c
		}
	case ext == ".yml" || ext == ".yaml":
		{
			err1 := yaml.Unmarshal(data, c)
			if err1 != nil {
				panic(err1)
			}

			GlobalConfig = c
		}
	default:
		{
			panic("Unsupported configuration extension.")
		}
	}
}
