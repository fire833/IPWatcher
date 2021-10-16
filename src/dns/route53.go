package dns

type Route53Updater struct {
}

type Route53Config struct {
}

func init() {

}

func (d *Route53Updater) UpdateDNS(OldHost string, NewEntry *DNSEntry) error {
	return nil
}

func (d *Route53Updater) Name() string { return "Route 53" }
