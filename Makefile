
DEFAULTGOOS = "linux"
DEFAULTGOARCH = "amd64"

all: 
	go build -ldflags "" cmd/ipwatcher/main.go -o bin/ipwatcher

install:
	
	install cmd/ipwatcher/ipwatcher /usr/bin/ipwatcher
	chmod 755 /usr/bin/ipwatcher
	install ipwatcher.service /usr/lib/systemd/system/ipwatcher.service

	systemctl daemon-reload
	systemctl enable ipwatcher.service
	systemctl start ipwatcher.service

remove: 

	systemctl stop ipwatcher.service
	systemctl disable ipwatcher.service

	rm -rf /usr/bin/ipwatcher
	rm -rf /usr/lib/systemd/system/ipwatcher.service

	systemctl daemon-reload

config:

	mkdir /etc/ipwatcher && touch /etc/ipwatcher/config.json