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
