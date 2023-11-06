package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	input string
	debug bool
)

func Reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i] // reverse the slice
	}
}

func GetBlockerOrEdgeIndex(treeHeight int, treeRow []int) int {
	if treeRow == nil {
		return 0
	}

	var index int

	for i, height := range treeRow {
		i += 1
		if height >= treeHeight {
			index = i // where we find a tree with the same or
			break
		} else {
			index = len(treeRow)
		}
	}

	return index
}

type Tree struct {
	x int
	y int
}

func (t *Tree) ScenicScore(woods [][]int) int {
	tree := woods[t.y][t.x]

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

	// because the slice should be represented as it
	// is in the matrix
	topTreeCopy := make([]int, len(topTree))
	copy(topTreeCopy, topTree)

	leftTreeCopy := make([]int, len(leftTree))
	copy(leftTreeCopy, leftTree)

	Reverse(topTreeCopy)
	Reverse(leftTreeCopy)

	top := GetBlockerOrEdgeIndex(tree, topTreeCopy)
	right := GetBlockerOrEdgeIndex(tree, rightTree)
	bottom := GetBlockerOrEdgeIndex(tree, bottomTree)
	left := GetBlockerOrEdgeIndex(tree, leftTreeCopy)

	return top * right * bottom * left
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

	var maxScenicScore int
	rows := len(woods.trees)
	cols := len(woods.trees[0])

	for y := 1; y <= rows-2; y++ {
		for x := 1; x <= cols-2; x++ {
			tree := Tree{x: x, y: y}
			scenicScore := tree.ScenicScore(woods.trees)

			if debug {
				fmt.Printf("tree: %v, sc: %v\n", woods.trees[y][x], scenicScore)
			}

			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	fmt.Println("-----------------------------------------------------------------")
	fmt.Printf("Number of rows: %v\n", rows)
	fmt.Printf("Number of cols: %v\n", cols)
	fmt.Printf("Number of trees: %v\n\n", cols*rows)

	fmt.Printf("What is the highest scenic score possible for any tree?: %v\n", maxScenicScore)
}
