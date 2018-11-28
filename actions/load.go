package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/metrosystems-cpe/DDOM/config"
	"github.com/metrosystems-cpe/DDOM/ddObjects"
	datadog "github.com/zorkian/go-datadog-api"
)

func loadDashboard(path string, orgCfg config.Organisation) {
	files, err := ioutil.ReadDir(path) // plm ??
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		jsonFile, err := os.Open(path + "/" + f.Name())
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var result datadog.Dashboard
		json.Unmarshal([]byte(byteValue), &result)
		err = ddObjects.CreateDashboards(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, &result)
		if err != nil {
			fmt.Println(fmt.Sprintf("Could not create datadog dashboard from file %v\n", f.Name()))
		} else {
			fmt.Println(fmt.Sprintf("Dashboard created successfully from file %v\n", f.Name()))
		}
	}
}

func loadMonitor(path string, orgCfg config.Organisation) {
	files, err := ioutil.ReadDir(path) // plm ??
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		// fmt.Println(f.Name())
		jsonFile, err := os.Open(path + "/" + f.Name())
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var result datadog.Monitor
		json.Unmarshal([]byte(byteValue), &result)
		if _, ok := result.GetQueryOk(); ok {
			err = ddObjects.CreateMonitors(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, &result)
			if err == nil {
				fmt.Println(fmt.Sprintf("Monitor created successfully from file: %v\n", f.Name()))
				// fmt.Printf("\nMonitor created successfully from file")
			} else {
				fmt.Printf("Could not create datadog monitor from file %v\n", f.Name())
			}
		} else {
			fmt.Printf("Not a monitor json - %v\n", f.Name())
		}

	}
}

func loadTimeboard(path string, orgCfg config.Organisation) {
	files, err := ioutil.ReadDir(path) // plm ??
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		jsonFile, err := os.Open(path + "/" + f.Name())
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var result datadog.Screenboard
		json.Unmarshal([]byte(byteValue), &result)
		err = ddObjects.CreateScreenboards(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, &result)
		if err != nil {
			fmt.Println(fmt.Sprintf("Could not create datadog screenboard from file %v\n", f.Name()))
		} else {
			fmt.Println(fmt.Sprintf("Screeboard created successfully from file %v\n", f.Name()))
		}
	}
}
