package parsers

import "testing"

func TestWhatsMyIP(t *testing.T) {

	info := WhatsMyIPAddrParser{}
	if err := info.Get(); err != nil {
		t.Fail()
		t.Log(err)
	}

	info.Body()

}
