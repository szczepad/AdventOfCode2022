package main

import "fmt"

func main() {
	inputString := LoadFile("input.txt")

	inventories := CreateInventoryFromStringArray(inputString)

	//Part 1 -> Get the Elf with the most Calories in its inventory
	result1 := GetMaxCalories(inventories)

	fmt.Printf("The elf with the most calories carries %d calories \n", result1)

	//Part 2 -> Get the sum of the top 3 Inventories
	result2 := GetSumTopThreeMaxCalories(inventories)

	fmt.Printf("The sum of the top3 inventories is %d calories \n", result2)

}
