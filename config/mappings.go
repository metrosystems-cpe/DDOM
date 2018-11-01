package config

type Organisation struct {
	AppKey string
	APIKey string
	URL    string
}

type OrganisationList map[string]Organisation

func (list OrganisationList) Find(key string) (val Organisation, found bool) {
	val, found = list[key]
	return
}
