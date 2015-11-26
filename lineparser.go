package main

import (
	"regexp"
	"strings"
)

func skipComment(line string) string {
	re := regexp.MustCompile("#.*")
	line = re.ReplaceAllString(line, "")
	return line
}

func parseLine(line string) HostEntry {
	entry := HostEntry{}
	entry.rawLine = strings.ToLower(strings.TrimSpace(skipComment(line)))
	if !entry.isEmpty() {
		entry.parse()
	}
	return entry
}
