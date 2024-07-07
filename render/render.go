package render

import (
	"projects/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	running             = true
	rectangleSize int32 = 100
	screenWidth   int32
	screenHeight  int32
)

func drawScene() {
	//for switch between menu etc
}

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		game.CurrentDirection = game.Up
	} else if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		game.CurrentDirection = game.Left
	} else if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		game.CurrentDirection = game.Down
	} else if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		game.CurrentDirection = game.Right
	}
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	//breaks of rows || columns < 10
	for rows := 0; rows < int(screenHeight/rectangleSize); rows++ {
		for columns := 0; columns < int(screenWidth/rectangleSize); columns++ {
			if game.Grid[rows][columns] == game.SNAKE {
				rl.DrawRectangle(int32(columns)*rectangleSize, int32(rows)*rectangleSize, rectangleSize, rectangleSize, rl.Black)
			} else if game.Grid[rows][columns] == game.FOOD {
				rl.DrawRectangle(int32(columns)*rectangleSize, int32(rows)*rectangleSize, rectangleSize, rectangleSize, rl.Red)
			} else if (rows+columns)%2 == 0 {
				rl.DrawRectangle(int32(columns)*rectangleSize, int32(rows)*rectangleSize, rectangleSize, rectangleSize, rl.Green)
			} else {
				rl.DrawRectangle(int32(columns)*rectangleSize, int32(rows)*rectangleSize, rectangleSize, rectangleSize, rl.Lime)
			}
		}
	}
	drawScene()

	rl.EndDrawing()
}

func update() {
	running = !rl.WindowShouldClose()

}

func Main(col *int, rows *int) {
	screenWidth = int32(*col) * rectangleSize
	screenHeight = int32(*rows) * rectangleSize
	rl.InitWindow(screenWidth, screenHeight, "Hello there")
	defer rl.CloseWindow()

	//rl.SetExitKey(0)
	rl.SetTargetFPS(60)
	for running {
		input()
		update()
		println(rl.IsKeyDown(rl.KeyD))
		render()
	}
}
