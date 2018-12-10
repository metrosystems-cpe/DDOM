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
		fmt.Println(fmt.Sprintf("%v -- could not create %s from file", name, objType))
	} else {
		fmt.Println(fmt.Sprintf("%v -- %s created successfully from file", name, objType))
	}
}

func LoadDashboard(path string, orgCfg config.Organisation) []int {
	files := getFiles(path)
	var ids []int
	for _, f := range files {
		var result datadog.Dashboard
		unmarshalData(path, f.Name(), &result)
		id, err := ddObjects.CreateDashboards(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, &result)
		logOutput(err, "dashboard", f.Name())
		if id != 0 {
			ids = append(ids, id)
		}
	}
	return ids
}

func LoadMonitor(path string, orgCfg config.Organisation) []int {
	files := getFiles(path)
	var ids []int
	for _, f := range files {
		var result datadog.Monitor
		unmarshalData(path, f.Name(), &result)
		if _, ok := result.GetQueryOk(); ok {
			id, err := ddObjects.CreateMonitors(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, &result)
			logOutput(err, "monitor", f.Name())
			if id != 0 {
				ids = append(ids, id)
			}
		} else {
			fmt.Printf("%v -- Not a valid monitor json.\n", f.Name())
		}
	}
	return ids

}

func LoadTimeboard(path string, orgCfg config.Organisation) []int {
	files := getFiles(path)
	var ids []int
	for _, f := range files {
		var result datadog.Screenboard
		unmarshalData(path, f.Name(), &result)
		id, err := ddObjects.CreateScreenboards(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, &result)
		logOutput(err, "screenboard", f.Name())
		if id != 0 {
			ids = append(ids, id)
		}
	}
	return ids
}
