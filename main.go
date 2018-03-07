package main

import (
	"encoding/json"
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
	// Pass first CLI arugment as config file location
	config := LoadConfiguration(CliArguments[0])
	// Print config to stop "declared and not used error"
	fmt.Println(config)
	fmt.Println("Main() - Ending application")
}

// LoadConfiguration file then return Config struct
func LoadConfiguration(filename string) Config {
	fmt.Println("\nLoadConfiguration() - Loading config")
	fmt.Println("LoadConfiguration() - Using config file -", filename)
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		fmt.Println("ERROR: LoadConfiguration() - Configuration file loading has ERRORS")
		fmt.Println(err.Error())
	} else {
		fmt.Println("LoadConfiguration() - Config file has loaded with NO errors")
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	fmt.Println("LoadConfiguration() - Loaded -", config)
	return config
}

// LoadICAL and return current event summary ON/OFF
// Check ICAL current event payload
// SetServerState power state from desired state
// Set payload to match desired state
