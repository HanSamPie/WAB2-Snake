package render

import (
	"log"
	"projects/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	headVertical = iota
	headHorizontal
	vertical
	horizontal
	leftTop
	leftDown
	rightTop
	rightDown
)

func renderSnake() {
	//TODO fix bracket, typecasting and general math hell in this file
	//TODO make tail smooth
	for i, part := range gameState.Snake {
		neighbor := checkNeighbors(i)
		x := int32(part.X + 1)
		y := int32(part.Y + 1)
		if neighbor == headVertical {
			tubeHeadVertical(x, y)
		} else if neighbor == headHorizontal {
			tubeHeadHorizontal(x, y)
		} else if neighbor == horizontal {
			tubeHorizontal(x, y)
		} else if neighbor == vertical {
			tubeVertical(x, y)
		} else if neighbor == leftTop {
			cornerLeftTop(x, y)
		} else if neighbor == leftDown {
			cornerLeftDown(x, y)
		} else if neighbor == rightTop {
			cornerRightTop(x, y)
		} else if neighbor == rightDown {
			cornerRightDown(x, y)
		} else if neighbor == -1 && gameState.CurrentDirection != game.GameOver {
			log.Fatal("render snake neighbor -1; or in other words I fucked up")
		}
	}
}

func checkNeighbors(n int) int {
	if n == 0 {
		//case head
		if lastDirection == game.Up || lastDirection == game.Down {
			return headVertical
		} else if lastDirection == game.Left || lastDirection == game.Right {
			return headHorizontal
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

func tubeHeadVertical(x int32, y int32) {
	if lastDirection == game.Down {
		rl.DrawRectangle((x*rectangleSize + rectangleSize*3/10), y*rectangleSize, rectangleSize*2/5, rectangleSize/2, rl.Black)
		//eyes
		rl.DrawRectangle((x*rectangleSize + rectangleSize*3/10), y*rectangleSize, rectangleSize/10, rectangleSize/10, rl.White)
		rl.DrawRectangle(x*rectangleSize+rectangleSize*6/10, y*rectangleSize, rectangleSize/10, rectangleSize/10, rl.White)
	} else {
		rl.DrawRectangle((x*rectangleSize + rectangleSize*3/10), y*rectangleSize+(rectangleSize/2), rectangleSize*2/5, rectangleSize/2, rl.Black)
		//eyes
		rl.DrawRectangle((x*rectangleSize + rectangleSize*3/10), y*rectangleSize+rectangleSize*9/10, rectangleSize/10, rectangleSize/10, rl.White)
		rl.DrawRectangle(x*rectangleSize+rectangleSize*6/10, y*rectangleSize+rectangleSize*9/10, rectangleSize/10, rectangleSize/10, rl.White)
	}
}

func tubeHeadHorizontal(x int32, y int32) {
	if lastDirection == game.Right {
		rl.DrawRectangle(x*rectangleSize, y*rectangleSize+rectangleSize*3/10, rectangleSize/2, rectangleSize*2/5, rl.Black)
		//eyes
		rl.DrawRectangle(x*rectangleSize, y*rectangleSize+rectangleSize*3/10, rectangleSize/10, rectangleSize/10, rl.White)
		rl.DrawRectangle(x*rectangleSize, y*rectangleSize+rectangleSize*6/10, rectangleSize/10, rectangleSize/10, rl.White)
	} else {
		rl.DrawRectangle(x*rectangleSize+(rectangleSize/2), y*rectangleSize+rectangleSize*3/10, rectangleSize/2, rectangleSize*2/5, rl.Black)
		//eyes
		rl.DrawRectangle(x*rectangleSize+rectangleSize*9/10, y*rectangleSize+rectangleSize*3/10, rectangleSize/10, rectangleSize/10, rl.White)
		rl.DrawRectangle(x*rectangleSize+rectangleSize*9/10, y*rectangleSize+rectangleSize*6/10, rectangleSize/10, rectangleSize/10, rl.White)
	}
}

func tubeVertical(x int32, y int32) {
	rl.DrawRectangle((x*rectangleSize + int32(int(rectangleSize)*3/10)), y*rectangleSize, int32(int(rectangleSize)*2/5), rectangleSize, rl.Black)
}

func tubeHorizontal(x int32, y int32) {
	rl.DrawRectangle((x * rectangleSize), y*rectangleSize+int32(int(rectangleSize)*3/10), rectangleSize, int32(int(rectangleSize)*2/5), rl.Black)
}

func cornerLeftTop(x int32, y int32) {
	// leftCenter
	rl.DrawRectangle(x*rectangleSize, y*rectangleSize+int32(int(rectangleSize)*3/10), int32(int(rectangleSize)*7/10), int32(int(rectangleSize)*2/5), rl.Black)
	// topCenter
	rl.DrawRectangle(x*rectangleSize+int32(int(rectangleSize)*3/10), y*rectangleSize, int32(int(rectangleSize)*2/5), int32(int(rectangleSize)*7/10), rl.Black)
}

func cornerLeftDown(x int32, y int32) {
	// leftCenter
	rl.DrawRectangle(x*rectangleSize, y*rectangleSize+int32(int(rectangleSize)*3/10), int32(int(rectangleSize)*7/10), int32(int(rectangleSize)*2/5), rl.Black)
	// downCenter
	rl.DrawRectangle(x*rectangleSize+int32(int(rectangleSize)*3/10), y*rectangleSize+int32(int(rectangleSize*3/10)), int32(int(rectangleSize)*2/5), int32(int(rectangleSize)*7/10), rl.Black)
}

func cornerRightTop(x int32, y int32) {
	// topCenter
	rl.DrawRectangle(x*rectangleSize+int32(int(rectangleSize)*3/10), y*rectangleSize, int32(int(rectangleSize)*2/5), int32(int(rectangleSize)*7/10), rl.Black)
	// rightCenter
	rl.DrawRectangle(x*rectangleSize+int32(int(rectangleSize*3/10)), y*rectangleSize+int32(int(rectangleSize)*3/10), int32(int(rectangleSize)*7/10), int32(int(rectangleSize)*2/5), rl.Black)
}

func cornerRightDown(x int32, y int32) {
	// rightCenter
	rl.DrawRectangle(x*rectangleSize+int32(int(rectangleSize*3/10)), y*rectangleSize+int32(int(rectangleSize)*3/10), int32(int(rectangleSize)*7/10), int32(int(rectangleSize)*2/5), rl.Black)
	// downCenter
	rl.DrawRectangle(x*rectangleSize+int32(int(rectangleSize)*3/10), y*rectangleSize+int32(int(rectangleSize*3/10)), int32(int(rectangleSize)*2/5), int32(int(rectangleSize)*7/10), rl.Black)
}
