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
	"log"
	"os/exec"
	"runtime"

	"github.com/fire833/ipwatcher/src/config"
)

var OSIsUsed bool = false

type OSNotification struct {
	l *Limit
	e string
}

func init() {
	n := new(OSNotification)
	config.RegisterConfig(n.Name(), nil, OSIsUsed, false)
}

func (n *OSNotification) Name() string { return "OS-Notification" }

func (n *OSNotification) Send(msg *Message) error {
	switch runtime.GOOS {
	case "linux":
		{
			path, err := exec.LookPath("notify-send")
			if err != nil {
				log.Default().Printf("Error with sending %s, error finding 'notify-send': %v", n.Name(), err)
				return err
			}

			err1 := exec.Command(path, msg.DeviceName, msg.Message).Run()
			if err1 != nil {
				return err1
			}
		}
	case "windows":
		{
			// TODO get notifications working in windows
		}
	}

	return nil
}

func (n *OSNotification) GetLimit() *Limit {
	return n.l
}

func (n *OSNotification) Error() string {
	return n.e
}
