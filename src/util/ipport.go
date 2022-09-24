package util

import (
	ghfsUtil "mjpclab.dev/ghfs/src/util"
)

func ExtractIPPort(strIPPort string) (ip, port string) {
	ip, port = ghfsUtil.ExtractHostnamePort(strIPPort)

	if len(ip) > 1 && ip[0] == '[' && ip[len(ip)-1] == ']' { // IPv6
		ip = ip[1 : len(ip)-1]
	}

	if len(port) > 0 && port[0] == ':' {
		port = port[1:]
	}

	return
}
