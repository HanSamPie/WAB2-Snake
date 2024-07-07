package main

import (
	"fmt"
	"projects/game"
	"projects/render"
)

func main() {
	fmt.Println("hello there")

	game.Columns = 8
	game.Rows = 8

	game.InitGame()

	render.Main(&game.Columns, &game.Rows)
}
