package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/christiangda/Golang-AdventOfCode2022/day_07/puzzle_01/stack"
	"github.com/christiangda/Golang-AdventOfCode2022/day_07/puzzle_01/tree"
)

var (
	input   string
	maxSize int
	debug   bool
)

type Element struct {
	Kind     int
	Content  string
	Argument string
}

func ParseLine(s string) *Element {
	// https://regex101.com/r/cV061E/1
	cmd := regexp.MustCompile(`^\$ (?P<command>[a-z]{2,})(?P<argument> .{0,}){0,}`)
	cmdMatch := cmd.FindStringSubmatch(s)
	if len(cmdMatch) > 0 {
		return &Element{
			Kind:     tree.Command,
			Content:  strings.TrimSpace(cmdMatch[cmd.SubexpIndex("command")]),
			Argument: strings.TrimSpace(cmdMatch[cmd.SubexpIndex("argument")]),
		}
	}

	// https://regex101.com/r/MV5BvH/1
	dir := regexp.MustCompile(`^(?P<kind>[a-z]{0,}) (?P<name>[a-z]{1,})`)
	dirMatch := dir.FindStringSubmatch(s)
	if len(dirMatch) > 0 {
		return &Element{
			Kind:     tree.Dir,
			Content:  strings.TrimSpace(dirMatch[dir.SubexpIndex("kind")]),
			Argument: strings.TrimSpace(dirMatch[dir.SubexpIndex("name")]),
		}
	}

	// https://regex101.com/r/JKjAYe/1
	files := regexp.MustCompile(`^(?P<size>[0-9]{1,}) (?P<name>[a-z](.[a-z]{0,}){0,})`)
	filesMatch := files.FindStringSubmatch(s)
	if len(filesMatch) > 0 {
		return &Element{
			Kind:     tree.File,
			Content:  strings.TrimSpace(filesMatch[files.SubexpIndex("name")]),
			Argument: strings.TrimSpace(filesMatch[files.SubexpIndex("size")]),
		}
	}

	return nil
}

func BuildTree(stack *stack.Stack[*tree.Node], root *tree.Node, value *Element) {
	switch value.Kind {
	case tree.Command:

		// change directory
		if value.Content == "cd" {

			if value.Argument == "/" {
				root.Name = value.Argument
				root.Kind = tree.Dir

				// first node in the stack
				stack.Push(root)

				if debug {
					fmt.Printf("  creating node: %v\n", root.Name)
				}
			}

			if value.Argument == ".." {
				n := stack.Pop()

				if debug {
					fmt.Printf("  leaving the node: %v\n", n.Name)

					c := stack.Peek()
					fmt.Printf("  current node: %v\n", c.Name)
				}

				break
			}

			if value.Argument != "/" && value.Argument != ".." {
				currentNode := stack.Peek()

				n := currentNode.FindChild(value.Argument, tree.Dir)
				stack.Push(n)

				if debug {
					fmt.Printf("  creating node: %v\n", n.Name)
				}

			}
		}

		// list
		if value.Content == "ls" {
			// nothing to do
			break
		}

	case tree.Dir:
		currentNode := stack.Peek()
		n := &tree.Node{
			Name: value.Argument,
			Kind: tree.Dir,
		}

		if currentNode.Children == nil {
			// first child
			currentNode.Children = []*tree.Node{n}
		} else {
			// append child
			currentNode.Children = append(currentNode.Children, n)
		}

		if debug {
			fmt.Printf("    adding dir: %v -> node: %v\n", n.Name, currentNode.Name)
		}

	case tree.File:
		size, err := strconv.Atoi(value.Argument)
		if err != nil {
			// ignore
			break
		}

		currentNode := stack.Peek()
		n := &tree.Node{
			Name: value.Content,
			Size: size,
			Kind: tree.File,
		}

		if currentNode.Children == nil {
			// first child
			currentNode.Children = []*tree.Node{n}
		} else {
			// append child
			currentNode.Children = append(currentNode.Children, n)
		}

		if debug {
			fmt.Printf("    adding file: %v -> node: %v\n", n.Name, currentNode.Name)
		}
	}
}

func init() {
	flag.StringVar(&input, "input", "input.txt", "input: a file path with the content of your https://adventofcode.com/2022/day/6/input")
	flag.IntVar(&maxSize, "maxSize", 100000, "Maximum size of directories to retrieve")
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

	root := &tree.Node{}

	stack := stack.New[*tree.Node]()
	rc := bufio.NewScanner(file)

	for rc.Scan() {
		line := rc.Text()

		value := ParseLine(line)
		if value == nil {
			// cannot be parser, jump
			continue
		}
		BuildTree(stack, root, value)
	}

	root.UpdateNodesDirSize()
	gotSize := root.GetSumOfSize(maxSize)

	fmt.Printf("What is the sum of the total sizes of those directories?: %v\n", gotSize)
}
