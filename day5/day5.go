package day5

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type crate rune

// TODO: extract to collections library, make generic
// https://github.com/golang/go/wiki/SliceTricks
type stack []rune

func push(s stack, r rune) stack {
	return append(s, r)
}

func pushn(s stack, r []rune) stack {
	return append(s, r...)
}

func prepend(s stack, r rune) stack {
	return append(stack{r}, s...)
}

func pop(s stack) (stack, rune) {
	return s[:len(s)-1], s[len(s)-1]
}

func popn(s stack, n int) (stack, []rune) {
	return s[:len(s)-n], s[len(s)-n:]
}

type move struct {
	from int
	to   int
	num  int
}

func (m move) String() string {
	return fmt.Sprintf("move %d from %d to %d", m.num, m.from, m.to)
}

var numbers = regexp.MustCompile("[0-9]+")

// TODO: this seems overly wordy
func parseMove(s string) move {
	res := numbers.FindAll([]byte(s), -1)
	num, _ := strconv.Atoi(string(res[0]))
	from, _ := strconv.Atoi(string(res[1]))
	to, _ := strconv.Atoi(string(res[2]))

	return move{
		from: from - 1,
		to:   to - 1,
		num:  num,
	}
}

type State struct {
	stacks []stack
}

// CrateMover 9000 carries one crate at a time
func (s *State) applyMove(m move) {
	for i := 0; i < m.num; i++ {
		var r rune
		s.stacks[m.from], r = pop(s.stacks[m.from])
		s.stacks[m.to] = push(s.stacks[m.to], r)
	}
}

// CrateMover 9001 moves a stack of crates at a time
func (s *State) applyMoveCrateMover9001(m move) {
	var r []rune
	s.stacks[m.from], r = popn(s.stacks[m.from], m.num)
	s.stacks[m.to] = pushn(s.stacks[m.to], r)
}

func isAlpha(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func parse(input io.Reader) (State, []move) {
	scanner := bufio.NewScanner(input)

	// shortcut; we know there's at most 9 stacks
	stacks := make([]stack, 11)
	var moves []move

	mode := 0 // 0 = crates, 1 = steps
	for scanner.Scan() {
		line := scanner.Text()
		// skip empty lines
		if line == "" {
			continue
		}

		// end of situation section
		if strings.Contains(line, " 1   2   3") {
			mode = 1
			continue
		}

		// parse crates
		if mode == 0 {
			for n, s := range line {
				if isAlpha(s) {
					// argh my brain
					idx := math.Floor((float64(n) - 1) / 4)
					stacks[int(idx)] = prepend(stacks[int(idx)], s)
				}
			}
		}

		if mode == 1 {
			moves = append(moves, parseMove(line))
		}
	}

	return State{stacks}, moves
}

func getTopCrates(s State) string {
	result := ""
	for _, s := range s.stacks {
		// skip empty stacks
		if len(s) == 0 {
			continue
		}
		result += string(s[len(s)-1])
	}
	return result
}

func Process(input io.Reader) string {
	state, moves := parse(input)
	for _, m := range moves {
		state.applyMove(m)
	}

	return getTopCrates(state)
}

func Process9001(input io.Reader) string {
	state, moves := parse(input)
	for _, m := range moves {
		state.applyMoveCrateMover9001(m)
	}

	return getTopCrates(state)
}
