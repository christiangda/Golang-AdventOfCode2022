package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func Priority(c string) int {
	// [a-z] -> [1-26]  # ASCII [97-122]
	// [A-Z] -> [27-52] # ASCII [65-90]

	b := []rune(c)
	asciiVal := int(b[0])

	if asciiVal >= 65 && asciiVal <= 90 {
		return asciiVal - 38
	} else if asciiVal >= 97 && asciiVal <= 122 {
		return asciiVal - 96
	} else {
		return 0
	}
}

func CommonItem(c1, c2 string) string {
	var common string

	for _, c := range c1 {
		if strings.Contains(c2, string(c)) {
			common = string(c)
			break
		}
	}

	return common
}

var (
	input string
	debug bool
)

func init() {
	flag.StringVar(&input, "input", "input.txt", "input: a file path with the content of your https://adventofcode.com/2022/day/3/input")
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

	var priorities []int
	rc := bufio.NewScanner(file)
	for rc.Scan() {
		rucksack := rc.Text()
		length := len(rucksack)

		compartment1 := rucksack[:length/2]
		compartment2 := rucksack[length/2:]

		ci := CommonItem(compartment1, compartment2)
		p := Priority(ci)

		if debug {
			fmt.Printf("common item: %v, priority: %v\n", ci, p)
		}

		priorities = append(priorities, p)
	}

	var sum int
	for _, v := range priorities {
		sum += v
	}

	fmt.Printf("Number of repeated Items: %v\n", len(priorities))
	fmt.Printf("Sum of the priorities: %v\n", sum)
}
