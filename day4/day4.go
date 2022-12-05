package day4

import (
	"bufio"
	"github.com/fwielstra/AoC2022/math"
	"io"
	"strconv"
	"strings"
)

type assignment struct {
	from int
	to   int
}

func (a assignment) length() int {
	return math.Abs(a.from - a.to)
}

type assignmentSet struct {
	one   assignment
	other assignment
}

func (s assignmentSet) hasFullOverlap() bool {
	// check if the smallest set fits in the larger set.
	// Doesn't matter if they're the same length.
	biggest, smallest := s.one, s.other
	if biggest.length() < smallest.length() {
		biggest, smallest = smallest, biggest
	}
	return smallest.from >= biggest.from && smallest.to <= biggest.to
}

func (s assignmentSet) hasOverlap() bool {
	// I cheated: https://stackoverflow.com/a/41352616/204840
	res := (s.other.to-s.one.from)*(s.one.to-s.other.from) >= 0
	//fmt.Printf("one: %+v, other: %+v, overlap %t\n", s.one, s.other, res)
	return res

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

func CountFullContains(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		set := parseAssignmentSet(line)
		if set.hasFullOverlap() {
			sum++
		}
	}

	return sum
}

func CountOverlaps(input io.Reader) int {
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
