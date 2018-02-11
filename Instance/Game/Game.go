package main

import (
	"fmt"
	"os"
	"os/signal"

	cons "../../ykconstant"
)

const (
	AboutEntity = cons.Entity_Game
)

func main() {
	fmt.Println("Game Server Start...")

	game := &Game{}
	game.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
