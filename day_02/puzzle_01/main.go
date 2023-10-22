package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/christiangda/Golang-AdventOfCode2022/day_02/puzzle_01/shapes"
)

var (
	input string
	debug bool
)

func init() {
	flag.StringVar(&input, "input", "input.txt", "input: a file path with the content of your https://adventofcode.com/2022/day/1/input")
	flag.BoolVar(&debug, "debug", false, "debug (default: false)")

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

	var results []int

	rc := bufio.NewScanner(file)
	for rc.Scan() {
		line := rc.Text()
		round := strings.Split(line, " ")
		if len(round) != 2 {
			fmt.Printf("Round do not have 2 players, has %v", len(round))
			os.Exit(1)
		}

		p1Token := round[0]
		p2Token := round[1]

		p1 := shapes.GetShape(p1Token)
		if p1 == nil {
			fmt.Print("Player 1 shape is null")
			os.Exit(1)
		}

		p2 := shapes.GetShape(p2Token)
		if p2 == nil {
			fmt.Print("Player 2 shape is null")
			os.Exit(1)
		}

		result := p1.PlayWith(p2)

		if debug {
			fmt.Printf("p1: {token: %v, name: %v, value: %v}, p2: {token: %v, name: %v, value: %v} -> round result: %v\n", p1Token, p1, p1.Value(), p2Token, p2, p2.Value(), result)
		}

		results = append(results, result)
	}

	fmt.Printf("Number of rounds: %v\n", len(results))

	var total int
	for _, v := range results {
		total += v
	}

	fmt.Printf("Total score according to my strategy: %v\n", total)
}
