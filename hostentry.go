package main

import (
	"log"
	"net"
	"strings"
)

type HostEntry struct {
	IP        net.IP
	Hostnames []Hostname
	rawLine   string
}

func (e *HostEntry) addHostnames(names []Hostname) []Hostname {
	for _, hostname := range names {
		hostname.IsCanonical = false
		e.AddHostname(hostname)
	}

	return e.Hostnames
}

func (e *HostEntry) AddHostname(h Hostname) []Hostname {
	for _, hostname := range e.Hostnames {
		// Ignore if its a duplicate host name
		if hostname.Name == h.Name {
			return e.Hostnames
		}
	}

	e.Hostnames = append(e.Hostnames, h)
	return e.Hostnames
}

func (e *HostEntry) parse() {
	tokens := strings.Fields(e.rawLine)
	ip := net.ParseIP(tokens[0])
	hostnames := tokens[1:]

	if ip == nil {
		log.Fatal("%s is invalid IP address", tokens[0])
	}

	e.IP = ip

	for index, name := range hostnames {
		hostname := Hostname{Name: name, IsCanonical: index == 0}
		e.AddHostname(hostname)
	}
}

func (e *HostEntry) isEmpty() bool {
	return len(e.rawLine) <= 0
}
