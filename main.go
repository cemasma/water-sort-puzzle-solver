package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	file := flag.String("file", "647.json", "specify the mapping file you want to solve\n--file 647.json")

	flag.Parse()

	startDate := time.Now()
	game := NewGame(*file)

	result := game.solve()
	endDate := time.Now()

	diff := endDate.Sub(startDate)
	fmt.Printf("Solution: %v\nduration: %s\n", result.Bottles, diff.String())

	for _, v := range result.Moves {
		fmt.Println(v)
	}
}
