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
