package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/christiangda/Golang-AdventOfCode2022/day_01/puzzle_01/elves"
	"github.com/christiangda/Golang-AdventOfCode2022/day_01/puzzle_01/manager"
)

var input string

func init() {
	flag.StringVar(&input, "input", "input.txt", "input: a file path with the content of your https://adventofcode.com/2022/day/1/input")

	flag.Parse()
}

func main() {
	if input == "" {
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open(input)
	if err != nil {
		fmt.Printf("Error: Looks like the input file is not valid %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	elves := elves.New()

	manager := manager.New(&elves, file)
	err = manager.Assign()
	if err != nil {
		fmt.Printf("Error: looks like the values in the file are not valid %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total Elves: %d\n", len(elves))
	fmt.Printf("Total Calories: %d\n", elves.TotalCalories())
	fmt.Printf("Elf with most Calories: %v, Calories: %v\n", elves.MostCalories().Name, elves.MostCalories().TotalCalories())
}
