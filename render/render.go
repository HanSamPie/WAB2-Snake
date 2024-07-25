package render

import (
	"fmt"
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
	gameState     *game.GameState
	lastDirection = game.Right
)

func input() {
	if (rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp)) && lastDirection.Y != 1 {
		gameState.CurrentDirection = game.Up
	} else if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) && lastDirection.X != 1 {
		gameState.CurrentDirection = game.Left
	} else if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) && lastDirection.Y != -1 {
		gameState.CurrentDirection = game.Down
	} else if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) && lastDirection.X != -1 {
		gameState.CurrentDirection = game.Right
	}
}

func renderBoard() {
	//renders boundary
	rl.ClearBackground(rl.DarkBlue)
	//renders checkered board
	for rows := 0; rows < gameState.Rows; rows++ {
		for columns := 0; columns < gameState.Columns; columns++ {
			/*if gameState.Grid[rows][columns] == game.SNAKE {
				rl.DrawRectangle(int32(columns+1)*rectangleSize, int32(rows+1)*rectangleSize, rectangleSize, rectangleSize, rl.Black)
			} else*/if gameState.Grid[rows][columns] == game.FOOD {
				rl.DrawRectangle(int32(columns+1)*rectangleSize, int32(rows+1)*rectangleSize, rectangleSize, rectangleSize, rl.Red)
			} else if (rows+columns)%2 == 0 {
				rl.DrawRectangle(int32(columns+1)*rectangleSize, int32(rows+1)*rectangleSize, rectangleSize, rectangleSize, rl.Green)
			} else {
				rl.DrawRectangle(int32(columns+1)*rectangleSize, int32(rows+1)*rectangleSize, rectangleSize, rectangleSize, rl.Lime)
			}
		}
	}
}

func render() {
	rl.BeginDrawing()

	//TODO implement game over screen
	renderBoard()

	renderSnake()

	//renders score
	//TODO maybe create render ui function
	score := "length: " + strconv.Itoa(len(gameState.Snake))
	rl.DrawText(score, 15, 15, rectangleSize/3, rl.White)

	rl.EndDrawing()
}

func update() {
	running = !rl.WindowShouldClose()
	frameCount++
	if frameCount == FPS/3 && gameState.CurrentDirection != game.GameOver {
		game.MoveSnake()

		frameCount = 0
		if gameState.CurrentDirection != game.GameOver {
			lastDirection = gameState.CurrentDirection
		}

		if gameState.Debug {
			fmt.Println(gameState.CurrentDirection)
		}

	}

}

func Main(state *game.GameState) {
	gameState = state

	screenWidth = 100 + (int32(gameState.Columns) * rectangleSize)
	screenHeight = 100 + (int32(gameState.Rows) * rectangleSize)
	rl.InitWindow(screenWidth, screenHeight, "Hello there")
	defer rl.CloseWindow()

	//rl.SetExitKey(0)

	rl.SetTargetFPS(int32(FPS))

	for running {
		//TODO on checks input on FPS
		//Try to set custom tick rate for update
		input()
		update()
		render()
	}
}
