package game

import (
	"fmt"
	"time"

	"math/rand"
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
type Direction position
type Game struct {
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

	NumberInputsToFruit = 0
	pathLength          = 0
	optimalPath         = 0
)

var DirectionMap = map[Direction]string{
	Up:    "Up",
	Down:  "Down",
	Right: "Right",
	Left:  "Left",
}

func (g *Game) InitGame(columns int, rows int, debug bool) *Game {
	//Initialize Game
	g.Debug = debug
	g.CurrentDirection = Right
	g.Columns = columns
	g.Rows = rows
	//Initialize grid with Snake & Food
	g.Grid = make([][]cell, g.Rows)
	for i := range g.Grid {
		g.Grid[i] = make([]cell, g.Columns)
	}
	initialPosition := position{X: g.Rows / 3, Y: g.Columns / 2}
	g.Snake = append(g.Snake, initialPosition)
	//update Grid
	g.Grid[g.Snake[0].Y][g.Snake[0].X] = SNAKE
	g.Grid[g.Columns/2][2*g.Rows/3] = FOOD

	//Initialize Metrics
	g.Metrics.SessionID = "asdf"
	g.Metrics.PlayerID = "Hans"
	g.Metrics.StartTime = time.Now()
	g.Metrics.TimeToLength = []LengthTime{{
		Length:    1,
		TimeSince: 0,
	}}
	g.Metrics.DirectionChanges = make([]DirectionChange, 0)
	//setting initial Optimal Path
	optimalPath = calcOPtimalPath(g.Snake[0].X, g.Columns/2, g.Snake[0].Y, 2*g.Rows/3)
	fmt.Println(g.Snake[0].X, g.Columns/3, g.Snake[0].Y, g.Rows/2, optimalPath)

	//Debug
	if g.Debug {
		g.test()
	}

	return g
}

func (g *Game) placeFood() (int, int) {
	var x, y int
	for {
		x = rand.Intn(g.Columns)
		y = rand.Intn(g.Rows)
		if g.Grid[y][x] == EMPTY {
			g.Grid[y][x] = FOOD
			break
		}
	}
	//add element to timeToLength
	g.timeToLength()
	//add element to InputsToFruit
	g.inputsToFruit()
	return x, y
}

func (g *Game) MoveSnake() {
	//new position
	newHead := position{
		X: (g.Snake[0].X + g.CurrentDirection.X),
		Y: (g.Snake[0].Y + g.CurrentDirection.Y),
	}

	//check collision
	if newHead.X >= g.Columns || newHead.Y >= g.Rows {
		//handle Game Over
		g.setGameOver("border")
		g.CurrentDirection = Stop
		return
	} else if newHead.X < 0 || newHead.Y < 0 { //check boundary collision
		//handle game over
		g.setGameOver("border")
		g.CurrentDirection = Stop
		return
	} else if g.Grid[newHead.Y][newHead.X] == SNAKE && g.Snake[len(g.Snake)-1] != newHead {
		//handle game over
		g.setGameOver("tail")
		g.CurrentDirection = Stop
		return
	}

	//increase pathLength
	pathLength++

	//check food eaten
	if g.Grid[newHead.Y][newHead.X] == FOOD {

		x, y := g.placeFood()
		g.Grid[newHead.Y][newHead.X] = SNAKE

		//add head
		newSnake := append([]position{newHead}, g.Snake...)
		g.Snake = newSnake

		//add element to pathFitness
		g.pathFitness(x, y)
	} else {
		//remove tail
		tail := g.Snake[len(g.Snake)-1]
		g.Snake = g.Snake[:len(g.Snake)-1]
		//add newHead
		newSnake := append([]position{newHead}, g.Snake...)
		g.Snake = newSnake

		//update grid
		g.Grid[tail.Y][tail.X] = EMPTY
	}

	//add newHead to grid
	g.Grid[newHead.Y][newHead.X] = SNAKE

	if g.Debug {
		g.test()
	}
}

func (g *Game) test() {
	/*g.printDirection()
	fmt.Println(g.Snake)
	g.printGrid()*/

	fmt.Println(g.Metrics)
}

func (g *Game) printDirection() {
	switch g.CurrentDirection {
	case Up:
		print("UP")
	case Right:
		print("RIGHT")
	case Down:
		print("DOWN")
	case Left:
		print("LEFT")
	}
	print(g.CurrentDirection.X)
	println(g.CurrentDirection.Y)
}

func (g *Game) printGrid() {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			print(g.Grid[i][j])
		}
		print("\n")
	}
	print("\n")
}
