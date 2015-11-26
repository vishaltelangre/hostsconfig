package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"
)

type Config struct {
	HostEntries []HostEntry
}

func (c *Config) AddEntry(e HostEntry) []HostEntry {
	for _, entry := range c.HostEntries {
		// Add host names into host entry if ip address already exists
		if entry.IP.String() == e.IP.String() {
			entry.addHostnames(e.Hostnames)
			return c.HostEntries
		}
	}

	c.HostEntries = append(c.HostEntries, e)
	return c.HostEntries
}

func (c *Config) IPv4Entries() []HostEntry {
	ipv4Entries := make([]HostEntry, 0)
	for _, entry := range c.HostEntries {
		if entry.IP.To4() != nil {
			ipv4Entries = append(ipv4Entries, entry)
		}
	}

	return ipv4Entries
}

func (c *Config) IPv6Entries() []HostEntry {
	ipv6Entries := make([]HostEntry, 0)
	for _, entry := range c.HostEntries {
		if entry.IP.To4() == nil {
			ipv6Entries = append(ipv6Entries, entry)
		}
	}

	return ipv6Entries
}

func (c *Config) StandardPreview() {
	preview(c, standardTemplate, "standard")
}

func (c *Config) PrettyPreview() {
	preview(c, humanizedTemplate, "humanized")
}

func preview(config *Config, tmpl string, tmplBase string) {
	buf := new(bytes.Buffer)
	t := template.New(tmplBase)
	t, err := t.ParseFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	entries := make(map[string][]HostEntry)
	entries["IPv4"] = config.IPv4Entries()
	entries["IPv6"] = config.IPv6Entries()

	err = t.Execute(buf, entries)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n\n", strings.TrimSpace(buf.String()))
}
