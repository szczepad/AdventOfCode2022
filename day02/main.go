package main

import "fmt"

func main() {
	file := LoadFile("input.txt")
	matches := CreateMatchesFromInput(file)

	//Part 1
	result := ExecuteGame(matches)
	fmt.Printf("The result after following the strategy guide is : %d \n", result)

	//Part 2
	result2 := ExecuteGamePart2(matches)
	fmt.Printf("The result after following part 2 of the strategy guide is : %d \n", result2)
}
