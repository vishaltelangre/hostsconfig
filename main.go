package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	humanizedTemplate = "./templates/humanized"
	standardTemplate  = "./templates/standard"
	DefaultPath       = "/etc/hosts"
	Version           = "0.0.1"
)

var (
	showVersion         *bool   = flag.Bool("version", false, "Show version info")
	showHelp            *bool   = flag.Bool("help", false, "Show help and usage of command")
	hostsFilePath       *string = flag.String("path", DefaultPath, "Path of hosts file")
	previewNeeded       *bool   = flag.Bool("preview", false, "Standard preview of hosts file")
	prettyPreviewNeeded *bool   = flag.Bool("pretty-preview", false, "Pretty preview of hosts file")
)

func printUsage() {
	fmt.Println(`
Hosts file manager

Usage:
  hostsconfig [OPTIONS]

The OPTIONS are:
  -sp, --standard-preview       Standard preview of file
  -pp, --pretty-preview         Pretty preview of file
  -f, --path                    Path to hosts file (Default: /etc/hosts)
  -h, --help                    Show this usage help
  -v, --version                 Display version
`)
}

func parseFlags(config Config) {
	flag.Parse()

	if *showVersion {
		fmt.Printf("hostsconfig - v%s\n", Version)
		return
	}

	if *previewNeeded {
		config.StandardPreview()
		return
	}

	if *prettyPreviewNeeded {
		config.PrettyPreview()
		return
	}

	printUsage()

}

func main() {
	file, err := os.Open(*hostsFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var config Config

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := parseLine(scanner.Text())
		if !entry.isEmpty() {
			config.AddEntry(entry)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	parseFlags(config)
}

func init() {
	flag.BoolVar(showVersion, "v", false, "Show version info")
	flag.BoolVar(showHelp, "h", false, "Show help and usage of command")
	flag.StringVar(hostsFilePath, "f", DefaultPath, "Path of hosts file")
	flag.BoolVar(previewNeeded, "sp", false, "Standard preview of hosts file")
	flag.BoolVar(prettyPreviewNeeded, "pp", false, "Pretty preview of hosts file")
}