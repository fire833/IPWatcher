package backend

import (
	"encoding/json"
	"fmt"
	"log/syslog"

	"github.com/fire833/ipwatcher/src/config"
)

var SLC *SyslogConfig
var SyslogIsUsed bool = false

type SyslogConfig struct {
	Network    string `json:"network"`
	RemoteAddr string `json:"r_addr"`
}

type SyslogNotification struct {
	l *Limit
	e string
}

func (c *SyslogConfig) UnmarshalConfig(input []byte) {
	if SyslogIsUsed {
		SLC = new(SyslogConfig)
	} else {
		return
	}

	if err := json.Unmarshal(input, SLC); err == nil {
		n := new(SyslogNotification)
		RegisterNotifier(n)
	} else {
		return
	}
}

func init() {
	n := new(SyslogNotification)
	config.RegisterConfig(n.Name(), SLC, SyslogIsUsed, false)
}

func (n *SyslogNotification) Send(msg *Message) error {
	writer, err := syslog.Dial(SLC.Network, SLC.RemoteAddr, syslog.LOG_NOTICE, "IPWatcher")
	defer writer.Close()
	if err != nil {
		return err
	}

	if err := writer.Notice(fmt.Sprintf("%s: %s", msg.Title, msg.Message)); err != nil {
		return err
	}

	return nil
}

func (n *SyslogNotification) GetLimit() *Limit {
	return n.l
}

func (n *SyslogNotification) Name() string { return "syslog" }

func (n *SyslogNotification) Error() string {
	return n.e
}
