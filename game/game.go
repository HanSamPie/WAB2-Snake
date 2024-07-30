package game

import (
	"fmt"

	"math/rand"
)

const (
	EMPTY = iota
	SNAKE
	FOOD
)

type cell int
type position struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type Direction position
type GameState struct {
	CurrentDirection Direction
	Snake            []position
	Grid             [][]cell
	Columns          int
	Rows             int
	Debug            bool
	Metrics          Metrics
}

var (
	Up    = Direction{X: 0, Y: -1}
	Down  = Direction{X: 0, Y: 1}
	Right = Direction{X: 1, Y: 0}
	Left  = Direction{X: -1, Y: 0}
	Stop  = Direction{X: 0, Y: 0}

	gameState *GameState
)

func InitGame(columns int, rows int, debug bool) *GameState {
	var state GameState
	gameState = &state
	state.Debug = debug
	state.CurrentDirection = Right
	state.Columns = columns
	state.Rows = rows

	state.Grid = make([][]cell, state.Rows)
	for i := range state.Grid {
		state.Grid[i] = make([]cell, state.Columns)
	}
	initialPosition := position{X: state.Rows / 3, Y: state.Columns / 2}
	state.Snake = append(state.Snake, initialPosition)
	state.Grid[state.Snake[0].Y][state.Snake[0].X] = SNAKE
	state.Grid[state.Rows/2][2*state.Columns/3] = FOOD

	//Initialize Metrics

	if debug {
		test()
	}

	return &state
}

func placeFood() {
	for {
		x := rand.Intn(gameState.Columns)
		y := rand.Intn(gameState.Rows)
		if gameState.Grid[y][x] == EMPTY {
			gameState.Grid[y][x] = FOOD
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
	if newHead.X >= gameState.Columns || newHead.Y >= gameState.Rows {
		//handle Game Over
		gameState.CurrentDirection = Stop
		return
	} else if newHead.X < 0 || newHead.Y < 0 { //check boundary collision
		//handle game over
		gameState.CurrentDirection = Stop
		return
	} else if gameState.Grid[newHead.Y][newHead.X] == SNAKE {
		//handle game over
		gameState.CurrentDirection = Stop
		return
	}
	//check food eaten
	if gameState.Grid[newHead.Y][newHead.X] == FOOD {
		placeFood()
		gameState.Grid[newHead.Y][newHead.X] = SNAKE

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
		gameState.Grid[tail.Y][tail.X] = EMPTY
	}

	//add newHead to grid
	gameState.Grid[newHead.Y][newHead.X] = SNAKE

	if gameState.Debug {
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
	for i := 0; i < gameState.Rows; i++ {
		for j := 0; j < gameState.Columns; j++ {
			print(gameState.Grid[i][j])
		}
		print("\n")
	}
	print("\n")
}
