package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/christiangda/Golang-AdventOfCode2022/day_05/puzzle_01/ship"
)

var (
	input string
	debug bool
)

func ExtractCrates(line string, row int) []ship.Crate {
	column := 1
	crateSize := 3
	crates := make([]ship.Crate, 0)

	// each crate has 3 characters [\d]
	for len(line) >= crateSize {
		var subset string

		if len(line) > crateSize {
			subset = line[0 : crateSize+1]
		} else {
			subset = line[:]
		}

		cr := regexp.MustCompile(`(\[[A-Z]\])`)
		match := cr.Match([]byte(subset))

		if debug {
			fmt.Printf("column: %v, row: %v, size: %v, content: %s, is crate?: %v\n", column, row, len(line), subset, match)
		}

		if match {
			n := ship.Crate{
				Content: strings.Trim(subset, "[] "),
				Column:  column,
				Row:     row,
			}
			crates = append(crates, n)
		}

		if len(line) > crateSize { // remove the crate and the white space
			line = line[crateSize+1:]
		} else { // remove the last crate
			line = line[crateSize:]
		}

		column += 1
	}

	return crates
}

func ExtractInstructions(line string) ship.Instruction {
	var instruction ship.Instruction

	testInstruction := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	iMatches := testInstruction.FindStringSubmatch(line)

	if debug {
		fmt.Printf("instructions: %s\n", iMatches)
	}

	if len(iMatches) == 4 {
		m, err := strconv.Atoi(iMatches[1])
		if err != nil {
			fmt.Printf("Error converting: %v, to string", iMatches[1])
		}
		f, err := strconv.Atoi(iMatches[2])
		if err != nil {
			fmt.Printf("Error converting: %v, to string", iMatches[2])
		}
		t, err := strconv.Atoi(iMatches[3])
		if err != nil {
			fmt.Printf("Error converting: %v, to string", iMatches[3])
		}

		i := ship.Instruction{
			Text:  iMatches[0],
			Moves: m,
			From:  f,
			To:    t,
		}

		instruction = i
	}

	return instruction
}

func Parse(line string, row int) ([]ship.Crate, []ship.Instruction) {
	if debug {
		fmt.Printf("processing line: '%s'\n", line)
	}

	var creates []ship.Crate
	var instructions []ship.Instruction

	// crates
	// https://regex101.com/r/z2bHOX/1
	// `(\[[A-Z]\])`
	testCrates := regexp.MustCompile(`(\[[A-Z]\])`)
	hasCrates := testCrates.Match([]byte(line))
	if hasCrates {
		c := ExtractCrates(line, row)
		creates = append(creates, c...)
	}

	// instructions
	// https://regex101.com/r/zQvkff/1
	//`move (\d+) from (\d+) to (\d+)`
	instruction := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	hasInstructions := instruction.Match([]byte(line))
	if hasInstructions {
		i := ExtractInstructions(line)
		instructions = append(instructions, i)
	}

	return creates, instructions
}

func init() {
	flag.StringVar(&input, "input", "input.txt", "input: a file path with the content of your https://adventofcode.com/2022/day/1/input")
	flag.BoolVar(&debug, "debug", false, "debug (default: false)")
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

	row := 1
	var crates []ship.Crate
	var instructions []ship.Instruction
	rc := bufio.NewScanner(file)
	for rc.Scan() {
		if rc.Text() == "" {
			continue
		}
		c, i := Parse(rc.Text(), row)
		crates = append(crates, c...)
		instructions = append(instructions, i...)

		row += 1
	}

	s := ship.New()
	s.AddCrates(crates)
	s.AddInstructions(instructions)

	ship.Crane(s)

	var end []string
	for _, s := range s.Stacks {
		end = append(end, s.Pop().Content)
	}

	fmt.Printf("Number of crates: %v\n", len(crates))
	fmt.Printf("Number of instructions: %v\n", len(instructions))
	fmt.Printf("Number of stacks: %v\n", len(s.Stacks))
	fmt.Printf("Which crate will end up on top of each stack: %v\n", strings.Join(end, ""))
}
