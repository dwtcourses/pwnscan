// git: github.com/nolimitcarter/pwnscan
// A modular framework designed to automate security tests.
// Licensed under the GPL

package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"time"
)

func whois(domain string) (string, error) {
	return whoisTimeout(domain, time.Second*5)
}

func whoisTimeout(domain string, timeout time.Duration) (result string, err error) {
	var (
		parts      []string
		zone       string
		buffer     []byte
		connection net.Conn
	)
	parts = strings.Split(domain, ".")
	if len(parts) < 2 {
		err = fmt.Errorf("Domain(%s) name is wrong", domain)
		return
	}

	zone = parts[len(parts)-1]
	server, ok := servers[zone]

	if !ok {
		err = fmt.Errorf("No such server for zone %s. Domain %s", zone, domain)
		return
	}
	connection, err = net.DialTimeout("tcp", net.JoinHostPort(server, "43"), timeout)

	if err != nil {
		// returns the net.Conn error
		return
	}
	defer connection.Close()
	connection.Write([]byte(domain + "\r\n"))
	buffer, err = ioutil.ReadAll(connection)

	if err != nil {
		return
	}
	result = string(buffer[:])
	return
}
