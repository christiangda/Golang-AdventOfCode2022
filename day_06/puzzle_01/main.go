package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	input string
	debug bool
)

func AreDifferent(r []rune) bool {
	var equals []int

	// brute force check (short string)
	for i := range r {
		for j := range r {
			if i != j {
				if debug {
					fmt.Printf("r1,r2: %v==%v\n", string(r[i]), string(r[j]))
				}
				if r[i] == r[j] {
					if debug {
						fmt.Println("found match, adding 1")
					}
					equals = append(equals, 1)
				}
			}
		}
	}

	var total int
	for _, e := range equals {
		total += e
	}

	return total == 0
}

func DetectSignal(line string, bufferSize int) int {
	lineLength := len(line)
	start := 0
	end := bufferSize
	var result int

	for lineLength > bufferSize-1 {
		toCheck := line[start:end]
		if debug {
			fmt.Printf("to check: %v\n", toCheck)
		}

		if AreDifferent([]rune(toCheck)) {
			if debug {
				fmt.Printf("string match: %v, index: %v\n", toCheck, end)
			}
			result = end
			break
		}
		start++
		end++
		lineLength -= 1
	}

	return result
}

func init() {
	flag.StringVar(&input, "input", "input.txt", "input: a file path with the content of your https://adventofcode.com/2022/day/6/input")
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

	bufferSize := 4
	var results []int

	rc := bufio.NewScanner(file)
	for rc.Scan() {
		line := rc.Text()
		if debug {
			fmt.Printf("line: %v\n", line)
		}
		i := DetectSignal(line, bufferSize)
		results = append(results, i)
	}

	for i, r := range results {
		fmt.Printf("Numbers of lines processed: %v\n", len(results))
		fmt.Printf("Line number: %v\n", i+1)
		fmt.Printf("Characters that need to be processed before the first start-of-packet marker is detected: %v\n", r)
	}
}
