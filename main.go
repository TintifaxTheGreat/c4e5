package main

import (
	"github.com/tintifaxthegreat/c4e5/cmd"
	"log"
	"os"
)

func main() {
	file, err := openLogFile("/home/eugen/c4e5.log") //TODO make this configurable
	if err != nil {
		panic("cannot create logfile")
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	log.Println("engine started")

	cmd.Execute()
	/*
		start := time.Now()
		d := 6
		g := engine.NewGame("", 0, 0, 0)
		n := engine.Perft(&g.Board, d)
		duration := time.Since(start)
		log.Print("Perft ", d, ": ", n)
		log.Print("Duration: ", duration)

		start = time.Now()
		n = engine.PerftPrime(&g.Board, d)
		duration = time.Since(start)
		log.Print("PerftPrime ", d, ": ", n)
		log.Print("Duration: ", duration)

	*/
}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
