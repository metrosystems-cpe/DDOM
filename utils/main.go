package utils

import "github.com/metrosystems-cpe/DDOM/config"

// AppConfig is used to encapsulate needid things for the app
type AppConfig struct {
	UsedObjectID uint
	Method       uint
	OrgCfg       config.OrganisationList
}
