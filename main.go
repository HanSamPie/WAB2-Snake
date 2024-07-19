package main

import (
	"projects/game"
	"projects/render"
)

func main() {
	gameState := game.InitGame(15, 15, true)
	render.Main(gameState)
}
