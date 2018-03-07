package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n------------------------------\nMain() - Starting application")
	fmt.Println("Main() - Ending application")
}

// LoadConfiguration file then return Config struct
// LoadICAL and return current event summary ON/OFF
// Check ICAL current event payload
// SetServerState power state from desired state
// Set payload to match desired state
