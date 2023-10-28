package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func areFullyContained(s1, s2 string) (bool, error) {
	sec1 := strings.Split(s1, "-")
	sec2 := strings.Split(s2, "-")

	if len(sec1) != 2 {
		return false, fmt.Errorf("section 1 is not completed, expected 2 , got: %v", len(sec1))
	}
	if len(sec2) != 2 {
		return false, fmt.Errorf("section 1 is not completed, expected 2 , got: %v", len(sec1))
	}

	t1, err := strconv.Atoi(sec1[0])
	if err != nil {
		return false, fmt.Errorf("error converting string to int: %v", err)
	}
	t2, err := strconv.Atoi(sec1[1])
	if err != nil {
		return false, fmt.Errorf("error converting string to int: %v", err)
	}

	u1, err := strconv.Atoi(sec2[0])
	if err != nil {
		return false, fmt.Errorf("error converting string to int: %v", err)
	}
	u2, err := strconv.Atoi(sec2[1])
	if err != nil {
		return false, fmt.Errorf("error converting string to int: %v", err)
	}

	if t1 <= u1 && t2 >= u2 {
		return true, nil
	}

	if u1 <= t1 && u2 >= t2 {
		return true, nil
	}

	return false, nil
}

var (
	input string
	debug bool
)

func init() {
	flag.StringVar(&input, "input", "input.txt", "input: a file path with the content of your https://adventofcode.com/2022/day/4/input")
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

	var got int
	var totalSections int
	rc := bufio.NewScanner(file)
	for rc.Scan() {
		sections := strings.Split(rc.Text(), ",")

		if len(sections) != 2 {
			fmt.Printf("bad entry from sections, expected 2, got %v", len(sections))
		}

		ok, err := areFullyContained(sections[0], sections[1])
		if err != nil {
			fmt.Printf("Error returned from function, err: %v", err)
		}

		if debug {
			fmt.Printf("sections 1: %v, 2: %v, are contained?: %v\n", sections[0], sections[1], ok)
		}

		if ok {
			got += 1
		}
		totalSections += len(sections)
	}

	fmt.Printf("Number of pairs of sections checked: %v\n", totalSections/2)
	fmt.Printf("Number of sections checked: %v\n", totalSections)
	fmt.Printf("Number of assignment pairs that one range fully contain the other: %v\n", got)
}
