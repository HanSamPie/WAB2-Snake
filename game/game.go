package game

import (
	"fmt"

	"golang.org/x/exp/rand"
)

const (
	EMPTY = iota
	SNAKE
	FOOD
)

type cell int
type position struct {
	X, Y int
}
type direction struct {
	X, Y int
}
type GameState struct {
	CurrentDirection direction
	Snake            []position
	Grid             [][]cell
	debug            bool
}

var (
	//TODO fix direction
	Right    = direction{X: 0, Y: 1}
	Left     = direction{X: 0, Y: -1}
	Down     = direction{X: 1, Y: 0}
	Up       = direction{X: -1, Y: 0}
	GameOver = direction{X: 0, Y: 0}

	Columns int
	Rows    int

	gameState *GameState
)

func InitGame(columns int, rows int, debug bool) *GameState {
	var state GameState
	gameState = &state
	state.debug = debug
	Columns = columns
	Rows = rows

	state.Grid = make([][]cell, Rows)
	for i := range state.Grid {
		state.Grid[i] = make([]cell, Columns)
	}
	initialPosition := position{X: Rows / 2, Y: Columns / 3}
	state.Snake = append(state.Snake, initialPosition)
	state.Grid[state.Snake[0].X][state.Snake[0].Y] = SNAKE
	state.Grid[Rows/2][2*Columns/3] = FOOD

	if debug {
		test()
	}

	return &state
}

func placeFood() {
	for {
		x := rand.Intn(Columns)
		y := rand.Intn(Rows)
		if gameState.Grid[x][y] == EMPTY {
			gameState.Grid[x][y] = FOOD
			break
		}
	}
}

func MoveSnake() {
	//new position
	newHead := position{
		X: (gameState.Snake[0].X + gameState.CurrentDirection.X),
		Y: (gameState.Snake[0].Y + gameState.CurrentDirection.Y),
	}

	//check collision
	if gameState.Grid[newHead.X][newHead.Y] == SNAKE {
		//handle Game Over
		gameState.CurrentDirection = GameOver
	} else if newHead.X < 0 || newHead.Y < 0 { //check boundary collision
		//handle game over
		gameState.CurrentDirection = GameOver
	} else if newHead.X >= Columns || newHead.Y >= Rows {
		//handle game over
		gameState.CurrentDirection = GameOver
	}

	//check food eaten
	if gameState.Grid[newHead.X][newHead.Y] == FOOD {
		placeFood()
		gameState.Grid[newHead.X][newHead.Y] = SNAKE

		//add head
		newSnake := append([]position{newHead}, gameState.Snake...)
		gameState.Snake = newSnake
	} else {
		//remove tail
		tail := gameState.Snake[len(gameState.Snake)-1]
		gameState.Snake = gameState.Snake[:len(gameState.Snake)-1]
		//add head
		newSnake := append([]position{newHead}, gameState.Snake...)
		gameState.Snake = newSnake

		//update grid
		gameState.Grid[tail.X][tail.Y] = EMPTY
	}
	for _, part := range gameState.Snake {
		gameState.Grid[part.X][part.Y] = SNAKE
	}
	if gameState.debug {
		test()
	}
}

func test() {
	printDirection()
	fmt.Println(gameState.Snake)
	printGrid()
}

func printDirection() {
	switch gameState.CurrentDirection {
	case Up:
		print("UP")
	case Right:
		print("RIGHT")
	case Down:
		print("DOWN")
	case Left:
		print("LEFT")
	}
	print(gameState.CurrentDirection.X)
	println(gameState.CurrentDirection.Y)
}

func printGrid() {
	for i := 0; i < Rows; i++ {
		for j := 0; j < Columns; j++ {
			print(gameState.Grid[i][j])
		}
		print("\n")
	}
	print("\n")
}
