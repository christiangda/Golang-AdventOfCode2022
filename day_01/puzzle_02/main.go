package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/christiangda/Golang-AdventOfCode2022/day_01/puzzle_02/elves"
	"github.com/christiangda/Golang-AdventOfCode2022/day_01/puzzle_02/manager"
)

var input string

func init() {
	flag.StringVar(&input, "input", "input.txt", "input: a file path with the content of your https://adventofcode.com/2022/day/1/input")
}

func main() {
	flag.Parse()

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

	e := elves.New()

	manager := manager.New(&e, file)
	err = manager.Assign()
	if err != nil {
		fmt.Printf("Error: looks like the values in the file are not valid %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total Elves: %d\n", len(e))
	fmt.Printf("Total Calories: %d\n", e.TotalCalories())

	topThree := e.TopThree()
	fmt.Println("\nTop Three Elf most Calories:")
	for _, elf := range topThree {
		fmt.Printf("-> Name: %v, Calories: %v\n", elf.Name, elf.TotalCalories())
	}
	fmt.Printf("\nTotal carrying calories (Top Three Elf): %v\n", topThree.TotalCalories())
}
