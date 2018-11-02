package config

type Organisation struct {
	AppKey string `yaml:"AppKey"`
	APIKey string `yaml:"APIKey"`
	URL    string `yaml:"URL"`
}

type OrganisationList map[string]Organisation

func (list OrganisationList) Find(key string) (val Organisation, found bool) {
	val, found = list[key]
	return
}
