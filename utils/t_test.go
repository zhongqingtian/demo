package utils

import "testing"

func TestIpAddrList(t *testing.T) {
	ip,err := IpAddrList()
	t.Log(ip,err)
}