package ddObjects

import (
	"strconv"

	"github.com/metrosystems-cpe/DDOM/helpers"
	datadog "github.com/zorkian/go-datadog-api"
)

func buildDDClient(APIKey string, AppKey string, orgURL string) *datadog.Client {
	dClient := datadog.NewClient(APIKey, AppKey)
	dClient.SetBaseUrl(orgURL)
	return dClient
}

// List datadog objects
func List(APIKey string, AppKey string, orgURL string, objType uint) [][]interface{} {
	dClient := buildDDClient(APIKey, AppKey, orgURL)
	var result [][]interface{}

	switch objType {
	case 1:
		// monitors
		monitors, err := dClient.GetMonitors()
		helpers.LogError(err)
		for _, monitor := range monitors {
			// creator := monitor.GetCreator()
			result = append(result, []interface{}{strconv.Itoa(monitor.GetId()), monitor.GetName()})
		}
	case 2:
		// Dashboards
		dashboards, err := dClient.GetDashboards()
		helpers.LogError(err)
		for _, dash := range dashboards {
			result = append(result, []interface{}{strconv.Itoa(dash.GetId()), dash.GetTitle()})
		}
	case 3:
		// timeboards
		timeboards, err := dClient.GetScreenboards()
		helpers.LogError(err)
		for _, t := range timeboards {
			result = append(result, []interface{}{strconv.Itoa(t.GetId()), t.GetTitle()})
		}
	}
	return result
}

// MonitorDetails retreives an monitor based on it's ID
func MonitorDetails(APIKey string, AppKey string, orgURL string, monitorID string) *datadog.Monitor {
	client := buildDDClient(APIKey, AppKey, orgURL)
	ID, err := strconv.Atoi(monitorID)
	helpers.LogError(err)
	mon, err := client.GetMonitor(ID)
	helpers.LogError(err)
	return mon
}

// DashboardDetails retreives an monitor based on it's ID
func DashboardDetails(APIKey string, AppKey string, orgURL string, dashID string) *datadog.Dashboard {
	client := buildDDClient(APIKey, AppKey, orgURL)
	ID, err := strconv.Atoi(dashID)
	helpers.LogError(err)
	dash, err := client.GetDashboard(ID)
	helpers.LogError(err)
	return dash
}

// TimeboardDetails retreives an monitor based on it's ID
func TimeboardDetails(APIKey string, AppKey string, orgURL string, timeID string) (*datadog.Screenboard, error) {
	client := buildDDClient(APIKey, AppKey, orgURL)
	ID, err := strconv.Atoi(timeID)
	helpers.LogError(err)
	tb, err := client.GetScreenboard(ID)
	helpers.LogError(err)
	return tb, err
}

// CreateDashboards will push a dashboard to the dest organisation
func CreateDashboards(APIKey string, AppKey string, orgURL string, dash *datadog.Dashboard) error {
	client := buildDDClient(APIKey, AppKey, orgURL)
	_, err := client.CreateDashboard(dash)

	return err
}

// CreateMonitors will push a monitor to the dest organisation
func CreateMonitors(APIKey string, AppKey string, orgURL string, mon *datadog.Monitor) error {
	client := buildDDClient(APIKey, AppKey, orgURL)
	_, err := client.CreateMonitor(mon)

	return err
}

// CreateScreenboards will push a screenboard to the dest organisation
func CreateScreenboards(APIKey string, AppKey string, orgURL string, screen *datadog.Screenboard) error {
	client := buildDDClient(APIKey, AppKey, orgURL)
	_, err := client.CreateScreenboard(screen)

	return err
}
