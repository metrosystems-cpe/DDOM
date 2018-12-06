package actions

import (
	"fmt"

	"github.com/metrosystems-cpe/DDOM/config"
	"github.com/metrosystems-cpe/DDOM/ddObjects"
)

func transferDashboard(ids []string, sourceCfg config.Organisation, orgCfg config.Organisation) {
	for _, id := range ids {
		dash, _ := ddObjects.DashboardDetails(sourceCfg.APIKey, sourceCfg.AppKey, sourceCfg.URL, id)
		err := ddObjects.CreateDashboards(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, dash)
		if err != nil {
			fmt.Println(fmt.Sprintf("Could not create datadog dashboard to destionation organisation"))
		} else {
			fmt.Println(fmt.Sprintf("Dashboard created successfully to destionation organisation"))
		}
	}
}

func transferTimeboard(ids []string, sourceCfg config.Organisation, orgCfg config.Organisation) {
	for _, id := range ids {
		tb, err := ddObjects.TimeboardDetails(sourceCfg.APIKey, sourceCfg.AppKey, sourceCfg.URL, id)
		// dashOut, _ := json.Marshal(dash)
		// var result datadog.Dashboard
		// json.Unmarshal([]byte(byteValue), &result)
		err = ddObjects.CreateScreenboards(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, tb)
		if err != nil {
			fmt.Println(fmt.Sprintf("Could not create datadog timeboard to destionation organisation -- %v", err))
		} else {
			fmt.Println(fmt.Sprintf("Timeboard created successfully to destionation organisation"))
		}
	}
}

func transferMonitor(ids []string, sourceCfg config.Organisation, orgCfg config.Organisation) {
	for _, id := range ids {
		mon, _ := ddObjects.MonitorDetails(sourceCfg.APIKey, sourceCfg.AppKey, sourceCfg.URL, id)
		err := ddObjects.CreateMonitors(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, mon)
		if err != nil {
			fmt.Println(fmt.Sprintf("Could not create datadog monitor to destionation organisation"))
		} else {
			fmt.Println(fmt.Sprintf("Monitor created successfully to destionation organisation"))
		}
	}
}
