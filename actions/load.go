package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/metrosystems-cpe/DDOM/config"
	"github.com/metrosystems-cpe/DDOM/ddObjects"
	"github.com/metrosystems-cpe/DDOM/helpers"
	datadog "github.com/zorkian/go-datadog-api"
)

func getFiles(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path) // plm ??
	helpers.PanicIfError(err, "Could not read directory")
	return files
}

func unmarshalData(path string, name string, result interface{}) {
	jsonFile, err := os.Open(path + "/" + name)
	helpers.PanicIfError(err, "Could not open file")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), result)
}

func logOutput(err error, objType string, name string) {
	if err != nil {
		fmt.Println(fmt.Sprintf("could not create %s from file %v\n", name, objType))
	} else {
		fmt.Println(fmt.Sprintf("%s created successfully from file %v\n", name, objType))
	}
}

func loadDashboard(path string, orgCfg config.Organisation) {
	files := getFiles(path)
	for _, f := range files {
		var result datadog.Dashboard
		unmarshalData(path, f.Name(), &result)
		err := ddObjects.CreateDashboards(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, &result)
		logOutput(err, "dashboard", f.Name())
	}
}

func loadMonitor(path string, orgCfg config.Organisation) {
	files := getFiles(path)
	for _, f := range files {
		var result datadog.Monitor
		unmarshalData(path, f.Name(), &result)
		if _, ok := result.GetQueryOk(); ok {
			err := ddObjects.CreateMonitors(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, &result)
			logOutput(err, "monitor", f.Name())
		} else {
			fmt.Printf("Not a valid monitor json - %v\n", f.Name())
		}

	}
}

func loadTimeboard(path string, orgCfg config.Organisation) {
	files := getFiles(path)
	for _, f := range files {
		var result datadog.Screenboard
		unmarshalData(path, f.Name(), &result)
		err := ddObjects.CreateScreenboards(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, &result)
		logOutput(err, "screenboard", f.Name())
	}
}
