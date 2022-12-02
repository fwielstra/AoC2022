package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fwielstra/AoC2022/day1"
	"github.com/fwielstra/AoC2022/day2"
)

func main() {
	day := os.Args[1]
	args := os.Args[2:]

	switch day {
	case "day1":
		runDay1(args)
	case "day2":
		runDay2(args)
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

	scores := day2.GetScores(file)
	sum := day2.Sum(scores)
	fmt.Printf("Result for day 2: %d\n", sum)
}
