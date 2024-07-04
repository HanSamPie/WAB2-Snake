package main

import (
	"fmt"

	"golang.org/x/exp/rand"
)

const (
	EMPTY = iota
	SNAKE
	FOOD
)

type Cell int

var grid [][]Cell

type position struct {
	x, y int
}

var columns int
var rows int

var snake []position

type direction struct {
	x, y int
}

func initGame() {
	grid = make([][]Cell, rows)
	for i := range grid {
		grid[i] = make([]Cell, columns)
	}
	snake = []position{{x: columns / 3, y: rows / 2}}
	grid[2*columns/3][rows/2] = FOOD
	direction.x = 1
}

func placeFood() {
	for {
		x := rand.Intn(columns)
		y := rand.Intn(rows)
		if grid[x][y] == EMPTY {
			grid[x][y] = FOOD
			break
		}
	}
}

func moveSnake(direction, position) {
	//new position
	//newHead = position{
	//	x: (snake[0].x + direction.x + direction.y),
	//	y: (snake[0].x + direction.x + direction.y)
	//}

	//check collision
	//check boundary collision

	//check food eaten

	//update grid
}

func main() {
	fmt.Println("hello there")
}
