package actions

import (
	"github.com/metrosystems-cpe/DDOM/config"
	"github.com/metrosystems-cpe/DDOM/ddObjects"
)

type objectRetriever func(orgCfg config.Organisation, id string) (interface{}, error)

func getMonitor(orgCfg config.Organisation, id string) (interface{}, error) {
	return ddObjects.MonitorDetails(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, id)
}

func getDashboard(orgCfg config.Organisation, id string) (interface{}, error) {
	return ddObjects.DashboardDetails(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, id)
}

func getScreenboard(orgCfg config.Organisation, id string) (interface{}, error) {
	return ddObjects.TimeboardDetails(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, id)
}
