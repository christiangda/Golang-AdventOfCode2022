package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var (
	input string
	debug bool
)

type Tree struct {
	x int
	y int
}

func (t *Tree) Visible(woods [][]int) int {
	tree := woods[t.y][t.x]

	// topTree := woods[t.y-1][t.x : t.x+1]
	var topTree []int
	for y := 0; y < t.y; y++ {
		x := woods[y][t.x]
		topTree = append(topTree, x)
	}

	rightTree := woods[t.y][t.x+1:]

	var bottomTree []int
	for y := t.y + 1; y < len(woods); y++ {
		x := woods[y][t.x]
		bottomTree = append(bottomTree, x)
	}

	leftTree := woods[t.y][:t.x]

	top := tree > slices.Max(topTree)
	right := tree > slices.Max(rightTree)
	bottom := tree > slices.Max(bottomTree)
	left := tree > slices.Max(leftTree)

	if top || right || bottom || left {
		return 1
	}

	return 0
}

type Woods struct {
	trees [][]int
}

func NewWoods() *Woods {
	return &Woods{
		trees: make([][]int, 0),
	}
}

func (w *Woods) AddRowOfTrees(line string) error {
	trees := make([]int, 0)

	for _, r := range line {
		tree, err := strconv.Atoi(string(r))
		if err != nil {
			return fmt.Errorf("error converting string to int: %v", err)
		}
		trees = append(trees, tree)
	}

	w.trees = append(w.trees, trees)

	return nil
}

func (w *Woods) Visible() int {
	// edges
	left := len(w.trees)
	right := len(w.trees)
	top := len(w.trees[0]) - 2    // remove firs and last col
	bottom := len(w.trees[0]) - 2 // remove firs and last col

	return top + bottom + left + right
}

func init() {
	flag.StringVar(&input, "input", "input.txt", "input: a file path with the content of your https://adventofcode.com/2022/day/8/input")
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

	woods := NewWoods()
	rc := bufio.NewScanner(file)
	for rc.Scan() {
		woods.AddRowOfTrees(rc.Text())
	}

	var visibleTreesInside int
	visibleTreesOutside := woods.Visible()
	rows := len(woods.trees)
	cols := len(woods.trees[0])

	for y := 1; y <= rows-2; y++ {
		for x := 1; x <= cols-2; x++ {
			tree := Tree{x: x, y: y}
			visibleTreesInside += tree.Visible(woods.trees)
		}
	}
	visibleTrees := visibleTreesInside + visibleTreesOutside

	fmt.Println("-----------------------------------------------------------------")
	fmt.Printf("Number of rows: %v\n", rows)
	fmt.Printf("Number of cols: %v\n", cols)
	fmt.Printf("Number of trees: %v\n\n", cols*rows)
	fmt.Printf("Number of visible trees inside: %v\n", visibleTreesInside)
	fmt.Printf("Number of visible trees outside: %v\n\n", visibleTreesOutside)

	fmt.Printf("how many trees are visible from outside the grid?: %v\n", visibleTrees)
}
