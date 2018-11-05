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
