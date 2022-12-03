package main

import (
	"fmt"
	"github.com/fwielstra/AoC2022/day1"
	"github.com/fwielstra/AoC2022/day2"
	"github.com/fwielstra/AoC2022/day3"
	"os"
	"strconv"
	"time"
)

func main() {
	day := os.Args[1]
	args := os.Args[2:]

	switch day {
	case "day1":
		runDay1(args)
	case "day2":
		runDay2(args)
	case "day3":
		runDay3(args)
	default:
		fmt.Printf("unrecognized day %s\n", day)
	}
}

func runDay1(args []string) {
	if len(args) != 2 {
		fmt.Printf("invalid args")
		os.Exit(1)
	}

	filename := args[0]
	top, _ := strconv.Atoi(args[1])

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		fmt.Printf("error opening input file %s: %v", filename, err)
		os.Exit(1)
	}

	result := day1.FindHighestCalories(file, top)
	fmt.Printf("Day 1 result top %d for file %s: %d\n", top, filename, result)
}

func runDay2(args []string) {
	if len(args) != 1 {
		fmt.Printf("invalid args")
		os.Exit(1)
	}

	filename := args[0]
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		fmt.Printf("error opening input file %s: %v", filename, err)
		os.Exit(1)
	}

	start := time.Now()
	scores := day2.GetScores(file)
	sum := day2.Sum(scores)
	elapsed := time.Since(start)
	fmt.Printf("Result for day 2: %d in %s\n", sum, elapsed)

	file.Close()

	// simple approach
	file, _ = os.Open(filename)
	start = time.Now()
	dumb := day2.Dumb(file)
	elapsed = time.Since(start)
	fmt.Printf("Result for day 2 dumb: %d in %s\n", dumb, elapsed)

	// part 2

	file, _ = os.Open(filename)
	start = time.Now()
	scores = day2.GetScoresPartTwo(file)
	sum = day2.Sum(scores)
	elapsed = time.Since(start)
	fmt.Printf("Result for day 2 part 2: %d in %s\n", sum, elapsed)

	file.Close()
}

func runDay3(args []string) {
	filename := args[0]
	file, _ := os.Open(filename)
	defer file.Close()

	start := time.Now()
	sum := day3.SumCompartiments(file)
	elapsed := time.Since(start)
	fmt.Printf("Result for day 1 part 1: %d in %s\n", sum, elapsed)
}
