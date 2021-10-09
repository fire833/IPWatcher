package flag

import (
	"path/filepath"
	"runtime"
)

// Add flag vars to be referenced by main and other packages
var ConfigLocation string
var ConfigFile string

var SlackBackend bool = false
var Discordbackend bool = false

func init() {
	switch runtime.GOOS {
	case "windows":
		{
			ConfigLocation = filepath.Join("C:\\", "Program Files (x86)", "IPWatcher")
			ConfigFile = filepath.Join(ConfigLocation, "config.json")
		}
	case "linux":
		{
			ConfigLocation = filepath.Join("/", "etc", "ipwatcher")
			ConfigFile = filepath.Join(ConfigLocation, "config.json")
		}
	}
}
