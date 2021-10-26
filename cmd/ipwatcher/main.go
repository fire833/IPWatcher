package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fire833/ipwatcher/src/config"
	"github.com/fire833/ipwatcher/src/watcher"
)

func main() {

	config.Globalflags.Parse()

	config.LoadConfig()
	go watcher.WatcherThread()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGILL, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	s := <-sig
	log.Printf("Killing ipwatcher: %v", s)
	os.Exit(0)
}
