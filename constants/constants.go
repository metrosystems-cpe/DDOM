package constants

// Objects is a mapping for the used objects in DD
var Objects = map[uint]string{
	1: "Monitors",
	2: "Dashboards",
	3: "Timeboards",
}

// Methods mapping for the available methods
var Methods = map[uint]string{
	1: "Backup",
	2: "Transfer",
	3: "LoadFromFile",
}

// ConfigFilePath where organisation configs are stored
const ConfigFilePath = "./ddom.yaml"
