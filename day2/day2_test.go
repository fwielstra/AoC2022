package day2

import (
	"strings"
	"testing"
)

func TestDay2(t *testing.T) {
	input := `A Y
B X
C Z
	`

	r := strings.NewReader(input)
	result := GetScores(r)

	if len(result) != 3 {
		t.Errorf("GetScores(): expected 3 results, got %d", len(result))
	}

	if result[0] != 8 {
		t.Errorf("GetScores(): expected result 0 to be 8, got %d", result[0])
	}
	if result[1] != 1 {
		t.Errorf("GetScores(): expected result 1 to be 1, got %d", result[1])
	}
	if result[2] != 6 {
		t.Errorf("GetScores(): expected result 2 to be 6, got %d", result[2])
	}

	sum := Sum(result)
	if sum != 15 {
		t.Errorf("Sum(): expected 15, got %d", sum)
	}
}

// TODO: table test.
func TestWinScore(t *testing.T) {
	tie := Round{
		me:       Rock,
		opponent: Rock,
	}
	if tie.WinScore() != 3 {
		t.Errorf("WinScore(): expected 3 for a tie, got %d", tie.WinScore())
	}

	win := Round{
		me:       Rock,
		opponent: Scissors,
	}

	if win.WinScore() != 6 {
		t.Errorf("WinScore(): expected 6 for a win, got %d", tie.WinScore())
	}

	loss := Round{
		me:       Paper,
		opponent: Scissors,
	}

	if loss.WinScore() != 0 {
		t.Errorf("WinScore(): expected 0 for a loss, got %d", loss.WinScore())
	}
}

func TestScore(t *testing.T) {
	tie := Round{
		me:       Rock,
		opponent: Rock,
	}
	if tie.Score() != 4 {
		t.Errorf("Score(): expected 4 for a tie with Rock, got %d", tie.Score())
	}

	win := Round{
		me:       Paper,
		opponent: Rock,
	}

	if win.Score() != 8 {
		t.Errorf("Score(): expected 8 for a win with Paper, got %d", tie.WinScore())
	}

	loss := Round{
		me:       Scissors,
		opponent: Rock,
	}

	if loss.Score() != 3 {
		t.Errorf("Score(): expected 3 for a loss with scissors, got %d", loss.Score())
	}
}

func TestDay2Simple(t *testing.T) {
	input := `A Y
B X
C Z
	`

	r := strings.NewReader(input)
	sum := Dumb(r)
	if sum != 15 {
		t.Errorf("Dumb(): expected 15, got %d", sum)
	}
}

func TestPart2(t *testing.T) {
	input := `A Y
B X
C Z
	`

	r := strings.NewReader(input)
	result := GetScoresPartTwo(r)

	if result[0] != 4 {
		t.Errorf("GetScoresPartTwo(): expected result 0 to be 4, got %d", result[0])
	}
	if result[1] != 1 {
		t.Errorf("GetScoresPartTwo(): expected result 1 to be 1, got %d", result[1])
	}
	if result[2] != 7 {
		t.Errorf("GetScoresPartTwo(): expected result 2 to be 7, got %d", result[2])
	}

	sum := Sum(result)
	if sum != 12 {
		t.Errorf("GetScoresPartTwo(): expected 15, got %d", sum)
	}
}
