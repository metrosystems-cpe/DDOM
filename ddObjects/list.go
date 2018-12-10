package ddObjects

import (
	"fmt"
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
func MonitorDetails(APIKey string, AppKey string, orgURL string, monitorID string) (*datadog.Monitor, error) {
	client := buildDDClient(APIKey, AppKey, orgURL)
	ID, err := strconv.Atoi(monitorID)
	helpers.LogError(err)
	mon, err := client.GetMonitor(ID)
	helpers.LogError(err)
	return mon, err
}

// DashboardDetails retreives an monitor based on it's ID
func DashboardDetails(APIKey string, AppKey string, orgURL string, dashID string) (*datadog.Dashboard, error) {
	client := buildDDClient(APIKey, AppKey, orgURL)
	ID, err := strconv.Atoi(dashID)
	helpers.LogError(err)
	dash, err := client.GetDashboard(ID)
	helpers.LogError(err)
	return dash, err
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
func CreateDashboards(APIKey string, AppKey string, orgURL string, dash *datadog.Dashboard) (int, error) {
	client := buildDDClient(APIKey, AppKey, orgURL)
	actual, err := client.CreateDashboard(dash)
	id := actual.GetId()
	if id != 0 {
		fmt.Printf("Dashboard with id %v was created.\n", id)
	}
	return id, err
}

// CreateMonitors will push a monitor to the dest organisation
func CreateMonitors(APIKey string, AppKey string, orgURL string, mon *datadog.Monitor) (int, error) {
	client := buildDDClient(APIKey, AppKey, orgURL)
	actual, err := client.CreateMonitor(mon)
	id := actual.GetId()
	if id != 0 {
		fmt.Printf("Monitor with id %v was created.\n", id)
	}
	return id, err
}

// CreateScreenboards will push a screenboard to the dest organisation
func CreateScreenboards(APIKey string, AppKey string, orgURL string, screen *datadog.Screenboard) (int, error) {
	client := buildDDClient(APIKey, AppKey, orgURL)
	actual, err := client.CreateScreenboard(screen)
	id := actual.GetId()
	if id != 0 {
		fmt.Printf("Monitor with id %v was created.\n", id)
	}
	return id, err
}

func DeleteDashboard(APIKey string, AppKey string, orgURL string, id int) {
	client := buildDDClient(APIKey, AppKey, orgURL)
	err := client.DeleteDashboard(id)
	if err != nil {
		fmt.Println("Dashboard could not be deleted")
	} else {
		fmt.Printf("Dashboard with id %v was deleted\n", id)
	}
}

func DeleteMonitor(APIKey string, AppKey string, orgURL string, id int) {
	client := buildDDClient(APIKey, AppKey, orgURL)
	err := client.DeleteMonitor(id)
	if err != nil {
		fmt.Println("Monitor could not be deleted")
	} else {
		fmt.Printf("Monitor with id %v was deleted\n", id)
	}
}

func DeleteTimeboard(APIKey string, AppKey string, orgURL string, id int) {
	client := buildDDClient(APIKey, AppKey, orgURL)
	err := client.DeleteScreenboard(id)
	if err != nil {
		fmt.Println("Timeboard could not be deleted")
	} else {
		fmt.Printf("Timeboard with id %v was deleted\n", id)
	}
}
