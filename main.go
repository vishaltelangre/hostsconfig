package main

import (
	"flag"
	"fmt"
)

const (
	DefaultPath = "/etc/hosts"

	showVersionDesc   = "Show version info"
	showHelpDesc      = "Show help and usage of command"
	hostsFilePathDesc = "`Path` of hosts file"
	showListDesc      = "List hosts file on standard output"
	beautifyDesc      = "Display beautified hosts file on standard output"
	addHostDesc       = "Add a host `entry`"
	deleteHostDesc    = "Delete a host `entry`"
	saveDesc          = "Save changes"
)

var (
	showVersion   bool
	showHelp      bool
	hostsFilePath string
	showList      bool
	beautify      bool
	addHost       string
	deleteHost    string
	save          bool
)

func init() {
	flag.BoolVar(&showVersion, "version", false, showVersionDesc)
	flag.BoolVar(&showVersion, "v", false, showVersionDesc)
	flag.BoolVar(&showHelp, "help", false, showHelpDesc)
	flag.BoolVar(&showHelp, "h", false, showHelpDesc)
	flag.StringVar(&hostsFilePath, "path", DefaultPath, hostsFilePathDesc)
	flag.StringVar(&hostsFilePath, "f", DefaultPath, hostsFilePathDesc)
	flag.BoolVar(&showList, "list", false, showListDesc)
	flag.BoolVar(&showList, "l", false, showListDesc)
	flag.BoolVar(&beautify, "beautify", false, beautifyDesc)
	flag.BoolVar(&beautify, "b", false, beautifyDesc)
	flag.StringVar(&addHost, "add", "", addHostDesc)
	flag.StringVar(&addHost, "a", "", addHostDesc)
	flag.StringVar(&deleteHost, "delete", "", deleteHostDesc)
	flag.StringVar(&deleteHost, "d", "", deleteHostDesc)
	flag.BoolVar(&save, "s", false, saveDesc)
	flag.BoolVar(&save, "save", false, saveDesc)
}

func printUsage() {
	fmt.Println(`
Hosts file utility

Usage:
  hostsconfig [OPTIONS]

The OPTIONS are:
  -l, --list           List hosts file on standard output
  -b, --beautify       Display beautified hosts file on standard output
  -a=, --add=          Add a host entry
  -d=, --delete=       Delete a host entry
  -s, --save           Save changes
  -f=, --path=         Path to hosts file (Default: /etc/hosts)
  -h, --help           Show this usage help
  -v, --version        Display version
`)
}

func main() {

	flag.Parse()

	if showVersion {
		fmt.Printf("hostsconfig - v%s\n", Version)
		return
	}

	if showHelp {
		printUsage()
		return
	}

	var config Config

	config.Parse()

	if len(addHost) > 0 {
		config.AddEntryFromLine(addHost)
	}

	if len(deleteHost) > 0 {
		config.DeleteEntryFromLine(deleteHost)
	}

	if save {
		config.SaveChanges()
	}

	if showList {
		fmt.Printf("\n%s\n\n", config.StandardPreview())
		return
	}

	if beautify {
		fmt.Printf("\n%s\n\n", config.BeautifiedPreview())
		return
	}

	printUsage()
}
