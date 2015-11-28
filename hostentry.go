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

func (e *HostEntry) addHostnames(names []Hostname) {
	for _, hostname := range names {
		hostname.IsCanonical = false
		e.AddHostname(hostname)
	}
}

func (e *HostEntry) AddHostname(h Hostname) {
	for _, hostname := range e.Hostnames {
		// Ignore if its a duplicate host name
		if hostname.Name == h.Name {
			return
		}
	}

	e.Hostnames = append(e.Hostnames, h)
}

func (e *HostEntry) deleteHostnames(names []Hostname) {
	for _, hostname := range names {
		e.DeleteHostname(hostname)
	}
}

func (e *HostEntry) DeleteHostname(h Hostname) bool {
	if h.Name == "localhost" {
		log.Fatal("'localhost' is a loopback host can't delete it")
	}

	for i := range e.Hostnames {
		if e.Hostnames[i].Name == h.Name {
			copy(e.Hostnames[i:], e.Hostnames[i+1:])
			e.Hostnames[len(e.Hostnames)-1] = Hostname{}
			e.Hostnames = e.Hostnames[:len(e.Hostnames)-1]

			// Make sure that first entry should be canonical
			if i == 0 && len(e.Hostnames) > 0 {
				hostname := &e.Hostnames[0]
				hostname.IsCanonical = true
			}

			return true
		}
	}
	return false
}

func (e *HostEntry) parse() {
	fields := strings.Fields(e.rawLine)
	ip := net.ParseIP(fields[0])
	hostnames := fields[1:]

	if ip == nil {
		log.Fatalf("'%s' is an invalid IP address", fields[0])
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
