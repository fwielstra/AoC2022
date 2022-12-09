package day6

import (
	"fmt"
	"testing"
)

func TestFindStartMarker(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 7},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 5},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", want: 6},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 10},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 11},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf(`FindStartMarker("%s")`, tt.input), func(t *testing.T) {
			if got := FindStartMarker(tt.input); got != tt.want {
				t.Errorf("FindStartMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindStartMessage(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 19},
		{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 23},
		{input: "nppdvjthqldpwncqszvftbrmjlhg", want: 23},
		{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 29},
		{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 26},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf(`FindStartMessage("%s")`, tt.input), func(t *testing.T) {
			if got := FindStartMessage(tt.input); got != tt.want {
				t.Errorf("FindStartMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
