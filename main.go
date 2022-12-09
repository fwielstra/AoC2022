package main

import (
	"fmt"
	"github.com/fwielstra/AoC2022/day1"
	"github.com/fwielstra/AoC2022/day2"
	"github.com/fwielstra/AoC2022/day3"
	"github.com/fwielstra/AoC2022/day4"
	"github.com/fwielstra/AoC2022/day5"
	"github.com/fwielstra/AoC2022/day6"
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
	case "day4":
		runDay4(filename)
	case "day5":
		runDay5(filename)
	case "day6":
		runDay6(filename)
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
		fmt.Printf("Result for day 3 part 1: %d\n", sum)
	})

	withFile(filename, func(r io.Reader) {
		sum := day3.SumBadges(r)
		fmt.Printf("Result for day 3 part 2: %d\n", sum)
	})
}

func runDay4(filename string) {
	withFile(filename, func(r io.Reader) {
		pairs := day4.CountFullContains(r)
		fmt.Printf("Resullt for day 4 part 1: %d\n", pairs)
	})

	withFile(filename, func(r io.Reader) {
		overlaps := day4.CountOverlaps(r)
		fmt.Printf("Resullt for day 4 part 2: %d\n", overlaps)
	})
}

func runDay5(filename string) {
	withFile(filename, func(r io.Reader) {
		result := day5.Process(r)
		fmt.Printf("Result for day 5 part 1: %s\n", result)
	})

	withFile(filename, func(r io.Reader) {
		result := day5.Process9001(r)
		fmt.Printf("Result for day 5 part 2: %s\n", result)
	})
}

func runDay6(filename string) {
	withFile(filename, func(r io.Reader) {
		str, _ := io.ReadAll(r)
		result := day6.FindStartMarker(string(str))
		fmt.Printf("Result for day 6 part 1: %d\n", result)
	})

	withFile(filename, func(r io.Reader) {
		str, _ := io.ReadAll(r)
		result := day6.FindStartMessage(string(str))
		fmt.Printf("Result for day 6 part 2: %d\n", result)
	})
}
