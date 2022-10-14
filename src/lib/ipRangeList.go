package lib

import (
	"net/netip"
)

type ipRangeList []netip.Prefix

func (list ipRangeList) Len() int {
	return len(list)
}

func (list ipRangeList) Less(i, j int) bool {
	cmpResult := list[i].Addr().Compare(list[j].Addr())
	if cmpResult != 0 {
		return cmpResult < 0
	} else {
		return list[i].Bits() < list[j].Bits()
	}
}

func (list ipRangeList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list ipRangeList) BinaryMatchAddr(addr netip.Addr) bool {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) >> 1
		if list[mid].Contains(addr) {
			return true
		}

		if cmpResult := addr.Compare(list[mid].Addr()); cmpResult < 0 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}
