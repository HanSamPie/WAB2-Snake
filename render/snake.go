package render

import (
	"projects/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	vertical = iota
	horizontal
	leftTop
	leftDown
	rightTop
	rightDown
)

func renderSnake() {
	tubeVertical(0, 0)
	tubeHorizontal(1, 0)
	cornerLeftTop(0, 1)
	cornerLeftDown(1, 1)
	cornerRightTop(2, 1)
	cornerRightDown(3, 1)

	for i, part := range gameState.Snake {
		neighbour := checkNeighbours(part.X, part.Y)
		if i == 1 {
			if neighbour == horizontal {
				tubeVertical(int32(part.X), int32(part.Y))
			} else {
				tubeHorizontal(int32(part.X), int32(part.Y))
			}
		}
	}
}

func checkNeighbours(x int, y int) {

	if gameState.Grid[y][x-1] == game.SNAKE && gameState.Grid[y+1][x] == game.SNAKE && (x-1 >= 0 || y+1 <= gameState.Columns) {
		cornerLeftTop(int32(x), int32(y))
	}
}

func tubeVertical(x int32, y int32) {
	rl.DrawRectangle(((x+1)*rectangleSize + int32(int(rectangleSize)*3/10)), (y+1)*rectangleSize, int32(int(rectangleSize)*2/5), rectangleSize, rl.Black)
}

func tubeHorizontal(x int32, y int32) {
	rl.DrawRectangle(((x + 1) * rectangleSize), (y+1)*rectangleSize+int32(int(rectangleSize)*3/10), rectangleSize, int32(int(rectangleSize)*2/5), rl.Black)
}

func cornerLeftTop(x int32, y int32) {
	// leftCenter
	rl.DrawRectangle((x+1)*rectangleSize, (y+1)*rectangleSize+int32(int(rectangleSize)*3/10), int32(int(rectangleSize)*7/10), int32(int(rectangleSize)*2/5), rl.Black)
	// topCenter
	rl.DrawRectangle((x+1)*rectangleSize+int32(int(rectangleSize)*3/10), (y+1)*rectangleSize, int32(int(rectangleSize)*2/5), int32(int(rectangleSize)*7/10), rl.Black)
}

func cornerLeftDown(x int32, y int32) {
	// leftCenter
	rl.DrawRectangle((x+1)*rectangleSize, (y+1)*rectangleSize+int32(int(rectangleSize)*3/10), int32(int(rectangleSize)*7/10), int32(int(rectangleSize)*2/5), rl.Black)
	// downCenter
	rl.DrawRectangle((x+1)*rectangleSize+int32(int(rectangleSize)*3/10), (y+1)*rectangleSize+int32(int(rectangleSize*3/10)), int32(int(rectangleSize)*2/5), int32(int(rectangleSize)*7/10), rl.Black)
}

func cornerRightTop(x int32, y int32) {
	// topCenter
	rl.DrawRectangle((x+1)*rectangleSize+int32(int(rectangleSize)*3/10), (y+1)*rectangleSize, int32(int(rectangleSize)*2/5), int32(int(rectangleSize)*7/10), rl.Black)
	// rightCenter
	rl.DrawRectangle((x+1)*rectangleSize+int32(int(rectangleSize*3/10)), (y+1)*rectangleSize+int32(int(rectangleSize)*3/10), int32(int(rectangleSize)*7/10), int32(int(rectangleSize)*2/5), rl.Black)
}

func cornerRightDown(x int32, y int32) {
	// rightCenter
	rl.DrawRectangle((x+1)*rectangleSize+int32(int(rectangleSize*3/10)), (y+1)*rectangleSize+int32(int(rectangleSize)*3/10), int32(int(rectangleSize)*7/10), int32(int(rectangleSize)*2/5), rl.Black)
	// downCenter
	rl.DrawRectangle((x+1)*rectangleSize+int32(int(rectangleSize)*3/10), (y+1)*rectangleSize+int32(int(rectangleSize*3/10)), int32(int(rectangleSize)*2/5), int32(int(rectangleSize)*7/10), rl.Black)
}
