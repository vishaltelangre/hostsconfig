package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
	"text/template"
)

const (
	creditsLine = `

# ===========================================================
# This file is generated using 'hostsconfig' utility.
# We recommend to use the same tool to modify this file.
#
# Download it from here:
# https://github.com/vishaltelangre/hostsconfig
# ===========================================================
`
)

type Config struct {
	HostEntries []HostEntry
}

func (c *Config) Parse() {
	file, err := os.Open(hostsFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c.AddEntryFromLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (c *Config) SaveChanges() {
	file, err := os.OpenFile(hostsFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprint(c.StandardPreview(), creditsLine))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\n Hosts file successfully updated!\n\n")
}

func (c *Config) AddEntry(e HostEntry) {
	for _, entry := range c.HostEntries {
		// Add host names into host entry if ip address already exists
		if entry.IP.String() == e.IP.String() {
			entry.addHostnames(e.Hostnames)
			return
		}
	}

	c.HostEntries = append(c.HostEntries, e)
	return
}

func (c *Config) AddEntryFromLine(line string) {
	if entry := parseLine(line); !entry.isEmpty() {
		c.AddEntry(entry)
		return
	}
}

func (c *Config) DeleteEntryFromLine(line string) {
	ip := net.ParseIP(line)
	if ip != nil {
		c.DeleteEntry(parseLine(line))
	}

	line = strings.ToLower(strings.TrimSpace(stripComment(line)))
	c.deleteHostnames(line)
}

func (c *Config) DeleteEntryAtIndex(i int) {
	ip := c.HostEntries[i].IP
	ensureIPIsNotLoopback(ip)

	copy(c.HostEntries[i:], c.HostEntries[i+1:])
	c.HostEntries[len(c.HostEntries)-1] = HostEntry{}
	c.HostEntries = c.HostEntries[:len(c.HostEntries)-1]
}

func (c *Config) DeleteEntryWithIP(ip net.IP) {
	ensureIPIsNotLoopback(ip)

	for i, entry := range c.HostEntries {
		if entry.IP.String() == ip.String() {
			c.DeleteEntryAtIndex(i)
			return
		}
	}
}

func (c *Config) DeleteEntry(e HostEntry) {
	for i, entry := range c.HostEntries {
		if entry.IP.String() == e.IP.String() {
			entry.deleteHostnames(e.Hostnames)

			// Delete entry if it has no hosts
			if len(entry.Hostnames) <= 0 {
				c.DeleteEntryAtIndex(i)
			}

			// Delete an entry with asked IP address
			if len(e.Hostnames) <= 0 {
				c.DeleteEntryAtIndex(i)
			}

			return
		}
	}
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

func (c *Config) StandardPreview() string {
	return preview(c, standardTemplate, "standard")
}

func (c *Config) BeautifiedPreview() string {
	return preview(c, humanizedTemplate, "humanized")
}

func preview(config *Config, tmpl string, tmplBase string) string {
	buf := new(bytes.Buffer)
	t := template.Must(template.New(tmplBase).Parse(tmpl))

	entries := make(map[string][]HostEntry)
	entries["IPv4"] = config.IPv4Entries()
	entries["IPv6"] = config.IPv6Entries()

	err := t.Execute(buf, entries)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s", strings.TrimSpace(buf.String()))
}

func (c *Config) deleteHostnames(line string) {
	hostnames := make([]Hostname, 0)
	for _, name := range strings.Fields(line) {
		hostnames = append(hostnames, Hostname{Name: name})
	}

	hostEntriesCopy := make([]HostEntry, len(c.HostEntries))
	copy(hostEntriesCopy, c.HostEntries)

	for _, entry := range hostEntriesCopy {
		entry.deleteHostnames(hostnames)

		// Delete an entry if it has no hosts
		if len(entry.Hostnames) <= 0 {
			c.DeleteEntryWithIP(entry.IP)
		}
	}
}

func ensureIPIsNotLoopback(ip net.IP) {
	if ip.IsLoopback() {
		log.Fatalf("'%s' is a loopback IP, can't delete it\n", ip.String())
	}
}

func stripComment(line string) string {
	re := regexp.MustCompile("#.*")
	line = re.ReplaceAllString(line, "")
	return line
}

func parseLine(line string) HostEntry {
	entry := HostEntry{}
	entry.rawLine = strings.ToLower(strings.TrimSpace(stripComment(line)))
	if !entry.isEmpty() {
		entry.parse()
	}
	return entry
}
