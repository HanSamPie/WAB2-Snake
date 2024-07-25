package render

import (
	"log"
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

	//TODO shorten head and add eyes
	//TODO visualize game over on snake somehow

	//TODO game crashes on collision

	//TODO on game over render last state
	//if gameState.CurrentDirection == game.GameOver {
	//	return
	//}

	for i, part := range gameState.Snake {
		neighbor := checkNeighbors(i)
		if neighbor == horizontal {
			tubeHorizontal(int32(part.X), int32(part.Y))
		} else if neighbor == vertical {
			tubeVertical(int32(part.X), int32(part.Y))
		} else if neighbor == leftTop {
			cornerLeftTop(int32(part.X), int32(part.Y))
		} else if neighbor == leftDown {
			cornerLeftDown(int32(part.X), int32(part.Y))
		} else if neighbor == rightTop {
			cornerRightTop(int32(part.X), int32(part.Y))
		} else if neighbor == rightDown {
			cornerRightDown(int32(part.X), int32(part.Y))
		} else if neighbor == -1 && gameState.CurrentDirection != game.GameOver {
			log.Fatal("render snake neighbor -1; or in other words I fucked up")
		}
	}
}

func checkNeighbors(n int) int {
	if n == 0 {
		//case head
		if lastDirection == game.Up || lastDirection == game.Down {
			return vertical
		} else if lastDirection == game.Left || lastDirection == game.Right {
			return horizontal
		}
	} else if len(gameState.Snake) >= 2 && len(gameState.Snake)-1 == n {
		//case tail
		part0 := gameState.Snake[n]
		previousPart := gameState.Snake[n-1]
		diffPrevious := game.Direction{X: part0.X - previousPart.X, Y: part0.Y - previousPart.Y}

		if diffPrevious.X == 1 || diffPrevious.X == -1 {
			return horizontal
		} else if diffPrevious.Y == 1 || diffPrevious.Y == -1 {
			return vertical
		}
	} else {
		//case body
		part0 := gameState.Snake[n]
		previousPart := gameState.Snake[n-1]
		nextPart := gameState.Snake[n+1]

		diffPrevious := game.Direction{X: part0.X - previousPart.X, Y: part0.Y - previousPart.Y}
		diffNext := game.Direction{X: part0.X - nextPart.X, Y: part0.Y - nextPart.Y}

		if (diffPrevious.X == 1 && diffNext.X == -1) || (diffPrevious.X == -1 && diffNext.X == 1) {
			//case horizontal; both right to left and left to right
			return horizontal
		} else if (diffPrevious.Y == 1 && diffNext.Y == -1) || (diffPrevious.Y == -1 && diffNext.Y == 1) {
			//case vertical; both up to down and down to up
			return vertical
		} else if (diffPrevious.X == 1 && diffNext.Y == 1) || (diffPrevious.Y == 1 && diffNext.X == 1) {
			//case corner leftTop
			return leftTop
		} else if (diffPrevious.X == 1 && diffNext.Y == -1) || (diffPrevious.Y == -1 && diffNext.X == 1) {
			//case leftDown
			return leftDown
		} else if (diffPrevious.X == -1 && diffNext.Y == 1) || (diffPrevious.Y == 1 && diffNext.X == -1) {
			//case rightTop
			return rightTop
		} else if (diffPrevious.X == -1 && diffNext.Y == -1) || (diffPrevious.Y == -1 && diffNext.X == -1) {
			//case rightDown
			return rightDown
		}
	}
	return -1
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
