package main

import (
	"github.com/abiosoft/ishell"
	"github.com/metrosystems-cpe/DDOM/config"
	"github.com/metrosystems-cpe/DDOM/menu"
	"github.com/metrosystems-cpe/DDOM/utils"
)

var (
	shell     *ishell.Shell
	appConfig utils.AppConfig
)

func init() {
	shell = ishell.New()
	shell.SetHistoryPath("./history")
	appConfig.OrgCfg = config.LoadfromFile()
	shell.Set("appConfig", &appConfig)
}

func main() {
	menu.ComputeCommands(shell)
	menu.Run(shell)
}
