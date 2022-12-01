package main

import "testing"

func TestLoadInputAsStrings(t *testing.T) {
	want := []string{"1000",
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
