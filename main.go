package main

import (
	"fmt"
	"time"
)

func main() {

	startDate := time.Now()
	game := NewGame("639.json")

	result := game.solve()
	endDate := time.Now()

	diff := endDate.Sub(startDate)
	fmt.Printf("Solution: %v\nduration: %s\n", result.Bottles, diff.String())

	for _, v := range result.Moves {
		fmt.Println(v)
	}
}
