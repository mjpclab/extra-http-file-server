package lib

import (
	"net/netip"
	"sort"
	"strings"
)

type IPRangeMan struct {
	sorted     bool
	ipv4Ranges ipRangeList
	ipv6Ranges ipRangeList
}

func (man *IPRangeMan) AddByString(strRange string) error {
	prefix, err := createRange(strRange)
	if err != nil {
		return err
	}

	addr := prefix.Addr()
	if addr.Is4() {
		man.ipv4Ranges = append(man.ipv4Ranges, prefix)
		man.sorted = false
	} else if addr.Is6() {
		man.ipv6Ranges = append(man.ipv6Ranges, prefix)
		man.sorted = false
	}
	return nil
}

func (man *IPRangeMan) sort() {
	if man.sorted {
		return
	}

	man.sorted = true
	sort.Sort(man.ipv4Ranges)
	sort.Sort(man.ipv6Ranges)
}

func (man *IPRangeMan) MatchAddr(addr netip.Addr) bool {
	if !man.sorted {
		man.sort()
	}

	if addr.Is4() {
		return man.ipv4Ranges.BinaryMatchAddr(addr)
	} else if addr.Is6() {
		return man.ipv6Ranges.BinaryMatchAddr(addr)
	} else {
		return false
	}
}

func (man *IPRangeMan) MatchStringAddr(strAddr string) bool {
	addr, err := netip.ParseAddr(strAddr)
	if err != nil {
		return false
	}
	return man.MatchAddr(addr)
}

func (man *IPRangeMan) HasData() bool {
	return len(man.ipv4Ranges) > 0 || len(man.ipv6Ranges) > 0
}

func NewIPRangeMan() *IPRangeMan {
	return &IPRangeMan{}
}

func createRange(strRange string) (prefix netip.Prefix, err error) {
	if slashIndex := strings.IndexByte(strRange, '/'); slashIndex >= 0 {
		prefix, err = netip.ParsePrefix(strRange)
	} else {
		var addr netip.Addr
		addr, err = netip.ParseAddr(strRange)
		if err == nil {
			prefix = netip.PrefixFrom(addr, addr.BitLen())
		}
	}
	return
}
