
VERSION == 0.01
DEFAULTGOOS == "linux"
DEFAULTGOARCH == "amd64"

all: 
	
	# export GOOS=${DEFAULTGOOS} && export GOARCH=${DEFAULTGOARCH}
	
	echo "Building binary from source..."
	commit="$(git rev-parse HEAD)"

	go build -o bin/ipwatcher -ldflags "-X 'config.Version=${VERSION}' -X 'config.Commit=${commit}' -X 'config.AllowExec=true'" cmd/ipwatcher/main.go

#Run as sudo
install:
	
	install cmd/ipwatcher/ipwatcher /usr/bin/ipwatcher
	chmod 755 /usr/bin/ipwatcher
	install ipwatcher.service /usr/lib/systemd/system/ipwatcher.service

	systemctl daemon-reload
	systemctl enable ipwatcher.service
	systemctl start ipwatcher.service

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