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

	ip, port = ExtractIPPort("[fe80::1]")
	if ip != "fe80::1" {
		t.Error(ip)
	}
	if port != "" {
		t.Error(port)
	}

	ip, port = ExtractIPPort("[fe80::2%eth0]")
	if ip != "fe80::2" {
		t.Error(ip)
	}
	if port != "" {
		t.Error(port)
	}

	ip, port = ExtractIPPort("[fe80::3]:1234")
	if ip != "fe80::3" {
		t.Error(ip)
	}
	if port != "1234" {
		t.Error(port)
	}

	ip, port = ExtractIPPort("[fe80::4%eth0]:1234")
	if ip != "fe80::4" {
		t.Error(ip)
	}
	if port != "1234" {
		t.Error(port)
	}
}
