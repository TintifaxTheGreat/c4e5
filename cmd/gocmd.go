package cmd

import (
	"github.com/abiosoft/ishell/v2"
	"time"
)

var gocmd = func(c *ishell.Context) {
	timer := time.NewTimer(game.MoveTime)
	go func() {
		<-timer.C
		game.Playing = false
	}()
	game.Playing = true
	move := game.FindMove()
	c.Println("bestmove", move.String())
}
