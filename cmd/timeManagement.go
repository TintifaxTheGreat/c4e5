package cmd

import (
	"github.com/tintifaxthegreat/c4e5/engine"
	"log"
	"time"
)

type timeManagement struct {
	tmWtime     int
	tmBtime     int
	tmWinc      int
	tmBinc      int
	tmMovesToGo int
}

func NewTimeManagement() *timeManagement {
	return &timeManagement{}
}

func (tm *timeManagement) SetGameTime(g *engine.Game) {
	var timeForAllMoves int

	if tm.tmMovesToGo == 0 {
		tm.tmMovesToGo = 40
	}

	if g.Board.Wtomove {
		timeForAllMoves = tm.tmWtime + (tm.tmMovesToGo-1)*tm.tmWinc
	} else {
		timeForAllMoves = tm.tmBtime + (tm.tmMovesToGo-1)*tm.tmBinc
	}

	g.MoveTime = time.Duration((timeForAllMoves*75/100)/tm.tmMovesToGo) * time.Millisecond
	log.Print("time set to ", g.MoveTime)
}
