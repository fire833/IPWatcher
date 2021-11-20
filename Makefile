#
#	Copyright (C) 2021  Kendall Tauser
#
#	This program is free software; you can redistribute it and/or modify
#	it under the terms of the GNU General Public License as published by
#	the Free Software Foundation; either version 2 of the License, or
#	(at your option) any later version.
#
#	This program is distributed in the hope that it will be useful,
#	but WITHOUT ANY WARRANTY; without even the implied warranty of
#	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#	GNU General Public License for more details.
#
#	You should have received a copy of the GNU General Public License along
#	with this program; if not, write to the Free Software Foundation, Inc.,
#	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
#


VERSION = 0.01
DEFAULTGOOS = "linux"
DEFAULTGOARCH = "amd64"
COMMIT = $(git rev-parse HEAD)

all: 
	
	export GOOS=$(DEFAULTGOOS) && export GOARCH=$(DEFAULTGOARCH)
	
	@echo "Building binary from source..."
	# export commit=$(git rev-parse HEAD)

	go build -o bin/ipwatcher -ldflags "-X 'config.Version=$(VERSION)' -X 'config.Commit=$(git rev-parse HEAD)'" cmd/ipwatcher/main.go

#Run as sudo
install:	config

	install bin/ipwatcher /usr/bin/ipwatcher
	chmod 755 /usr/bin/ipwatcher
	install ipwatcher.service /usr/lib/systemd/system/ipwatcher.service
	chmod 644 /usr/lib/systemd/system/ipwatcher.service

	systemctl daemon-reload
	systemctl enable ipwatcher.service
	systemctl start ipwatcher.service

	@echo "Successfully installer IPWatcher onto system."

# Run as sudo
remove: 

	systemctl stop ipwatcher.service
	systemctl disable ipwatcher.service

	rm -rf /usr/bin/ipwatcher
	rm -rf /usr/lib/systemd/system/ipwatcher.service

	systemctl daemon-reload

# Run as sudo
config:

	mkdir /etc/ipwatcher && touch /etc/ipwatcher/config.json