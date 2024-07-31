package render

import (
	"fmt"
	"image/color"
	"projects/game"
	"strconv"
	"time"

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
	gameState     *game.Game
)

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
		gameState.MoveSnake()

		frameCount = 0
		if gameState.CurrentDirection != game.Stop && gameState.CurrentDirection != lastDirection {
			lastDirection = gameState.CurrentDirection

			//add element to DirectionChanges
			data := game.DirectionChange{
				Direction: game.DirectionMap[lastDirection],
				Timestamp: time.Now(),
			}
			gameState.Metrics.DirectionChanges = append(gameState.Metrics.DirectionChanges, data)

		}

		if gameState.Debug {
			fmt.Println(gameState.CurrentDirection)
		}

	}

}

func Main(state *game.Game) {
	gameState = state

	screenWidth = 2*rectangleSize + (int32(gameState.Columns) * rectangleSize)
	screenHeight = 2*rectangleSize + (int32(gameState.Rows) * rectangleSize)
	rl.InitWindow(screenWidth, screenHeight, "Snake")
	defer rl.CloseWindow()

	//rl.SetExitKey(0)

	rl.SetTargetFPS(int32(FPS))

	for running {
		//or check for Stop on each move
		if gameState.CurrentDirection != game.Stop {
			input()
		}
		update()
		render()
	}
}
