package tree

import (
	"fmt"

	"github.com/christiangda/Golang-AdventOfCode2022/day_07/puzzle_01/stack"
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

// GetSumOfSize given the maxSize, return the sum of all dirs size which
// the size is at most the maxSize
// implement Depth-first search
func (n *Node) GetSumOfSize(maxSize int) int {
	stack := stack.New[*Node]()
	stack.Push(n)

	var total int

	for stack.Size() > 0 {
		next := stack.Pop()

		if next.Kind == Dir && next.Size <= maxSize {
			total += next.Size
			// fmt.Printf("Node: %v, Size: %v: total: %v\n", next.Name, next.Size, total)
		}

		if len(next.Children) > 0 {
			for _, child := range next.Children {
				if child.Kind == Dir {
					// fmt.Printf("%v -> %v ->%v\n", child.Name, child.Kind, child.Size)
					stack.Push(child)
				}
			}
		}
	}

	return total
}
