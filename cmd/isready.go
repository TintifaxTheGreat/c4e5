package cmd

import (
	"github.com/abiosoft/ishell/v2"
)

var isready = func(c *ishell.Context) {
	sendReadyOk(c)
}
