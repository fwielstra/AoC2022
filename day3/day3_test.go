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
