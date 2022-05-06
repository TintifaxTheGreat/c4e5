package cmd

import (
	"github.com/tintifaxthegreat/c4e5/engine"
	"log"
	"time"
)

type timeManagement struct {
	wtime     int
	btime     int
	winc      int
	binc      int
	movesToGo int
}

func NewTimeManagement() *timeManagement {
	return &timeManagement{}
}

func (tm *timeManagement) SetGameTime(g *engine.Game) {
	var timeForAllMoves int
	var timeUsagePercent = timeUsagePercentDefault

	if tm.movesToGo == 0 {
		tm.movesToGo = estimatedRestMoves
	}

	if g.Board.Wtomove {
		timeForAllMoves = tm.wtime + (tm.movesToGo-1)*tm.winc
	} else {
		timeForAllMoves = tm.btime + (tm.movesToGo-1)*tm.binc
	}

	if g.Board.Fullmoveno < earlyGameMaxMoves && tm.movesToGo > earlyGameMinRestMoves {
		timeUsagePercent = timeUsagePercentEarlyGame
	}

	g.MoveTime = time.Duration((timeForAllMoves*timeUsagePercent)/(tm.movesToGo*100)) * time.Millisecond
	log.Print("time set to ", g.MoveTime)
}
