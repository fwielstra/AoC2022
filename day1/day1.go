package day1

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

// returns the top N calorie carriers
func FindHighestCalories(input io.Reader, amount int) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	calories := make([]int, 0)
	var current = 0
	for scanner.Scan() {
		next := strings.TrimSpace(scanner.Text())
		if next == "" {
			calories = append(calories, current)
			current = 0
			continue
		}

		num, err := strconv.Atoi(next)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}

		current += num
	}

	sort.Ints(calories)

	// sum the top n result
	sum := 0
	for _, c := range calories[len(calories)-amount:] {
		sum += c
	}

	return sum
}
