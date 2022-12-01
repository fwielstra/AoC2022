package day1

import (
	"strings"
	"testing"
)

func TestBlep(t *testing.T) {
	input := `1000
	2000
	3000

	4000

	5000
	6000

	7000
	8000
	9000

	10000
	`

	r := strings.NewReader(input)
	result := FindHighestCalories(r, 1)
	if result != 24000 {
		t.Errorf("FindHighestCalories(r, 1): expected 24000, got %d", result)
	}

	r = strings.NewReader(input)
	result2 := FindHighestCalories(r, 3)

	if result2 != 45000 {
		t.Errorf("FindHighestCalories(r, 3): expected 45000, got %d", result)
	}

}
