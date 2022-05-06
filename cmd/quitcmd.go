package cmd

import (
	"github.com/abiosoft/ishell/v2"
	"os"
)

var quitcmd = func(c *ishell.Context) {
	os.Exit(0)
}
