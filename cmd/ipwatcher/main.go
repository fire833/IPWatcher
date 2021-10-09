package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fire833/ipwatcher/src/backend"
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
	backend.LoadNotifiers()
	go watcher.WatcherThread()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGILL, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	s := <-sig
	log.Printf("Killing ipwatcher: %v", s)
	os.Exit(0)
}
