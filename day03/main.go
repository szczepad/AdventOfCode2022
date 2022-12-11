package main

import (
	"fmt"

	"github.com/szczepad/AdventOfCode2022/common/fileloader"
)

func main() {
	inputString := fileloader.LoadFile("input.txt")
	rucksacks := CreateRucksacksFromStringArray(inputString)
	elfgroups := CreateElfGroupsFromRucksacks(rucksacks)

	//Part1
	result1 := GetRucksackPriorities(rucksacks)
	fmt.Printf("The priorities of the item types in part 1 are : %d", result1)

	//Part2
	result2 := GetElfBadgePriorities(elfgroups)
	fmt.Printf("The priorities of the badges in part 2 are : %d", result2)
}
