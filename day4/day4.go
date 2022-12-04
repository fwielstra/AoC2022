package day4

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type assignment struct {
	from int
	to   int
}

func (a assignment) length() int {
	return abs(a.from - a.to)
}

type assignmentSet struct {
	one   assignment
	other assignment
}

func (s assignmentSet) hasOverlap() bool {
	// check if the smallest set fits in the larger set.
	// Doesn't matter if they're the same length.
	biggest, smallest := s.one, s.other
	if biggest.length() < smallest.length() {
		biggest, smallest = smallest, biggest
	}
	return smallest.from >= biggest.from && smallest.to <= biggest.to
}

func parseAssignment(s string) assignment {
	parts := strings.Split(s, "-")
	from, _ := strconv.Atoi(parts[0])
	to, _ := strconv.Atoi(parts[1])

	return assignment{
		from: from,
		to:   to,
	}
}

func parseAssignmentSet(s string) assignmentSet {
	pairs := strings.Split(s, ",")

	return assignmentSet{
		one:   parseAssignment(pairs[0]),
		other: parseAssignment(pairs[1]),
	}
}

// TODO: only check if the smaller range fits in the larger
func CountFullContains(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		set := parseAssignmentSet(line)
		if set.hasOverlap() {
			sum++
		}
	}

	return sum
}
