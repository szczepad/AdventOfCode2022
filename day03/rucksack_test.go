package main

import "testing"

func TestCreateRucksacksFromStringArray(t *testing.T) {
	input := []string{"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}

	want := []Rucksack{{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"}, {"PmmdzqPrV", "vPwwTWBwg"}}

	rucksacks := CreateRucksacksFromStringArray(input)

	if len(rucksacks) != len(want) {
		t.Errorf("Not the same amount of Rucksacks. Got: %d, Want: %d", len(rucksacks), len(want))
	}
	for i, rucksack := range rucksacks {
		if rucksack.compartmentOne != want[i].compartmentOne || rucksack.compartmentTwo != want[i].compartmentTwo {
			t.Errorf("Rucksacks do not match. Got: %v-%v, Want: %v-%v", rucksack.compartmentOne, rucksack.compartmentTwo, want[i].compartmentOne, want[i].compartmentTwo)
		}

	}
}

func TestCreateElfGroupsFromRucksacks(t *testing.T) {
	input := CreateRucksacksFromStringArray([]string{"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	})

	want := []ElfGroup{
		{[3]Rucksack{
			{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
			{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
			{"PmmdzqPrV", "vPwwTWBwg"},
		}}, {[3]Rucksack{
			{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"},
			{"ttgJtRGJ", "QctTZtZT"},
			{"CrZsJsPPZsGz", "wwsLwLmpwMDw"},
		}}}

	got := CreateElfGroupsFromRucksacks(input)

	if len(got) != len(want) {
		t.Errorf("Elfgroup lengths do not match. Got: %d, Want: %d", len(got), len(want))
	}

	for i, elfgroup := range got {
		for j, rucksack := range elfgroup.rucksacks {
			if rucksack.compartmentOne != want[i].rucksacks[j].compartmentOne || rucksack.compartmentTwo != want[i].rucksacks[j].compartmentTwo {
				t.Errorf("Compartments do not match. Got: %v%v, Want: %v%v", rucksack.compartmentOne, rucksack.compartmentTwo, want[i].rucksacks[j].compartmentOne, want[i].rucksacks[j].compartmentTwo)
			}

		}

	}
}

func TestGetDuplicatedItem(t *testing.T) {
	type test struct {
		input Rucksack
		want  rune
	}

	tests := []test{{Rucksack{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}, 'p'},
		{Rucksack{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"}, 'L'},
		{Rucksack{"PmmdzqPrV", "vPwwTWBwg"}, 'P'}}

	for _, tc := range tests {
		got := tc.input.GetDuplicatedItem()

		if got != tc.want {
			t.Errorf("Could not get duplicated Item correctly. Got: %c, Want: %c", got, tc.want)
		}
	}
}

func TestGetItemPriority(t *testing.T) {
	type test struct {
		input rune
		want  int
	}

	tests := []test{{'a', 1}, {'A', 27}, {'p', 16}, {'L', 38}, {'P', 42}, {'v', 22}, {'t', 20}, {'s', 19}}

	for _, tc := range tests {
		got := GetItemPriority(tc.input)

		if got != tc.want {
			t.Errorf("Did not get right item priority. Got: %d, Want: %d", got, tc.want)
		}
	}
}

func TestGetRucksackPriorities(t *testing.T) {
	input := CreateRucksacksFromStringArray([]string{"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	})

	want := 157

	got := GetRucksackPriorities(input)

	if got != want {
		t.Errorf("Did not get the right Rucksack priority. Got: %d, Want: %d", got, want)
	}
}

func TestGetBadge(t *testing.T) {
	type test struct {
		input ElfGroup
		want  rune
	}

	tests := []test{
		{ElfGroup{[3]Rucksack{
			{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
			{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
			{"PmmdzqPrV", "vPwwTWBwg"},
		}}, 'r'},
		{ElfGroup{[3]Rucksack{
			{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"},
			{"ttgJtRGJ", "QctTZtZT"},
			{"CrZsJsPPZsGz", "wwsLwLmpwMDw"},
		}}, 'Z'}}

	for _, tc := range tests {
		got := tc.input.GetBadge()

		if got != tc.want {
			t.Errorf("Did not get right Badge. Got: %c, Want: %c", got, tc.want)
		}
	}
}

func TestGetElfBadgePriorities(t *testing.T) {
	input := []ElfGroup{
		{[3]Rucksack{
			{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
			{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
			{"PmmdzqPrV", "vPwwTWBwg"},
		}}, {[3]Rucksack{
			{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"},
			{"ttgJtRGJ", "QctTZtZT"},
			{"CrZsJsPPZsGz", "wwsLwLmpwMDw"},
		}}}

	want := 70

	got := GetElfBadgePriorities(input)

	if got != want {
		t.Errorf("Did not get the right Badge priorities. Got: %d, Want: %d", got, want)
	}

}
