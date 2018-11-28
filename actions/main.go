package actions

import (
	"fmt"
	"strings"

	"github.com/metrosystems-cpe/DDOM/config"
	"github.com/metrosystems-cpe/DDOM/helpers"

	"github.com/abiosoft/ishell"
	"github.com/metrosystems-cpe/DDOM/ddObjects"
	"github.com/metrosystems-cpe/DDOM/utils"
)

// Backup is used to store objects in file
func Backup(context *ishell.Context) {
	context.ShowPrompt(false)
	defer context.ShowPrompt(true)
	appCfg := context.Get("appConfig").(*utils.AppConfig)
	context.Print("Enter a DataDog organisation: ")
	org := context.ReadLine()
	context.Print("Enter a path where to store: ")
	outPath := context.ReadLine()
	if orgCfg, found := appCfg.OrgCfg.Find(org); found {
		list := ddObjects.List(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, appCfg.UsedObjectID)
		columns := []string{"ID", "Name"}
		table := helpers.BuildPrintableTable(columns, list)
		context.Println(table)
		context.Print("Enter comma separated ids to be stored: ")
		ids := context.ReadLine()
		saveData(outPath, ids, appCfg, orgCfg)
		// more backup logic here
	} else {
		context.Println("not a valid organisation")
	}
}

func LoadFromFile(context *ishell.Context) {
	appCfg := context.Get("appConfig").(*utils.AppConfig)
	context.Print("Enter a path from where to load the json: ")
	path := context.ReadLine()
	context.Print("Enter organization: ")
	org := context.ReadLine()
	if orgCfg, found := appCfg.OrgCfg.Find(org); found {
		fmt.Println(path)
		switch appCfg.UsedObjectID {
		case 1:
			// backup monitors
			loadMonitor(path, orgCfg)
		case 2:
			// backup dashboards
			loadDashboard(path, orgCfg)
		}

	} else {
		context.Println("Org not found")
	}

}

func Transfer(context *ishell.Context) {
	fmt.Println("will transfer")
}

func saveData(path string, ids string, appCfg *utils.AppConfig, orgCfg config.Organisation) {
	idsArr := strings.Split(ids, ",")
	switch appCfg.UsedObjectID {
	case 1:
		// backup monitors
		backupMonitors(path, idsArr, orgCfg)
	case 2:
		// backup dashboards
		backupDashboards(path, idsArr, orgCfg)
	case 3:
		// backup timeboards
		backupTimeboards(path, idsArr, orgCfg)
	}
}
