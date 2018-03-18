package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuloV/ics-golang"
)

// Config for user configuration used throughout
type Config struct {
	ServerAPIPath            string
	ServerAPIUser            string
	ServerAPIPass            string
	ServerAPIPayloadDefault  string
	CalendarValue001         string
	ServerAPIPayload001      string
	CalendarValue002         string
	ServerAPIPayload002      string
	CalendarValue003         string
	ServerAPIPayload003      string
	CalendarValue004         string
	ServerAPIPayload004      string
	CalendarValue005         string
	ServerAPIPayload005      string
	ServerAPIIgnoreCertError bool
	ICALPath                 string
}

func main() {
	fmt.Println("\n------------------------------\nMain() - Starting application")
	// READ command-line arguments to get config file path
	CliArguments := os.Args[1:]
	fmt.Println("Main() - CLI arguments loaded -", CliArguments)
	// Pass first CLI argument as config file location
	config := LoadConfiguration(CliArguments[0])
	// LoadICAL and return current event summary ON/OFF
	desiredStatus := LoadICAL(config)
	// Send request to API with desiredStatus
	SetServerState(config, desiredStatus)
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
	for _, e := range cal[0].GetEvents() {
		now := time.Now()
		if now.After(e.GetStart()) && now.Before(e.GetEnd()) {
			if strings.EqualFold(config.CalendarValue001, e.GetSummary()) {
				fmt.Println("LoadICAL() - Current even is", e.GetSummary(), "-", &e)
				return 1
			} else if strings.EqualFold(config.CalendarValue002, e.GetSummary()) {
				fmt.Println("LoadICAL() - Current even is", e.GetSummary(), "-", &e)
				return 2
			} else if strings.EqualFold(config.CalendarValue003, e.GetSummary()) {
				fmt.Println("LoadICAL() - Current even is", e.GetSummary(), "-", &e)
				return 3
			} else if strings.EqualFold(config.CalendarValue004, e.GetSummary()) {
				fmt.Println("LoadICAL() - Current even is", e.GetSummary(), "-", &e)
				return 4
			} else if strings.EqualFold(config.CalendarValue005, e.GetSummary()) {
				fmt.Println("LoadICAL() - Current even is", e.GetSummary(), "-", &e)
				return 5
			}
		}
	}

	fmt.Println("LoadICAL() - Loaded")
	return 0
}

// SetServerState power state from desired state
func SetServerState(config Config, state int) {
	fmt.Println("\nSetServerState() - Setting state")

	if config.ServerAPIIgnoreCertError == true {
		fmt.Println("SetServerState() - Ignoring self signed certificate errors")
		// CHANGE SECURITY TO IGNORE SELF SIGNED CERTIFICATE WARNINGS
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		_, err := http.Get("https://golang.org/")
		if err != nil {
			fmt.Println(err)
		}
	}

	// PERFORM REQUEST ON API
	// Set payload to match desired state
	var payload = strings.NewReader("")
	if state == 0 {
		fmt.Println("SetServerState() - Setting state to DEFAULT")
		payload = strings.NewReader(config.ServerAPIPayloadDefault)
	} else if state == 1 {
		fmt.Println("SetServerState() - Setting state to", config.CalendarValue001)
		payload = strings.NewReader(config.ServerAPIPayload001)
	} else if state == 2 {
		fmt.Println("SetServerState() - Setting state to", config.CalendarValue002)
		payload = strings.NewReader(config.ServerAPIPayload002)
	} else if state == 3 {
		fmt.Println("SetServerState() - Setting state to", config.CalendarValue003)
		payload = strings.NewReader(config.ServerAPIPayload003)
	} else if state == 4 {
		fmt.Println("SetServerState() - Setting state to", config.CalendarValue004)
		payload = strings.NewReader(config.ServerAPIPayload004)
	} else if state == 5 {
		fmt.Println("SetServerState() - Setting state to", config.CalendarValue005)
		payload = strings.NewReader(config.ServerAPIPayload005)
	}
	fmt.Println("SetServerState() - Sending request payload - ", payload)
	// Build & trigger the request
	req, _ := http.NewRequest("POST", config.ServerAPIPath, payload)
	req.SetBasicAuth(config.ServerAPIUser, config.ServerAPIPass)
	req.Header.Add("Cache-Control", "no-cache")
	res, _ := http.DefaultClient.Do(req)
	fmt.Println("SetServerState() - Request made")
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// Print the result of request
	fmt.Println(res)
	fmt.Println(string(body))
}
