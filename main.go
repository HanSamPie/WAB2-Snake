package main

import (
	"fmt"
	"projects/game"
	"projects/render"
)

func initGame(columns int, rows int) *game.GameState {
	return game.InitGame(columns, rows)

}

func main() {
	fmt.Println("hello there")

	gameState := initGame(9, 9)

	render.Main(gameState)
}
