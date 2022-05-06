package cmd

import (
	"github.com/abiosoft/ishell/v2"
	"github.com/dylhunn/dragontoothmg"
	"strings"
)

var position = func(c *ishell.Context) {

	var fen string
	position, args := c.Args[0], c.Args[1:]

	switch position {
	case "fen":
		fen = strings.Join(args, " ")
	case "startpos":
		fen = fenStart
	default:
		panic("unknown keyword after position")
	}

	// TODO consider moves

	game.Board = dragontoothmg.ParseFen(fen)
	game.StoreBoardHistory()
}
