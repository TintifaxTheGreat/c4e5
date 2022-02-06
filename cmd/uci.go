package cmd

import (
	"github.com/abiosoft/ishell/v2"
)

var uci = func(c *ishell.Context) {
	sendId(c)
	sendOptions(c)
	sendUciOk(c)
}
