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
	// go watcher.LocalWatcherThread()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGILL, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	s := <-sig
	log.Printf("Killing ipwatcher: %v", s)
	os.Exit(0)
}
