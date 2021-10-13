
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
install:

	install bin/ipwatcher /usr/bin/ipwatcher
	chmod 755 /usr/bin/ipwatcher
	install ipwatcher.service /usr/lib/systemd/system/ipwatcher.service

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