package config

import (
	"io/ioutil"

	"github.com/metrosystems-cpe/DDOM/constants"
	"github.com/metrosystems-cpe/DDOM/helpers"
	yaml "gopkg.in/yaml.v2"
)

func LoadfromFile() OrganisationList {
	byteContent, err := ioutil.ReadFile(constants.ConfigFilePath)

	helpers.PanicIfError(err, "File could not be read")

	list := make(OrganisationList)
	err = yaml.Unmarshal(byteContent, &list)

	helpers.PanicIfError(err, "Corupted data")

	return list
}
