package cmd

import (
	"github.com/abiosoft/ishell/v2"
)

func sendId(c *ishell.Context) {
	c.Println("id name C4E5")
	c.Println("id author Eugen Lindorfer")
}

func sendUciOk(c *ishell.Context) {
	c.Println("uciok")
}

func sendOptions(c *ishell.Context) {
	c.Println("option") //TODO extend this
}

func sendReadyOk(c *ishell.Context) {
	c.Println("readyok")
}

func debug(c *ishell.Context, s string) {
	c.Println("DEBUG ", s)
}
