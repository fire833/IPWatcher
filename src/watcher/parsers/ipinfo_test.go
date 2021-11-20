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

package parsers

import (
	"fmt"
	"testing"
)

type TestingConfig struct {
	Ip string `json:"ip"`
}

func TestIPInfo(t *testing.T) {
	c := &TestingConfig{}
	fmt.Print(c.Ip)
	// file, err := os.ReadFile("keyfile.json")
	// if err != nil {
	// 	t.Fail()
	// 	t.Log(err)
	// }

	// c := &TestingConfig{}
	// _ = json.Unmarshal(file, c)

	// info := IPInfoParser{}

	// err1 := info.Get(net.ParseIP(c.ip))
	// if err1 != nil {
	// 	t.Fail()
	// 	t.Log(err)
	// }

}
