package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fwielstra/AoC2022/day1"
)

func main() {
	day := os.Args[1]
	args := os.Args[2:]

	switch day {
	case "day1":
		runDay1(args)
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
	}

	result := day1.FindHighestCalories(file, top)
	fmt.Printf("Day 1 result top %d for file %s: %d\n", top, filename, result)
}
