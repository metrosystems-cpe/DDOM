package constants

type Endpoint struct {
	Method string
	Url    string
}

// Objects is a mapping for the used objects in DD
var Objects = map[uint]string{
	1: "Monitors",
	2: "Dashboards",
	3: "Timeboards",
}

var Methods = map[uint]string{
	1: "Backup",
	2: "Transfer",
	3: "LoadFromFile",
}

const ConfigFilePath = "./ddom.yaml"

var DDAPIEndpoints = map[uint]interface{}{
	1: map[string]interface{}{
		"index":  Endpoint{Method: "GET", Url: "https://api.datadoghq.com/api/v1/monitor"},
		"create": Endpoint{Method: "POST", Url: "https://api.datadoghq.com/api/v1/monitor"},
		"update": Endpoint{Method: "PUT", Url: "https://api.datadoghq.com/api/v1/monitor/"},
	},
	2: map[string]interface{}{
		"index":  Endpoint{Method: "GET", Url: "https://api.datadoghq.com/api/v1/dashboard/lists/manual"},
		"create": Endpoint{Method: "POST", Url: "https://api.datadoghq.com/api/v1/dashboard/lists/manual"},
		"update": Endpoint{Method: "PUT", Url: "https://api.datadoghq.com/api/v1/dashboard/lists/manual/"},
	},
	3: map[string]interface{}{
		"index":  Endpoint{Method: "GET", Url: "https://api.datadoghq.com/api/v1/dash"},
		"create": Endpoint{Method: "POST", Url: "https://api.datadoghq.com/api/v1/dash"},
		"update": Endpoint{Method: "PUT", Url: "https://api.datadoghq.com/api/v1/dash/"},
	},
}
