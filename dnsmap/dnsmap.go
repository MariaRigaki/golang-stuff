package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var hostname = flag.String("t", "google.com", "target domain")

func main() {
	flag.Parse()

	if isDomainName(*hostname) {
		addrs, err := net.LookupHost(*hostname)
		if err != nil {
			fmt.Println(err)
		}
		for _, a := range addrs {
			if isPrivate(a) {
				fmt.Print("[+] ")
			}
			fmt.Println(a)
		}

	} else {
		fmt.Println("Invalid domain name: " + *hostname)
	}

}

func isPrivate(ip string) bool {
	ip4 := net.ParseIP(ip)
	if ip4 != nil {
		splits := strings.SplitN(ip, ".", 4)
		if (splits[0] == "10") || (splits[0] == "192" && splits[1] == "168") {
			return true
		}
		b, err := strconv.Atoi(splits[1])
		if err != nil {
			return false
		}
		if splits[0] == "172" && (b >= 16 || b <= 31) {
			return true
		}
	}
	return false
}

func isDomainName(s string) bool {
	// See RFC 1035, RFC 3696.
	if len(s) == 0 {
		return false
	}
	if len(s) > 255 {
		return false
	}

	last := byte('.')
	ok := false // Ok once we've seen a letter.
	partlen := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		default:
			return false
		case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_':
			ok = true
			partlen++
		case '0' <= c && c <= '9':
			// fine
			partlen++
		case c == '-':
			// Byte before dash cannot be dot.
			if last == '.' {
				return false
			}
			partlen++
		case c == '.':
			// Byte before dot cannot be dot, dash.
			if last == '.' || last == '-' {
				return false
			}
			if partlen > 63 || partlen == 0 {
				return false
			}
			partlen = 0
		}
		last = c
	}
	if last == '-' || partlen > 63 {
		return false
	}

	return ok
}
