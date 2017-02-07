package testhelpers

import (
	"fmt"
	"net"
	"strconv"
)

func GetOpenPort() uint16 {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(fmt.Sprintf("Failed to get open port: %s", err))
	}
	l.Close()

	return extractPortFromAddr(l.Addr().String())
}

func extractPortFromAddr(addr string) uint16 {
	_, pstr, err := net.SplitHostPort(addr)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse addr to get port: %s", err))
	}

	p, err := strconv.ParseUint(pstr, 10, 16)
	if err != nil {
		panic(fmt.Sprintf("Failed to convert port string: %s", err))
	}

	return uint16(p)
}
