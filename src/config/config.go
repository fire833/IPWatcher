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
	// Define what resolver you want used to find your public IP. can either be my-ip or whatsmyip.
	IPresolver string `json:"resolver" yaml:"resolver"`
	// Define what info-gatherer you want to use to gather information about your new IP. currently only uses ipinfo.
	IPInfoGatherer string `json:"info_gatherer" yaml:"info_gatherer"`
	// Configuration for the pushover notification backend
	Pushover *PushoverConfig `json:"pushover,omitempty" yaml:"pushover,omitempty"`
	Discord  *DiscordConfig  `json:"discord,omitempty" yaml:"discord,omitempty"`
	Slack    *SlackConfig    `json:"slack,omitmepty" yaml:"slack,omitmepty"`
	Teams    *TeamsConfig    `json:"teams,omitempty" yaml:"teams,omitempty"`
	Webhook  *WebhookConfig  `json:"webhook,omitempty" yaml:"webhook,omitempty"`
}

type Webhook string

// Configuration for the pushover notification backend
type PushoverConfig struct {
	ApiKey string   `json:"api_key" yaml:"apiKey"`
	Users  []string `json:"users" yaml:"users"`
}

type SlackConfig struct {
	Webhooks []Webhook `json:"hooks" yaml:"hooks"`
}

type DiscordConfig struct {
	Webhooks []Webhook `json:"hooks" yaml:"hooks"`
}

type TeamsConfig struct {
	Webhooks []Webhook `json:"hooks" yamls:"hooks"`
}

type WebhookConfig struct {
	Webhooks []Webhook `json:"hooks" yaml:"hooks"`
}

type TelegramConfig struct {
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
