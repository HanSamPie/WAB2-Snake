package main

import (
	"fmt"
	"projects/game"
	"projects/render"
)

func main() {
	fmt.Println("hello there")

	game.Columns = 9
	game.Rows = 9

	render.Main(&game.Columns, &game.Rows)
}
