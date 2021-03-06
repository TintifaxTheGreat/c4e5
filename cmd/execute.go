package cmd

import (
	"github.com/abiosoft/ishell/v2"
	"github.com/tintifaxthegreat/c4e5/engine"
)

var game *engine.Game

func Execute() {
	game = engine.NewGame("", 0, 0, 0)

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

	shell.AddCmd(&ishell.Cmd{
		Name: "quit",
		Func: quitcmd,
	})

	shell.Run()
}
