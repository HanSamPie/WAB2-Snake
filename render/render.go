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
	rectangleSize int32 = 100
	FPS           int   = 90
	frameCount    int   = 0
	screenWidth   int32
	screenHeight  int32
	lastDirection = game.Right
	gameState     *game.Game
	renderState   int = 0
)

const (
	getInformation = iota
	startGame
	gameRunning
	gameOver
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

	//TODO for now setting it to running
	renderState = startGame
	switch renderState {
	case getInformation:

	case startGame:
		renderBoard()
		renderSnake()

		//TODO either move text or add see through filter over game

		start := "Press Enter to start the game!"
		fontSize := 20
		textWidth := rl.MeasureText(start, rectangleSize/2)
		textHeight := fontSize // Simplified; you might want to adjust based on actual font metrics

		// Calculating the center position
		x := (screenWidth / 2) - (textWidth / 2)
		y := (screenHeight / 2) - int32((textHeight / 2))
		rl.DrawText(start, x, y, rectangleSize/2, rl.Black)
		rl.DrawText("You can control the snake by using WASD or the Arrow keys.", 15, 15, rectangleSize/2, rl.Black)
		explanation := "Explanation:\n" +
			"The goal of the game is to make your snake as long as possible.\n" +
			"To increase your snakes length you have to eat the food (red square).\n" +
			"Crashing into yourself or the border will end the game."
		rl.DrawText(explanation, rectangleSize/2, screenHeight-3*rectangleSize, rectangleSize/3, rl.Black)
	case gameRunning:
		renderBoard()
		renderSnake()

		//renders score
		score := "length: " + strconv.Itoa(len(gameState.Snake))
		rl.DrawText(score, 15, 15, rectangleSize/2, rl.White)
	case gameOver:
		//TODO implement game over screen
	}

	rl.EndDrawing()
}

func update() {
	running = !rl.WindowShouldClose()
	//ensure that game is only running when the render state == gameRunning
	if renderState == gameRunning {
		frameCount++
		//define the number of updates per seconds
		if frameCount == FPS/6 && gameState.CurrentDirection != game.Stop {
			gameState.MoveSnake()

			frameCount = 0

			//change direction and add data to metrics
			if gameState.CurrentDirection != game.Stop && gameState.CurrentDirection != lastDirection {
				lastDirection = gameState.CurrentDirection

				//add element to DirectionChanges
				data := game.DirectionChange{
					Direction: game.DirectionMap[lastDirection],
					Timestamp: time.Now(),
				}
				gameState.Metrics.DirectionChanges = append(gameState.Metrics.DirectionChanges, data)
				game.NumberInputsToFruit++
			}

			if gameState.Debug {
				fmt.Println(gameState.CurrentDirection)
			}

		}
	}

}

func Main(state *game.Game) {
	gameState = state

	//TODO render height based on screen size
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
