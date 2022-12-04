package day3

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

// TODO: merge findDiff and findDiffs
// assumes there's only ever one character that is different
func findDiff(s1 string, s2 string) rune {
	for _, c := range s1 {
		if strings.ContainsRune(s2, c) {
			return c

		}
	}

	fmt.Printf("cannont find diff character in %s and %s", s1, s2)
	return 0
}

func findDiffs(s1 string, s2 string) string {
	result := ""
	for _, c := range s1 {
		if strings.ContainsRune(s2, c) {
			result += string(c)

		}
	}

	return result
}

func priority(r rune) int {
	// normally we'd be able to just subtract the rune by a number and get the score,
	// but uppercase has a lower int value than its priority should be.
	if unicode.IsUpper(r) {
		return int(r - 38)
	} else {
		return int(r - 96)
	}
}

func SumCompartiments(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)
		left := line[0 : length/2]
		right := line[(length / 2):length]
		diff := findDiff(left, right)
		sum += priority(diff)
	}

	return sum
}

func findBadge(buffer [3]string) rune {
	first := findDiffs(buffer[0], buffer[1])
	second := findDiffs(first, buffer[2])
	// note: result may contain multiple instances of the matching character. This is fine.

	return rune(second[0])
}

// TODO: clean up

func SumBadges(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	buffer := [3]string{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		buffer[i] = line
		i++
		if i == 3 {
			badge := findBadge(buffer)
			sum += priority(badge)
			i = 0
		}
	}

	return sum
}
