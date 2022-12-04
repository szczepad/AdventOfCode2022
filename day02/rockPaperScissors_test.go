package main

import "testing"

func TestCreateMatchesFromInput(t *testing.T) {
	input := []string{"A Y", "B X", "C Z"}

	want := []match{{"A", "Y"}, {"B", "X"}, {"C", "Z"}}

	got := CreateMatchesFromInput(input)

	if len(got) != len(want) {
		t.Errorf("Arrays do not have equal length. Got: %d, Want: %d", len(got), len(want))
	}
	for i, v := range want {
		if v != got[i] {
			t.Errorf("Array elements do not match. Index: %d, Got: %s, want: %s", i, got[i], v)
		}
	}

}

func TestExecuteMatch(t *testing.T) {
	type test struct {
		input match
		want  int
	}

	tests := []test{{match{"A", "Y"}, 8}, {match{"B", "X"}, 1}, {match{"C", "Z"}, 6}}

	for _, tc := range tests {
		got := tc.input.ExecuteMatch()

		if got != tc.want {
			t.Errorf("Wrong Match result. Got: %d, Want: %d", got, tc.want)
		}

	}
}

func TestExecuteMatchPart2(t *testing.T) {
	type test struct {
		input match
		want  int
	}

	tests := []test{{match{"A", "Y"}, 4}, {match{"B", "X"}, 1}, {match{"C", "Z"}, 7}}

	for _, tc := range tests {
		got := tc.input.ExecuteMatchPart2()

		if got != tc.want {
			t.Errorf("Wrong Match result. Got: %d, Want: %d", got, tc.want)
		}

	}

}

func TestExecuteGame(t *testing.T) {
	input := []match{{"A", "Y"}, {"B", "X"}, {"C", "Z"}}

	want := 15

	got := ExecuteGame(input)

	if got != want {
		t.Errorf("Wrong end result of game. Got: %d, Want: %d", got, want)
	}

}

func TestExecuteGamePart2(t *testing.T) {
	input := []match{{"A", "Y"}, {"B", "X"}, {"C", "Z"}}

	want := 12

	got := ExecuteGamePart2(input)

	if got != want {
		t.Errorf("Wrong end result of game. Got: %d, Want: %d", got, want)
	}
}
