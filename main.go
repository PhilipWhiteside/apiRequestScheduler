package main

import (
	"fmt"
	"os"
)

// Config for user configuraiton used throughout
type Config struct {
	ConfigVersion            int
	AppDebug                 bool
	ServerAPIPath            string
	ServerAPIUser            string
	ServerAPIPass            string
	ServerAPIPayloadDefault  string
	ServerAPIIgnoreCertError bool
	ICALPath                 string
}

func main() {
	fmt.Println("\n------------------------------\nMain() - Starting application")
	// READ command-line arguments to get config file path
	CliArguments := os.Args[1:]
	fmt.Println("Main() - CLI arguments loaded -", CliArguments)
	fmt.Println("Main() - Ending application")
}

// LoadConfiguration file then return Config struct
// LoadICAL and return current event summary ON/OFF
// Check ICAL current event payload
// SetServerState power state from desired state
// Set payload to match desired state
