package main

import (
	"flag"
)

// CLI options
var (
	DateFrom         = flag.String("-date-from", "", "Fetch virtual machines from this date")
	DateTo           = flag.String("-date-to", "", "Fetch virtual machines up to this date")
	Period           = flag.String("-period", "", "Fetch virtual machine for this period (year,month,day)")
	IncludeGroups    = flag.String("-include-groups", "", "Groups whitelist. Comma-separated list of groups")
	ExcludeGroups    = flag.String("-exclude-groups", "", "Groups blacklist. Comma-separated list of groups")
	OneEndpoint      = flag.String("-source", "", "OpenNebula server IP:PORT/RPC2")
	OneKey           = flag.String("-key", "", "OpenNebula authentication key")
	GroupsFile       = flag.String("-groups-file", "", "Groups configuration file path")
	GoatServerIP     = flag.String("-destination", "", "Goat server IP:PORT")
	ClientIdentifier = flag.String("-identifier", "", "Client identifier string")
)

func main() {
	if flag.NArg() == 0 {
		flag.PrintDefaults()
		return
	}

	flag.Parse()
}
