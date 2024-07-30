package main

import (
	"projects/game"
	"projects/render"
)

// for MLP create head in game and pass direction to MLP
func main() {
	game := new(game.Game)
	game.InitGame(15, 15, false)
	render.Main(game)
}
