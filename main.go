package main

import "fmt"

func main() {
	game := NewGame("game.json")

	fmt.Println(game.solve(game.Bottles))
}
