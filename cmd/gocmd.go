package cmd

import (
	"github.com/abiosoft/ishell/v2"
)

var gocmd = func(c *ishell.Context) {
	m := game.FindMove()
	c.Println("bestmove", m.String())
}
