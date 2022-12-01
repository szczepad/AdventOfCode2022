package main

import "testing"

func TestCreateInventoryFromStringArray(t *testing.T) {
	input := []string{"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000"}

	want := []inventory{{[]int{1000, 2000, 3000}}, {[]int{4000}}, {[]int{5000, 6000}}, {[]int{7000, 8000, 9000}}, {[]int{10000}}}

	got := CreateInventoryFromStringArray(input)

	if len(got) != len(want) {
		t.Errorf("Arrays do not have equal length. Got: %d, Want: %d", len(got), len(want))
	}

	for i, inventory := range want {
		for j, value := range inventory.calories {
			if value != got[i].calories[j] {
				t.Errorf("Inventory values do not match. Got: %d, want: %d", got[i].calories[j], value)
			}

		}

	}
}

func TestCalculateTotalCalories(t *testing.T) {
	input := inventory{[]int{1000, 2000, 3000}}

	want := 6000

	got := input.CalculateTotalCalories()

	if got != want {
		t.Errorf("Could not correctly calculate total calories. Got: %d, Want: %d", got, want)
	}
}

func TestGetMaxCalories(t *testing.T) {
	input := []inventory{{[]int{1000, 2000, 3000}}, {[]int{4000}}, {[]int{5000, 6000}}, {[]int{7000, 8000, 9000}}, {[]int{10000}}}

	want := 24000

	got := GetMaxCalories(input)

	if got != want {
		t.Errorf("Max Calories were not calculated correctly. Got: %d, Want: %d", got, want)
	}
}

func TestGetSumTopThreeMaxCalories(t *testing.T) {
	input := []inventory{{[]int{1000, 2000, 3000}}, {[]int{4000}}, {[]int{5000, 6000}}, {[]int{7000, 8000, 9000}}, {[]int{10000}}}

	want := 45000

	got := GetSumTopThreeMaxCalories(input)

	if got != want {
		t.Errorf("Max Calories were not calculated correctly. Got: %d, Want: %d", got, want)
	}
}
