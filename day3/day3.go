package day3

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

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
