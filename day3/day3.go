package day3

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

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
		diff := findDiffs(left, right)
		sum += priority(rune(diff[0]))
	}

	return sum
}

func findBadge(buffer [3]string) rune {
	first := findDiffs(buffer[0], buffer[1])
	second := findDiffs(first, buffer[2])
	// note: result may contain multiple instances of the matching character. This is fine.
	return rune(second[0])
}

func SumBadges(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	sum := 0
	buffer := [3]string{}
	i := 0

	for scanner.Scan() {
		line := scanner.Text()
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
