package util

import (
	ghfsUtil "mjpclab.dev/ghfs/src/util"
	"strings"
)

func ExtractIPPort(strIPPort string) (ip, port string) {
	ip, port = ghfsUtil.ExtractHostnamePort(strIPPort)

	if len(ip) > 1 && ip[0] == '[' && ip[len(ip)-1] == ']' { // IPv6
		ip = ip[1 : len(ip)-1]
		if percentIndex := strings.IndexByte(ip, '%'); percentIndex >= 0 {
			ip = ip[:percentIndex]
		}
	}

	if len(port) > 0 && port[0] == ':' {
		port = port[1:]
	}

	return
}
