package main

import (
	"github.com/fire833/ipwatcher/src/config"
	"github.com/fire833/ipwatcher/src/flag"
	"github.com/fire833/ipwatcher/src/watcher"
	"github.com/integrii/flaggy"
)

func main() {

	flaggy.SetName("ipwatcher")
	flaggy.SetDescription("A daemon to track your public IP address and report changes to a backend notification API.")
	flaggy.SetVersion("")

	flaggy.String(&flag.ConfigFile, "c", "config", "Define a custom configuration location for ipwatcher to parse/utilize.")

	flaggy.Parse()

	config.LoadConfig()
	watcher.WatcherThread()

}
