package config

import (
	"path/filepath"
	"runtime"

	"github.com/integrii/flaggy"
)

// Add flag vars to be referenced by main and other packages
var ConfigLocation string
var ConfigFile string

var Globalflags *flaggy.Parser = flaggy.NewParser("ipwatcher")

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

	Globalflags.Description = "A daemon to track your public IP address and report changes to a backend notification API."
	Globalflags.Version = Version + "\nGit commit: " + Commit + "\nGo version: " + Go + "\nOS: " + Os + "\nArchitecture: " + Arch

	Globalflags.String(&ConfigFile, "c", "config", "Define a custom configuration location for ipwatcher to parse/utilize.")

}
