package actions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/metrosystems-cpe/DDOM/config"
	"github.com/metrosystems-cpe/DDOM/helpers"
)

func writeToFile(fileName string, data []byte) {
	outFile, err := os.Create(fileName)
	helpers.PanicIfError(err, "Could not create file")
	w := bufio.NewWriter(outFile)
	_, err = w.WriteString(string(data))
	helpers.PanicIfError(err, "Could Not write data into file")
	w.Flush()
}

func backup(path string, ids []string, orgCfg config.Organisation, r objectRetriever) {
	for _, id := range ids {
		obj, err := r(orgCfg, id)
		if err != nil {
			fmt.Printf("Object with ID: %v could not pe saved. Error: %v", id, err)
		} else {
			objOut, _ := json.Marshal(obj)
			fileName := fmt.Sprintf("%s/%s.json", path, id)
			writeToFile(fileName, objOut)
		}
	}
}

func backupMonitors(path string, ids []string, orgCfg config.Organisation) {
	backup(path, ids, orgCfg, getMonitor)
}

func backupDashboards(path string, ids []string, orgCfg config.Organisation) {
	backup(path, ids, orgCfg, getDashboard)
}

func backupTimeboards(path string, ids []string, orgCfg config.Organisation) {
	backup(path, ids, orgCfg, getScreenboard)
}
