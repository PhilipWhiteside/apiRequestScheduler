package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/PuloV/ics-golang"
)

// Config for user configuraiton used throughout
type Config struct {
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
	// LoadICAL and return current event summary ON/OFF
	desiredStatus := LoadICAL(config)
	// Print desiredStatus to stop "declared and not used error"
	fmt.Println(desiredStatus)
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
func LoadICAL(config Config) int {
	fmt.Println("\nLoadICAL() - Loading ICAL")
	//  create new parser
	parser := ics.New()
	// set the filepath for the ics files
	ics.FilePath = "tmp/new/"
	// we dont want to delete the temp files
	ics.DeleteTempFiles = false
	ics.RepeatRuleApply = true
	// get the input chan
	inputChan := parser.GetInputChan()
	// send the calendar urls to be parsed
	inputChan <- config.ICALPath
	fmt.Println("LoadICAL() - Created ICS filesystem settings & load ICAL")
	//  wait for the calendar to be parsed
	parser.Wait()
	// get all calendars in this parser
	cal, _ := parser.GetCalendars()
	// Check ICAL payload for errors
	fmt.Println("LoadICAL() - Events loaded -", cal)
	if len(cal) > 0 {
		fmt.Println("LoadICAL() - Calender entries have been pulled. Slice is populated")
	} else {
		fmt.Println("LoadICAL() - Calender entries have not been pulled. Slice is empty")
		fmt.Println("LoadICAL() - App will panic")
		fmt.Println("LoadICAL() - Please check the URL of the calendar ICS")
	}
	// Check ICAL current event payload

	fmt.Println("LoadICAL() - Loaded")
	return 0
}

// SetServerState power state from desired state
// Set payload to match desired state
