package day4

import (
	"strings"
	"testing"
)

func TestDay4(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

	result := CountFullContains(strings.NewReader(input))
	if result != 2 {
		t.Errorf("CountFullContains(): expected 2, got %d", result)
	}
}
