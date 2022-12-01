package main

import "strconv"

type inventory struct{
	calories []int
}

func (i *inventory) CalculateTotalCalories() int {
	var result int 

	for _,value := range i.calories {
		result += value
	}

	return result
}

func CreateInventoryFromStringArray(input []string) []inventory{
	result := []inventory{}
	var partialResult []int 

	for _,value := range input {
		if value != "" {
			intValue, _ := strconv.Atoi(value)
			partialResult = append(partialResult,intValue)
		} else {
			result = append(result,inventory{partialResult})
			partialResult = []int{}
		}

	}

	result = append(result,inventory{partialResult})

	return result
}

func GetMaxCalories(input []inventory) int {
	var result int 

	for _, inventory := range input {
		if inventory.CalculateTotalCalories() > result {
			result = inventory.CalculateTotalCalories()
		}
	}

	return result
}

func GetSumTopThreeMaxCalories(input []inventory) int {
	var result int
	var topThreeMaxCalories [3]int

	for _, inventory := range input {
		if inventory.CalculateTotalCalories() > topThreeMaxCalories[0] {
			topThreeMaxCalories[2] = topThreeMaxCalories[1]
			topThreeMaxCalories[1] = topThreeMaxCalories[0]
			topThreeMaxCalories[0] = inventory.CalculateTotalCalories()
		} else if inventory.CalculateTotalCalories() > topThreeMaxCalories[1] {
			topThreeMaxCalories[2] = topThreeMaxCalories[1]
			topThreeMaxCalories[1] = inventory.CalculateTotalCalories()
		} else if inventory.CalculateTotalCalories() > topThreeMaxCalories[2] {
			topThreeMaxCalories[2] = inventory.CalculateTotalCalories()
		}
	}

	for _, value := range topThreeMaxCalories{
		result += value
	}

	return result
}
