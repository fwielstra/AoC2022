package day5_test

import (
	"github.com/fwielstra/AoC2022/day5"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	result := day5.Process(strings.NewReader(input))
	if result != "CMZ" {
		t.Errorf("Process(): result should be CMZ, was %s", result)
	}
}

func TestProcess9001(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	result := day5.Process9001(strings.NewReader(input))
	if result != "MCD" {
		t.Errorf("Process9001(): result should be MCD, was %s", result)
	}
}
