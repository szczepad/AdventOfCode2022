package main

import "testing"

func TestLoadInputAsStrings(t *testing.T) {
	want := []string{
		"A Y",
		"B X",
		"C Z",
	}
	got := LoadFile("input_test.txt")

	if len(got) != len(want) {
		t.Errorf("Array lenght does not match. Got: %d, Want: %d", len(got), len(want))
	}
	for i, v := range want {
		if v != got[i] {
			t.Errorf("Array elements do not match. Index: %d, Got: %s, want: %s", i, got[i], v)
		}
	}

}
