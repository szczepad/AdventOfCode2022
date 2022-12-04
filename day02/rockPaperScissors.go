package main

import "strings"

type match struct {
	opponent, player string
}

func CreateMatchesFromInput(input []string) []match {
	var result []match

	for _, current := range input {
		partialResultStr := strings.Split(current, " ")
		partialResult := match{partialResultStr[0], partialResultStr[1]}

		result = append(result, partialResult)
	}

	return result
}

func (m *match) ExecuteMatch() int {
	var result int

	result += m.calculatePointsFromSelection()
	result += m.calculatePointsFromResult()

	return result
}

func (m *match) ExecuteMatchPart2() int {
	var result int

	result += m.calculatePointsFromSelectionPart2()
	result += m.calculatePointsFromResultPart2()

	return result
}

func (m *match) calculatePointsFromSelection() int {
	if m.player == "X" {
		return 1
	} else if m.player == "Y" {
		return 2
	} else {
		return 3
	}
}

func (m *match) calculatePointsFromSelectionPart2() int {
	if m.player == "X" {
		if m.opponent == "A" {
			return 3
		} else if m.opponent == "B" {
			return 1
		} else {
			return 2
		}
	} else if m.player == "Y" {
		if m.opponent == "A" {
			return 1
		} else if m.opponent == "B" {
			return 2
		} else {
			return 3
		}
	} else {
		if m.opponent == "A" {
			return 2
		} else if m.opponent == "B" {
			return 3
		} else {
			return 1
		}
	}
}

func (m *match) calculatePointsFromResult() int {
	if m.opponent == "A" {
		if m.player == "X" {
			return 3
		} else if m.player == "Y" {
			return 6
		} else {
			return 0
		}
	} else if m.opponent == "B" {
		if m.player == "X" {
			return 0
		} else if m.player == "Y" {
			return 3
		} else {
			return 6
		}

	} else {
		if m.player == "X" {
			return 6
		} else if m.player == "Y" {
			return 0
		} else {
			return 3
		}
	}
}

func (m *match) calculatePointsFromResultPart2() int {
	if m.player == "X" {
		return 0
	} else if m.player == "Y" {
		return 3
	} else {
		return 6
	}
}

func ExecuteGame(matches []match) int {
	var result int

	for _, match := range matches {
		result += match.ExecuteMatch()
	}

	return result
}

func ExecuteGamePart2(matches []match) int {
	var result int

	for _, match := range matches {
		result += match.ExecuteMatchPart2()
	}

	return result
}
