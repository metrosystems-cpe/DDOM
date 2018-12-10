package main

import (
	"testing"

	"github.com/metrosystems-cpe/DDOM/actions"
	"github.com/metrosystems-cpe/DDOM/config"
	"github.com/metrosystems-cpe/DDOM/ddObjects"
	"github.com/metrosystems-cpe/DDOM/test_constants"
)

func TestLoadDashboard(t *testing.T) {

	path := "tests"

	var orgCfg config.Organisation
	orgCfg = test_constants.OrgCfg

	id := actions.LoadDashboard(path, orgCfg)
	for _, i := range id {
		ddObjects.DeleteDashboard(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, i)
	}

}

func TestLoadMonitor(t *testing.T) {
	path := "tests"

	var orgCfg config.Organisation
	orgCfg = test_constants.OrgCfg

	id := actions.LoadMonitor(path, orgCfg)
	for _, i := range id {
		ddObjects.DeleteMonitor(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, i)
	}
}

func TestLoadTimeboard(t *testing.T) {
	path := "tests"

	var orgCfg config.Organisation
	orgCfg = test_constants.OrgCfg

	id := actions.LoadTimeboard(path, orgCfg)
	for _, i := range id {
		ddObjects.DeleteTimeboard(orgCfg.APIKey, orgCfg.AppKey, orgCfg.URL, i)
	}
}
