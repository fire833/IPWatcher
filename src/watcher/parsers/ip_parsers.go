package parsers

import (
	"net"
)

var IpResolver IpParserLocator
var IpInfoGatherer IpInformationGatherer

type IpParserLocator interface {
	// Use this IP resolver to get the public IP address your host is located on.
	// Will set the raw message body to be parsed by the other sub-methods defined above.
	// This method needs to be called first because it configres the object to meet the
	// other guarantees of the interface later on.
	Get() error
	// Parse an IP address and return the raw net.IP value for parsing down the line.
	ParseIP() net.IP
	// Extract the IP address from whatever response was returned from the site.
	GetStringIP() string
	// Returns true if an object returned contains an IPv4 address,
	// if not then it can be assumed it is an IPv6 address.
	IsV4() bool
	// Returns the body of the response that was returned for the IP resolver.
	Body() []byte
	// Return the name of the parser
	Name() string
}

type IpInformationGatherer interface {
	// Gets information about a specific IP address.
	// Makes the physical call to the remote API, and
	// internally configures itself to meet the other guarantees of the interface.
	Get(net.IP)
	// Return the ASN of the provided IP.
	GetASN() string
	// Return the hostname of the provided IP.
	GetHostname() string
	// Returns the string array of City, Region/State, and Country, in that order.
	// The array size should always return as size 3.
	GetLocality() []string
	// Return rganization of the provided IP.
	GetOrg() string
	// Returns the latitude and logitude of the Ip address geolocation approximation.
	// The first index of the returned array should be latitude, the second index longitude.
	// The returned array should always be of length 2.
	GetLocation() []float64
	// Returns the name of the IpInfoGatherer
	Name() string
}
