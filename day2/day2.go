package day2

import (
	"bufio"
	"io"
	"strings"
)

// making data types because we can

type RPS int

const (
	Rock RPS = iota
	Paper
	Scissors
)

var Scores = map[RPS]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

type Round struct {
	me       RPS
	opponent RPS
}

func (r Round) WinScore() int {
	// trick; use the numerical value of rock/paper/scissors to determine win/loss/tie:
	// https://stackoverflow.com/questions/11377117/rock-paper-scissors-determine-win-loss-tie-using-math

	// tie
	if r.me == r.opponent {
		return 3
	}
	// win
	if (r.me-r.opponent+3)%3 == 1 {
		return 6
	}
	// loss
	return 0
}

func (r Round) Score() int {
	score := r.WinScore()
	return score + Scores[r.me]
}

func (r Round) getWinningMove() RPS {
	switch r.opponent {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	}

	return -1 // error
}

func (r Round) getLosingMove() RPS {
	switch r.opponent {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	}

	return -1 // error
}

func (r Round) ScoreWithOutcome() int {
	if r.me == 0 {
		// lose
		return Scores[r.getLosingMove()]
	}

	if r.me == 1 {
		// draw
		return 3 + Scores[r.opponent]
	}

	if r.me == 2 {
		// win
		return 6 + Scores[r.getWinningMove()]
	}

	return 0 // should not happen
}

var mapping = map[string]RPS{
	"A": Rock,
	"X": Rock,
	"B": Paper,
	"Y": Paper,
	"C": Scissors,
	"Z": Scissors,
}

func parseRound(round string) Round {
	parts := strings.Split(round, " ")
	return Round{
		opponent: mapping[parts[0]],
		me:       mapping[parts[1]],
	}
}

func parseRounds(input io.Reader) []Round {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	rounds := make([]Round, 0)
	for scanner.Scan() {
		next := strings.TrimSpace(scanner.Text())
		if next == "" {
			continue
		}
		rounds = append(rounds, parseRound(next))
	}
	return rounds
}

func GetScores(input io.Reader) []int {
	rounds := parseRounds(input)
	scores := make([]int, len(rounds))
	for i, r := range rounds {
		scores[i] = r.Score()
	}

	return scores
}

func GetScoresPartTwo(input io.Reader) []int {
	rounds := parseRounds(input)
	scores := make([]int, len(rounds))
	for i, r := range rounds {
		scores[i] = r.ScoreWithOutcome()
	}

	return scores
}

func Sum(numbers []int) int {
	result := 0
	for _, n := range numbers {
		result += n
	}
	return result
}

func Dumb(input io.Reader) int {
	scores := map[string]int{
		"A X": 4,
		"B X": 1,
		"C X": 7,
		"A Y": 8,
		"B Y": 5,
		"C Y": 2,
		"A Z": 3,
		"B Z": 9,
		"C Z": 6,
	}

	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		sum += scores[scanner.Text()]
	}
	return sum
}
