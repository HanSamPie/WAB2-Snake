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
)

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
	g.Grid[g.Snake[0].Y][g.Snake[0].X] = SNAKE
	g.Grid[g.Rows/2][2*g.Columns/3] = FOOD

	//Initialize Metrics
	g.Metrics.SessionID = "asdf"
	g.Metrics.PlayerID = "Hans"
	g.Metrics.StartTime = time.Now()
	//TODO at GameOver remove empty values?
	g.Metrics.TimeToLength = []LengthTime{{
		Length:    1,
		TimeSince: 0,
	}}
	g.Metrics.DirectionChanges = make([]DirectionChange, 0)

	//Debug
	if g.Debug {
		g.test()
	}

	return g
}

func (g *Game) placeFood() {
	for {
		x := rand.Intn(g.Columns)
		y := rand.Intn(g.Rows)
		if g.Grid[y][x] == EMPTY {
			g.Grid[y][x] = FOOD
			break
		}
	}
	g.timeToLength()
}

func (g *Game) MoveSnake() {
	//new position
	newHead := position{
		X: (g.Snake[0].X + g.CurrentDirection.X),
		Y: (g.Snake[0].Y + g.CurrentDirection.Y),
	}

	//TODO Bug: if newHead hits last part game over but should not cause game over

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
	} else if g.Grid[newHead.Y][newHead.X] == SNAKE {
		//handle game over
		g.setGameOver("tail")
		g.CurrentDirection = Stop
		return
	}
	//check food eaten
	if g.Grid[newHead.Y][newHead.X] == FOOD {

		g.placeFood()
		g.Grid[newHead.Y][newHead.X] = SNAKE

		//add head
		newSnake := append([]position{newHead}, g.Snake...)
		g.Snake = newSnake
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
