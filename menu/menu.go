package menu

import (
	"github.com/abiosoft/ishell"
)

// ComputeCommands is used to add commands to our interactive menu
func ComputeCommands(shell *ishell.Shell) {
	shell.AddCmd(&ishell.Cmd{
		Name: "checkOrg",
		Help: "Shows the organisation config",
		Func: checkOrgHandler,
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "setObject",
		Help: "Set the object you want to use",
		Func: setObjectHandler,
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "setMethod",
		Help: "Set the method to be performed",
		Func: setMethodHandler,
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "run",
		Help: "Run the selected method for selected context",
		Func: runHandler,
	})
}

// Run is used to show up the menu
func Run(shell *ishell.Shell) {
	shell.Run()
}
