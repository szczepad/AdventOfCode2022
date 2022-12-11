package main

type Rucksack struct {
	compartmentOne, compartmentTwo string
}

type ElfGroup struct {
	rucksacks [3]Rucksack
}

func CreateRucksacksFromStringArray(input []string) []Rucksack {
	var result []Rucksack

	for _, entry := range input {
		result = append(result, Rucksack{entry[0:(len(entry) / 2)], entry[len(entry)/2:]})

	}

	return result
}

func CreateElfGroupsFromRucksacks(input []Rucksack) []ElfGroup {
	var result []ElfGroup

	var tempGroup [3]Rucksack
	for i, rucksack := range input {
		tempGroup[i%3] = rucksack
		if (i+1)%3 == 0 {
			result = append(result, ElfGroup{tempGroup})
		}
	}

	return result
}

func (r *Rucksack) GetDuplicatedItem() rune {
	var result rune

	for _, itemOne := range r.compartmentOne {
		for _, itemTwo := range r.compartmentTwo {
			if itemOne == itemTwo {
				result = itemOne
			}
		}
	}

	return result
}

func (e *ElfGroup) GetBadge() rune {
	var result rune

	var rucksack1 = e.rucksacks[0].compartmentOne + e.rucksacks[0].compartmentTwo
	var rucksack2 = e.rucksacks[1].compartmentOne + e.rucksacks[1].compartmentTwo
	var rucksack3 = e.rucksacks[2].compartmentOne + e.rucksacks[2].compartmentTwo

	for _, item := range rucksack1 {
		for _, comparison1 := range rucksack2 {
			if item == comparison1 {
				for _, comparison2 := range rucksack3 {
					if item == comparison2 {
						result = item
					}
				}
			}
		}
	}

	return result
}

func GetItemPriority(input rune) int {
	var result int

	result = int(input - '0')
	if result >= 49 {
		result = result - 48
	} else {
		result += 10
	}

	return result
}

func GetRucksackPriorities(input []Rucksack) int {
	var result int

	for _, rucksack := range input {
		incorrectItem := rucksack.GetDuplicatedItem()
		result += GetItemPriority(incorrectItem)
	}

	return result
}

func GetElfBadgePriorities(input []ElfGroup) int {
	var result int

	for _, elfgroup := range input {
		badge := elfgroup.GetBadge()
		result += GetItemPriority(badge)
	}

	return result
}
