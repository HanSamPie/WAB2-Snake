package render

import (
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
)

func drawScene() {
	//for switch between menu etc
}

func input() {
	if (rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp)) && gameState.CurrentDirection != game.Down {
		gameState.CurrentDirection = game.Up
	} else if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) && gameState.CurrentDirection != game.Right {
		gameState.CurrentDirection = game.Left
	} else if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) && gameState.CurrentDirection != game.Up {
		gameState.CurrentDirection = game.Down
	} else if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) && gameState.CurrentDirection != game.Left {
		gameState.CurrentDirection = game.Right
	}
}

func render() {
	rl.BeginDrawing()

	//TODO add border to map and move score into it
	//TODO implement game over screen
	//TODO render snake based on gameState.Snake
	rl.ClearBackground(rl.DarkBlue)
	for rows := 0; rows < gameState.Rows; rows++ {
		for columns := 0; columns < gameState.Columns; columns++ {
			if gameState.Grid[rows][columns] == game.SNAKE {
				rl.DrawRectangle(int32(columns)*rectangleSize+50, int32(rows)*rectangleSize+50, rectangleSize, rectangleSize, rl.Black)
			} else if gameState.Grid[rows][columns] == game.FOOD {
				rl.DrawRectangle(int32(columns)*rectangleSize+50, int32(rows)*rectangleSize+50, rectangleSize, rectangleSize, rl.Red)
			} else if (rows+columns)%2 == 0 {
				rl.DrawRectangle(int32(columns)*rectangleSize+50, int32(rows)*rectangleSize+50, rectangleSize, rectangleSize, rl.Green)
			} else {
				rl.DrawRectangle(int32(columns)*rectangleSize+50, int32(rows)*rectangleSize+50, rectangleSize, rectangleSize, rl.Lime)
			}
		}
	}
	score := "length: " + strconv.Itoa(len(gameState.Snake))
	rl.DrawText(score, 15, 15, rectangleSize/3, rl.White)

	drawScene()

	rl.EndDrawing()
}

func update() {
	running = !rl.WindowShouldClose()
	frameCount++
	if frameCount == FPS/3 && gameState.CurrentDirection != game.GameOver {
		game.MoveSnake()
		frameCount = 0
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
