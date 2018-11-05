package actions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/metrosystems-cpe/DDOM/config"
	"github.com/metrosystems-cpe/DDOM/ddObjects"
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

func backupMonitors(path string, ids []string, orgCfg config.Organisation) {
	for _, id := range ids {
		monitor := ddObjects.MonitorDetails(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, id)
		monOut, _ := json.Marshal(monitor)
		fileName := fmt.Sprintf("%s/%d.json", path, monitor.GetId())
		writeToFile(fileName, monOut)
	}
}

func backupDashboards(path string, ids []string, orgCfg config.Organisation) {
	for _, id := range ids {
		dash := ddObjects.DashboardDetails(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, id)
		dashOut, _ := json.Marshal(dash)
		fileName := fmt.Sprintf("%s/%d.json", path, dash.GetId())
		writeToFile(fileName, dashOut)
	}
}

func backupTimeboards(path string, ids []string, orgCfg config.Organisation) {
	for _, id := range ids {
		tb := ddObjects.TimeboardDetails(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, id)
		tbOut, _ := json.Marshal(tb)
		fileName := fmt.Sprintf("%s/%d.json", path, tb.GetId())
		writeToFile(fileName, tbOut)
	}
}
