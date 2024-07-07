package game

import "golang.org/x/exp/rand"

const (
	EMPTY = iota
	SNAKE
	FOOD
)

type Cell int

var Grid [][]Cell

type position struct {
	x, y int
}

type Direction struct {
	x, y int
}

var (
	Up    = Direction{x: 0, y: 1}
	Right = Direction{x: 1, y: 0}
	Down  = Direction{x: 0, y: -1}
	Left  = Direction{x: -1, y: 0}
)

var CurrentDirection Direction

var Columns int
var Rows int

var snake []position

func InitGame() {
	Grid = make([][]Cell, Rows)
	for i := range Grid {
		Grid[i] = make([]Cell, Columns)
	}
	snake = []position{{x: Rows / 2, y: Columns / 3}}
	Grid[snake[0].x][snake[0].y] = SNAKE
	Grid[Rows/2][2*Columns/3] = FOOD
}

func placeFood() {
	for {
		x := rand.Intn(Columns)
		y := rand.Intn(Rows)
		if Grid[x][y] == EMPTY {
			Grid[x][y] = FOOD
			break
		}
	}
}

func MoveSnake() {
	if len(snake) == 0 {
		InitGame()
	}

	//new position
	newHead := position{
		x: (snake[0].x + CurrentDirection.x + CurrentDirection.y),
		y: (snake[0].x + CurrentDirection.x + CurrentDirection.y),
	}

	//check collision
	if Grid[newHead.x][newHead.y] == SNAKE {
		//handle Game Over
	}
	//check boundary collision
	if newHead.x < 0 || newHead.y < 0 {
		//handle game over
	} else if newHead.x > Columns || newHead.y > Rows {
		//handle game over
	}
	//check food eaten
	if Grid[newHead.x][newHead.y] == FOOD {
		placeFood()
	} else {
		tail := snake[len(snake)-1]
		snake = snake[:len(snake)-1]
		Grid[tail.x][tail.y] = EMPTY
	}

	//update Grid
	Grid[newHead.x][newHead.y] = SNAKE

}

func Test() {
	//switch CurrentDirection {
	//case Up:
	//	println("UP")
	//case Right:
	//	println("RIGHT")
	//case Down:
	//	println("DOWN")
	//case Left:
	//	println("LEFT")
	//}
	for i := 0; i < Rows; i++ {
		for j := 0; j < Columns; j++ {
			print(Grid[i][j])
		}
		print("\n")
	}
}
