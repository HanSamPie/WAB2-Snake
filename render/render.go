package render

import (
	"fmt"
	"image/color"
	"projects/game"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	running             = true
	rectangleSize int32 = 50
	FPS           int   = 90
	frameCount    int   = 0
	screenWidth   int32
	screenHeight  int32
	lastDirection = game.Right
)

type renderer struct {
}

func input() {
	if (rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp)) && lastDirection.Y != 1 {
		//change direction to up
		gameState.CurrentDirection = game.Up
	} else if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) && lastDirection.X != 1 {
		//change direction to left
		gameState.CurrentDirection = game.Left
	} else if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) && lastDirection.Y != -1 {
		//change direction to down
		gameState.CurrentDirection = game.Down
	} else if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) && lastDirection.X != -1 {
		//change direction to right
		gameState.CurrentDirection = game.Right
	}
}

func renderBoard() {
	//renders boundary
	rl.ClearBackground(rl.DarkBlue)
	//renders checkered board
	for rows := 0; rows < gameState.Rows; rows++ {
		for columns := 0; columns < gameState.Columns; columns++ {
			var color color.RGBA
			if gameState.Grid[rows][columns] == game.FOOD {
				color = rl.Red
			} else if (rows+columns)%2 == 0 {
				color = rl.Green
			} else {
				color = rl.Lime
			}
			rl.DrawRectangle(int32(columns+1)*rectangleSize, int32(rows+1)*rectangleSize, rectangleSize, rectangleSize, color)
		}
	}
}

func render() {
	rl.BeginDrawing()

	//TODO implement game over screen
	renderBoard()
	renderSnake()

	//renders score
	score := "length: " + strconv.Itoa(len(gameState.Snake))
	rl.DrawText(score, 15, 15, rectangleSize/2, rl.White)

	rl.EndDrawing()
}

func update() {
	running = !rl.WindowShouldClose()
	frameCount++
	if frameCount == FPS/6 && gameState.CurrentDirection != game.Stop {
		game.MoveSnake()

		frameCount = 0
		if gameState.CurrentDirection != game.Stop {
			lastDirection = gameState.CurrentDirection
		}

		if gameState.Debug {
			fmt.Println(gameState.CurrentDirection)
		}

	}

}

func Renderer() {
	gameState = state

	screenWidth = 2*rectangleSize + (int32(gameState.Columns) * rectangleSize)
	screenHeight = 2*rectangleSize + (int32(gameState.Rows) * rectangleSize)
	rl.InitWindow(screenWidth, screenHeight, "Snake")
	defer rl.CloseWindow()

	//rl.SetExitKey(0)

	rl.SetTargetFPS(int32(FPS))

	for running {
		input()
		update()
		render()
	}
}
