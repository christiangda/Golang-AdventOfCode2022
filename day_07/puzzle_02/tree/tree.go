package tree

import (
	"fmt"
	"math"

	"github.com/christiangda/Golang-AdventOfCode2022/day_07/puzzle_02/stack"
)

const (
	Command int = iota
	File
	Dir
)

type Node struct {
	Name     string
	Size     int
	Kind     int
	Children []*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("\n{\n  Name: %s\n  Size: %v\n  Kind: %v\n  Children: %v\n}", n.Name, n.Size, n.Kind, n.Children)
}

// FindChild find in the child if the node with the name and kind exits
func (n *Node) FindChild(name string, kind int) *Node {
	if len(n.Children) > 0 {
		for _, child := range n.Children {
			if child.Name == name && child.Kind == kind {
				return child
			}
		}
	}

	return nil
}

// GetSumOfFilesSize traversing the tree and sum all the files size
// implement Depth-first search
func (n *Node) GetSumOfFilesSize() int {
	stack := stack.New[*Node]()
	stack.Push(n)

	var total int

	for stack.Size() > 0 {
		next := stack.Pop()

		if next.Kind == File {
			total += next.Size
		}

		if len(next.Children) > 0 {
			for _, child := range next.Children {
				stack.Push(child)
			}
		}
	}

	return total
}

// UpdateNodesDirSize update all the nodes kind=Dir with the accumulated size of nodes Kind=Files
// Traversing in Post order transversal
func (n *Node) UpdateNodesDirSize() {
	var acc int

	if n != nil {
		for _, c := range n.Children {
			c.UpdateNodesDirSize()
			acc += c.Size
			// fmt.Printf("Node: %v, Size: %v, kind: %v\n", c.Name, c.Size, c.Kind)
		}
	}

	if n.Kind == Dir {
		n.Size = acc
	}

	// fmt.Printf("Node: %v, Size: %v, kind: %v, acc:= %v\n", n.Name, n.Size, n.Kind, acc)
}

// GetGreaterOrEqual given the size, return the sorted closes node with size
// greater or equal to size from the tree
// implement Depth-first search
func (n *Node) GetGreaterOrEqual(size int) *Node {
	stack := stack.New[*Node]()

	var result *Node
	stack.Push(n)

	max := math.MaxInt

	for stack.Size() > 0 {
		next := stack.Pop()

		if next.Kind == Dir && next.Size >= size && next.Size < max {
			result = next
			max = result.Size
			// fmt.Printf("Adding candidate -> Name: %v, Size: %v, max: %v\n", next.Name, next.Size, max)
		}

		if len(next.Children) > 0 {
			for _, child := range next.Children {
				stack.Push(child)
			}
		}
	}

	return result
}
