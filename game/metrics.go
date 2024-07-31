package game

import (
	"time"
)

type Metrics struct {
	SessionID        string            `json:"session_id"`
	PlayerID         string            `json:"player_id"`
	StartTime        time.Time         `json:"start_time"`
	EndTime          time.Time         `json:"end_time"`
	FinalLength      int               `json:"final_length"`
	TimeToLength     []LengthTime      `json:"time_to_length"`
	MeanTimeToFruit  time.Duration     `json:"mean_time_to_fruit"`
	DirectionChanges []DirectionChange `json:"direction_changes"`
	InputsToFruit    []InputsToFruit   `json:"inputs_to_fruit"`
	PathFitness      []PathFitness     `json:"path_fitness"`
	Heatmap          []Heatmap         `json:"heatmap"`
	GameOver         GameOver          `json:"game_over"`
}

type LengthTime struct {
	Length    int           `json:"length"`
	TimeSince time.Duration `json:"time_since"`
}

type DirectionChange struct {
	Direction string    `json:"direction"`
	Timestamp time.Time `json:"timestamp"`
}

type InputsToFruit struct {
	FruitNumber int `json:"fruit_number"`
	Inputs      int `json:"inputs"`
}

type PathFitness struct {
	FruitNumber int     `json:"fruit_number"`
	ActualPath  int     `json:"actual_path"`
	OptimalPath int     `json:"optimal_path"`
	PathRatio   float32 `json:"path_ratio"`
}

type Heatmap struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Visits int `json:"visits"`
}

type GameOver struct {
	Cause    string `json:"cause"`
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"position"`
}

func (g *Game) setGameOver(gameOver string) {
	g.Metrics.EndTime = time.Now()
	g.Metrics.FinalLength = len(g.Snake)
	g.Metrics.GameOver = GameOver{
		Cause: gameOver,
		Position: struct {
			X int "json:\"x\""
			Y int "json:\"y\""
		}(g.Snake[0]),
	}
	g.Metrics.MeanTimeToFruit = (g.Metrics.EndTime.Sub(g.Metrics.StartTime)) / time.Duration(g.Metrics.FinalLength)
	g.test()
}

func (g *Game) timeToLength() {
	var passedTime time.Duration
	for _, time := range g.Metrics.TimeToLength {
		passedTime += time.TimeSince
	}
	data := LengthTime{
		Length:    len(g.Snake),
		TimeSince: time.Since(g.Metrics.StartTime) - passedTime,
	}
	g.Metrics.TimeToLength = append(g.Metrics.TimeToLength, data)
}

func (g *Game) inputsToFruit() {
	data := InputsToFruit{
		FruitNumber: len(g.Snake),
		Inputs:      NumberInputsToFruit,
	}
	g.Metrics.InputsToFruit = append(g.Metrics.InputsToFruit, data)
	NumberInputsToFruit = 0
}

// TODO first ratio = inf+
// TODO ratio can be below zero {4 11 12 0.9166667}
func (g *Game) pathFitness(x int, y int) {
	data := PathFitness{
		FruitNumber: len(g.Snake),
		ActualPath:  pathLength,
		OptimalPath: optimalPath,
		PathRatio:   float32(pathLength) / float32(optimalPath),
	}
	g.Metrics.PathFitness = append(g.Metrics.PathFitness, data)
	//reset pathLength and set new Optimal path
	pathLength = 0
	optimalPath = calcOPtimalPath(g.Snake[0].X, x, g.Snake[0].Y, y)
}

func calcOPtimalPath(x1, x2, y1, y2 int) int {
	//apparently go has no abs function for int
	//and I consider this better than type conversion to float
	difX := x1 - x2
	difY := y1 - y2
	if difX < 0 {
		difX = -difX
	}
	if difY < 0 {
		difY = -difY
	}
	dif := difX + difY

	return dif
}
