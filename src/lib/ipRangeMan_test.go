package lib

import "testing"

func TestIPRangeMan_MatchStringAddr(t *testing.T) {
	man := NewIPRangeMan()
	man.AddByString("192.168.2.0/24")
	man.AddByString("192.168.3.0/24")
	man.AddByString("192.168.1.0/24")
	man.AddByString("172.16.1.1")
	man.AddByString("fe80::1/64")

	if !man.MatchStringAddr("192.168.1.5") {
		t.Error()
	}

	if !man.MatchStringAddr("192.168.2.6") {
		t.Error()
	}

	if !man.MatchStringAddr("192.168.3.7") {
		t.Error()
	}

	if man.MatchStringAddr("192.168.4.8") {
		t.Error()
	}

	if !man.MatchStringAddr("172.16.1.1") {
		t.Error()
	}

	if man.MatchStringAddr("172.16.1.2") {
		t.Error()
	}

	if !man.MatchStringAddr("fe80::ff") {
		t.Error()
	}

	if man.MatchStringAddr("ff80::ff") {
		t.Error()
	}
}
