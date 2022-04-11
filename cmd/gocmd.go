package cmd

import (
	"github.com/abiosoft/ishell/v2"
	"strconv"
	"time"
)

var gocmd = func(c *ishell.Context) {
	var next, value string
	timeManagement := NewTimeManagement()

	commands := c.Args
	for hasNext(commands) {

		next, commands = getNext(commands)
		switch next {
		case "searchmoves":

		case "ponder":

		case "wtime":
			value, commands = getNext(commands)
			iValue, err := strconv.Atoi(value)
			if err == nil {
				timeManagement.tmWtime = iValue
			}

		case "btime":
			value, commands = getNext(commands)
			iValue, err := strconv.Atoi(value)
			if err == nil {
				timeManagement.tmBtime = iValue
			}

		case "winc":
			value, commands = getNext(commands)
			iValue, err := strconv.Atoi(value)
			if err == nil {
				timeManagement.tmWinc = iValue
			}

		case "binc":
			value, commands = getNext(commands)
			iValue, err := strconv.Atoi(value)
			if err == nil {
				timeManagement.tmBinc = iValue
			}

		case "movestogo":
			value, commands = getNext(commands)
			iValue, err := strconv.Atoi(value)
			if err == nil {
				timeManagement.tmMovesToGo = iValue
			}

		case "depth":
			value, commands = getNext(commands)
			iValue, err := strconv.Atoi(value)
			if err == nil {
				game.MaxDepth = iValue - 1
			}

		case "nodes":

		case "mate":

		case "movetime":
			value, commands = getNext(commands)
			iValue, err := strconv.Atoi(value)
			if err == nil {
				game.MoveTime = time.Duration(iValue) * time.Millisecond * 9 / 10
				goto startTimer // If movetime is given, there is no need to manage the time per game.
			}

		case "infinite":

		case "default":
		}
	}

	timeManagement.SetGameTime(game)

startTimer:
	timer := time.NewTimer(game.MoveTime)
	go func() {
		<-timer.C
		game.Playing = false
	}()
	game.Playing = true
	move := game.FindMove()
	game.Board.Apply(move)
	game.StoreBoardHistory()
	c.Println("bestmove", move.String())
}
