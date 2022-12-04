package day3

import (
	"strings"
	"testing"
)

func TestDay3(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

	result := SumCompartiments(strings.NewReader(input))
	if result != 157 {
		t.Errorf("SumCompartiments(): expected 157, got %d", result)
	}
}

func TestDay3PartTwo(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

	result := SumBadges(strings.NewReader(input))
	if result != 70 {
		t.Errorf("SumBadges(): expected 70, got %d", result)
	}
}
