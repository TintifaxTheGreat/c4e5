package main

import (
	"github.com/tintifaxthegreat/c4e5/cmd"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.Println("engine started")

	cmd.Execute()
}
