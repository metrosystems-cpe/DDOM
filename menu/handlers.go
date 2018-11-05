package menu

import (
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/metrosystems-cpe/DDOM/constants"
	"github.com/metrosystems-cpe/DDOM/helpers"
	"github.com/metrosystems-cpe/DDOM/utils"
)

func checkOrgHandler(c *ishell.Context) {
	if len(c.Args) == 0 || len(c.Args) > 1 {
		c.Println("Please provide a single organisation name")
		return
	}
	organisationName := c.Args[0]
	c.Printf("%s: \n\n", organisationName)
	confRef := c.Get("appConfig").(*utils.AppConfig)
	v, f := confRef.OrgCfg.Find(c.Args[0])
	if f {
		row := [][]interface{}{{v.URL, v.APIKey, v.AppKey}}
		c.Println(helpers.BuildPrintableTable([]string{"URL", "API KEY", "APP KEY"}, row))
	} else {
		c.Println("Organisation not found")
	}

}

func setObjectHandler(c *ishell.Context) {
	var (
		rawInput int
		err      error
	)
	refConfig := c.Get("appConfig").(*utils.AppConfig)
	c.Print("Use 1 for Monitors or 2 for Dashboards or 3 for Timeboards ")
	rawInput, err = strconv.Atoi(c.ReadLine())
	for err != nil {
		c.Println("Not a valid option. Use 1 / 2 / 3")
		rawInput, err = strconv.Atoi(c.ReadLine())
	}
	refConfig.UsedObjectID = uint(rawInput)
	c.Printf("Context set to %s\n", constants.Objects[refConfig.UsedObjectID])
}

func setMethodHandler(c *ishell.Context) {
	var (
		rawInput int
		err      error
	)
	refConfig := c.Get("appConfig").(*utils.AppConfig)

	if refConfig.UsedObjectID == 0 {
		c.Println("Context not set. Run setContext")
		return
	}
	c.Printf("1 for Backup\n2 for Transfer\n3 for LoadFromFile\n")
	rawInput, err = strconv.Atoi(c.ReadLine())
	for err != nil {
		c.Println("Not a valid option. Use 1 / 2 / 3")
		rawInput, err = strconv.Atoi(c.ReadLine())
	}
	refConfig.Method = uint(rawInput)
	c.Printf("Method set to %s\n", constants.Methods[refConfig.Method])
}

func runHandler(c *ishell.Context) {
	refConfig := c.Get("appConfig").(*utils.AppConfig)
	c.Printf("The app will run in context of %s and will perform %s\n", constants.Objects[refConfig.UsedObjectID], constants.Methods[refConfig.Method])
}
