package cmd

import (
	"github.com/abiosoft/ishell/v2"
	"github.com/tintifaxthegreat/c4e5/engine"
)

var game *engine.Game

func Execute() {
	game = engine.NewGame("")

	var shell *ishell.Shell
	shell = ishell.New()
	shell.ShowPrompt(false)

	shell.AddCmd(&ishell.Cmd{
		Name: "uci",
		Func: uci,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "isready",
		Func: isready,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "position",
		Func: position,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "go",
		Func: gocmd,
	})

	// TODO add quit

	shell.Run()
}
