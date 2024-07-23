package render

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderSnake() {
	tubeVertical(0, 0)
	tubeHorizontal(1, 0)
	cornerLeftTop(0, 1)
	cornerLeftDown(1, 1)
	cornerRightTop(2, 1)
	cornerRightDown(3, 1)
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
