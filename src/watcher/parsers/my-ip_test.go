package parsers

import "testing"

func TestMyIP(t *testing.T) {

	info := MyIPParser{}
	if err := info.Get(); err != nil {
		t.Fail()
		t.Log(err)
	}

}
