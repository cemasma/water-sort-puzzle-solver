package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

const EMPTY = "EMPTY"

type bottlesArray [][]string

type Game struct {
	Bottles bottlesArray `json:"bottles"`
}

type Flow struct {
	Bottles bottlesArray
	Moves []string
}

func newFlow(bottles bottlesArray) Flow {
	moves := make([]string, 0)

	return Flow{
		Moves: moves,
		Bottles: bottles,
	}
}

func (game Game) solve() Flow {
	queueState := make([]Flow, 0)
	queueState = append(queueState, newFlow(game.Bottles))

	visited := make(map[string]bool)
	var solvedState Flow

	for len(queueState) != 0 {
		currentState := queueState[len(queueState) - 1]
		queueState = queueState[:len(queueState) - 1]

		if visited[hashBottles(currentState.Bottles)] {
			continue
		}

		visited[hashBottles(currentState.Bottles)] = true

		if isDone(currentState.Bottles) {
			solvedState = currentState
			break
		}

		for i := 0; i < len(currentState.Bottles); i++ {

			bottle1 := currentState.Bottles[i]
			if isBottleDone(bottle1) || bottle1[0] == EMPTY {
				continue
			}

			for j := 0; j < len(currentState.Bottles); j++ {
				if i == j {
					continue
				}

				bottle2 := currentState.Bottles[j]

				if bottle2[len(bottle2) - 1] != EMPTY {
					continue
				}

				if isMovePossible(bottle1, bottle2) {
					cs := getCopyOfSituation(currentState.Bottles)
					makeMove(&cs, i, j)
					flow := newFlow(getCopyOfSituation(cs))
					newMovesArr := make([]string, len(currentState.Moves))
					copy(newMovesArr, currentState.Moves)
					flow.Moves = append(newMovesArr, strconv.Itoa(i) + " -> " + strconv.Itoa(j))

					if !isAlreadySolved(queueState, cs) {
						queueState = append(queueState, flow)
					}
				}
			}
		}
	}

	return solvedState
}

func isBottleDone(bottle []string) bool {
	color := bottle[0]

	for i := 1; i < len(bottle); i++ {
		if color == EMPTY {
			return false
		}

		if bottle[i] != color {
			return false
		}
	}

	return true
}

func getCopyOfSituation(situation bottlesArray) bottlesArray {
	copySlice := make(bottlesArray, len(situation))
	for index, val := range situation {
		newArr := make([]string, len(val))
		copy(newArr, val)
		copySlice[index] = newArr
	}

	return copySlice
}

func isAlreadySolved(solutions []Flow, bottles bottlesArray) bool {
	for _, val := range solutions {
		s1 := hashBottles(val.Bottles)
		s2 := hashBottles(bottles)

		if s1 == s2 {
			return true
		}
	}

	return false
}

func hashBottles(bottles bottlesArray) string {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(bottles)
	return string(b.Bytes())
}

func makeMove(currentSituation *bottlesArray, i, j int) {
	color, index := getTopColorOfBottle((*currentSituation)[i])
	emptyIndex := getBottomEmptyPlace((*currentSituation)[j])

	(*currentSituation)[i][index] = EMPTY
	(*currentSituation)[j][emptyIndex] = color

	if isMovePossible((*currentSituation)[i], (*currentSituation)[j]) {
		makeMove(currentSituation, i, j)
	}
}

func isMovePossible(bottle1, bottle2 []string) bool {
	topColorBottle1, _ := getTopColorOfBottle(bottle1)
	topColorBottle2, _ := getTopColorOfBottle(bottle2)

	if bottle2[len(bottle2)-1] != EMPTY {
		return false
	}

	if topColorBottle1 == bottle1[0] && bottle2[0] == EMPTY {
		return false
	}

	if topColorBottle1 != topColorBottle2 && topColorBottle2 != EMPTY {
		return false
	}

	return true
}

func getBottomEmptyPlace(bottle []string) (index int) {
	for i := range bottle {
		if bottle[i] == EMPTY {
			index = i
			return
		}
	}

	return
}

func getTopColorOfBottle(bottle []string) (color string, index int) {
	color = EMPTY

	for i := len(bottle) - 1; i >= 0; i-- {
		if bottle[i] != EMPTY {
			color = bottle[i]
			index = i

			return
		}
	}

	return
}

func isDone(bottles bottlesArray) bool {
	isDone := true
	for _, bottle := range bottles {
		var previousColor string
		for _, color := range bottle {
			if previousColor == "" {
				previousColor = color
				continue
			}
			if color != previousColor {
				isDone = false
				break
			}

			previousColor = color
		}

		if isDone == false {
			break
		}
	}

	return isDone
}

func NewGame(filename string) (game Game) {
	byt, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byt, &game)

	if err != nil {
		panic(err)
	}

	return
}
