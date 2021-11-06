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
