package main

import (
	"fmt"
	"github.com/fwielstra/AoC2022/day1"
	"github.com/fwielstra/AoC2022/day2"
	"github.com/fwielstra/AoC2022/day3"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	day := os.Args[1]
	filename := os.Args[2]
	args := os.Args[3:]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s: %s\n", filename, err)
		os.Exit(1)
	}

	switch day {
	case "day1":
		runDay1(filename, args)
	case "day2":
		runDay2(filename)
	case "day3":
		runDay3(filename)
	default:
		fmt.Printf("unrecognized day %s\n", day)
	}

	file.Close()
}

func withFile(filename string, callback func(r io.Reader)) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s: %s\n", filename, err)
		os.Exit(1)
	}
	start := time.Now()
	callback(file)
	elapsed := time.Since(start)
	fmt.Printf("processing file %s took %s\n", filename, elapsed)

	err = file.Close()
	if err != nil {
		fmt.Printf("error closing file %s: %s\n", filename, err)
		os.Exit(1)
	}
}

func runDay1(filename string, args []string) {
	withFile(filename, func(r io.Reader) {
		top, _ := strconv.Atoi(args[0])
		result := day1.FindHighestCalories(r, top)
		fmt.Printf("Day 1 result top %d for file %s: %d\n", top, filename, result)
	})
}

func runDay2(filename string) {
	withFile(filename, func(r io.Reader) {
		scores := day2.GetScores(r)
		sum := day2.Sum(scores)
		fmt.Printf("Result for day 2: %d\n", sum)
	})

	withFile(filename, func(r io.Reader) {
		// simple approach
		dumb := day2.Dumb(r)
		fmt.Printf("Result for day 2 dumb: %d\n", dumb)
	})

	withFile(filename, func(r io.Reader) {
		// part 2
		scores := day2.GetScoresPartTwo(r)
		sum := day2.Sum(scores)
		fmt.Printf("Result for day 2 part 2: %d\n", sum)
	})
}

func runDay3(filename string) {
	withFile(filename, func(r io.Reader) {
		sum := day3.SumCompartiments(r)
		fmt.Printf("Result for day 1 part 1: %d\n", sum)
	})

	withFile(filename, func(r io.Reader) {
		sum := day3.SumBadges(r)
		fmt.Printf("Result for day 1 part 2: %d\n", sum)
	})
}
