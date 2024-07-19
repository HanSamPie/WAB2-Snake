package main

import (
	"fmt"
	"projects/game"
	"projects/render"
)

func initGame(columns int, rows int, debug bool) *game.GameState {
	return game.InitGame(columns, rows, debug)
}

func main() {
	fmt.Println("hello there")

	gameState := initGame(9, 9, false)

	render.Main(gameState)
}
