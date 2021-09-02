package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

const EMPTY = "EMPTY"

type Game struct {
	Bottles [][]string `json:"bottles"`
}

func (game Game) solve(bottles [][]string) [][]string {
	queueState := make([][][]string, 0)
	queueState = append(queueState, bottles)

	visited := make([][][]string, 0)
	solvedState := make([][]string, 0)

	for len(queueState) != 0 {
		if len(queueState) % 1000 > 999 {
			fmt.Println(len(queueState))
		}

		currentState := queueState[len(queueState) - 1]
		queueState = queueState[:len(queueState) - 1]

		if isAlreadySolved(visited, currentState) {
			continue
		}

		visited = append(visited, getCopyOfSituation(currentState))

		if isDone(currentState) {
			solvedState = currentState
			break
		}

		for i := 0; i < len(currentState); i++ {
			for j := 0; j < len(currentState); j++ {
				if i == j {
					continue
				}

				bottle1 := currentState[i]
				bottle2 := currentState[j]

				if isBottleDone(bottle1) {
					continue
				}

				if isMovePossible(bottle1, bottle2) {
					cs := getCopyOfSituation(currentState)
					makeMove(&cs, i, j)

					if !isAlreadySolved(queueState, cs) {
						queueState = append(queueState, getCopyOfSituation(cs))
					}
				}
			}
		}
	}

	for _, v := range visited {
		fmt.Println(isDone(v))
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

func getCopyOfSituation(situation [][]string) [][]string {
	copySlice := make([][]string, len(situation))
	for index, val := range situation {
		newArr := make([]string, len(val))
		copy(newArr, val)
		copySlice[index] = newArr
	}

	return copySlice
}

func isAlreadySolved(solutions [][][]string, bottles [][]string) bool {
	for _, val := range solutions {
		s1 := getStringifiedBottles(val)
		s2 := getStringifiedBottles(bottles)

		if s1 == s2 {
			return true
		}
	}

	return false
}

func getStringifiedBottles(bottles [][]string) (str string) {
	for bottleIndex, bottle := range bottles {
		for colorIndex, color := range bottle {
			str += "|" + strconv.Itoa(bottleIndex) + ":" + strconv.Itoa(colorIndex) + ":" + color + "|"
		}
	}

	return str
}

func makeMove(currentSituation *[][]string, i, j int) {
	color, index := getTopColorOfBottle((*currentSituation)[i])
	emptyIndex := getBottomEmptyPlace((*currentSituation)[j])

	(*currentSituation)[i][index] = EMPTY
	(*currentSituation)[j][emptyIndex] = color
}

func isMovePossible(bottle1, bottle2 []string) bool {
	topColorBottle1, _ := getTopColorOfBottle(bottle1)
	topColorBottle2, _ := getTopColorOfBottle(bottle2)

	if bottle2[len(bottle2)-1] != EMPTY {
		return false
	}

	if topColorBottle1 != topColorBottle2 && topColorBottle2 != EMPTY {
		return false
	}

	return true
}

func getBottomEmptyPlace(bottle []string) (index int) {
	for i, _ := range bottle {
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

func IsDone(bottles [][]string) bool {
	return isDone(bottles)
}

func isDone(bottles [][]string) bool {
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
