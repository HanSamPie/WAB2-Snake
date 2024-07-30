package game

import (
	"time"
)

type Metrics struct {
	SessionID            string            `json:"session_id"`
	PlayerID             string            `json:"player_id"`
	StartTime            time.Time         `json:"start_time"`
	EndTime              time.Time         `json:"end_time"`
	FinalLength          int               `json:"final_length"`
	TimeToLength         []LengthTime      `json:"time_to_length"`
	SuccessChance        []SuccessChance   `json:"success_chance"`
	MeanTimeToFruit      []FruitTime       `json:"mean_time_to_fruit"`
	TotalMeanTimeToFruit time.Duration     `json:"total_mean_time_to_fruit"`
	DirectionChanges     []DirectionChange `json:"direction_changes"`
	InputsToFruit        []InputsToFruit   `json:"inputs_to_fruit"`
	PathFitness          []PathFitness     `json:"path_fitness"`
	Heatmap              []Heatmap         `json:"heatmap"`
	GameOver             GameOver          `json:"game_over"`
}

type LengthTime struct {
	Length int       `json:"length"`
	Time   time.Time `json:"time"`
}

type SuccessChance struct {
	LengthTarget int  `json:"length_target"`
	Reached      bool `json:"reached"`
}

type FruitTime struct {
	FruitNumber int       `json:"fruit_number"`
	Time        time.Time `json:"time"`
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
	FruitNumber int `json:"fruit_number"`
	ActualPath  int `json:"actual_path"`
	OptimalPath int `json:"optimal_path"`
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
	Time time.Time `json:"time"`
}
