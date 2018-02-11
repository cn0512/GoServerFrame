package main

import (
	"fmt"

	"../GamePlus"
)

type Game struct {
	logic GamePlus.IGamePlus
}

func (yk *Game) Start() {
	yk.logic = &GamePlus.GamePlus_Logic1{}
	fmt.Println(yk.logic.About())
	yk.logic.Init()
	yk.logic.Run()
}

func (yk *Game) Stop() {
	yk.logic.End()
}
