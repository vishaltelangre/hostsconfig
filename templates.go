package main

const (
	// Standard template
	standardTemplate = `
# Host Database

{{ if .IPv4 }}# IPv4 capable hosts{{ template "HostEntries" .IPv4 }}{{ end }}

{{ if .IPv6 }}# IPv6 capable hosts{{ template "HostEntries" .IPv6 }}{{ end }}

{{/* HostEntries Standard Template */}}
{{ define "HostEntries" }}{{ range $_, $entry := . }}
{{ printf "%-30s" $entry.IP.String }}{{ with $hostnames := $entry.Hostnames }}{{ range $_, $host := . }} {{ $host.Name }}{{ end }}{{ end }}{{ end }}{{ end }}`

	// Humanized/beautified template
	humanizedTemplate = `{{ if .IPv4 }}
=============
IPv4 Hosts
=============
{{ template "HostEntries" .IPv4 }}{{ end }}
{{ if .IPv6 }}
=============
IPv6 Hosts
=============
{{ template "HostEntries" .IPv6 }}{{ end }}

{{/* HostEntries Humanized Template */}}
{{ define "HostEntries" }}{{ range $_, $entry := . }}
{{ printf "%s" $entry.IP.String }}{{ with $hostnames := $entry.Hostnames }}{{ range $_, $host := . }}
  - {{ $host.Name }}{{ if $host.IsCanonical }} [CANONICAL] {{ end }}{{ end }}
{{ end }}{{ end }}{{ end }}`
)