package util

import "testing"

func TestExtractIPPort(t *testing.T) {
	var ip, port string

	ip, port = ExtractIPPort("127.0.0.1")
	if ip != "127.0.0.1" {
		t.Error(ip)
	}
	if port != "" {
		t.Error(port)
	}

	ip, port = ExtractIPPort("127.0.0.2:")
	if ip != "127.0.0.2" {
		t.Error(ip)
	}
	if port != "" {
		t.Error(port)
	}

	ip, port = ExtractIPPort("127.0.0.3:4567")
	if ip != "127.0.0.3" {
		t.Error(ip)
	}
	if port != "4567" {
		t.Error(port)
	}

	ip, port = ExtractIPPort(":5678")
	if ip != "" {
		t.Error(ip)
	}
	if port != "5678" {
		t.Error(port)
	}
}
