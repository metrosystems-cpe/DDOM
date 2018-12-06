package config

// Organisation is a data type used to represent an entry from the config organisations list
type Organisation struct {
	AppKey string `yaml:"AppKey"`
	APIKey string `yaml:"APIKey"`
	URL    string `yaml:"URL"`
}

// OrganisationList is a simple hash table to store configs for organisations
type OrganisationList map[string]Organisation

// Find is used to search for a name in the organisation list
func (list OrganisationList) Find(key string) (val Organisation, found bool) {
	val, found = list[key]
	return
}
