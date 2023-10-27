package ship

import (
	"github.com/christiangda/Golang-AdventOfCode2022/day_05/puzzle_01/stack"
)

type Crate struct {
	Content string
	Column  int
	Row     int
}

type Instruction struct {
	Text  string
	Moves int
	From  int
	To    int
}

type Ship struct {
	Stacks       []stack.Stack[Crate]
	Instructions []Instruction
}

func New() *Ship {
	return &Ship{
		Stacks:       make([]stack.Stack[Crate], 0),
		Instructions: make([]Instruction, 0),
	}
}

func (s *Ship) AddInstructions(instruction []Instruction) {
	s.Instructions = append(s.Instructions, instruction...)
}

func (s *Ship) AddCrates(crates []Crate) {
	// find the numbers of columns
	var columns int
	for _, crate := range crates {
		if crate.Column > columns {
			columns = crate.Column
		}
	}

	// add one stack per column
	s.Stacks = make([]stack.Stack[Crate], columns)

	// fill the stacks with the crates using column information
	for _, crate := range crates {
		s.Stacks[crate.Column-1].Push(crate)
	}

	// the stack were created from top to bottom, so let's reverse these
	for _, stack := range s.Stacks {
		stack.Reverse()
	}
}

func Crane(ship *Ship) {
	for _, i := range ship.Instructions {
		for j := 0; j < i.Moves; j++ {
			ship.Stacks[i.To-1].Push(ship.Stacks[i.From-1].Pop())
		}
	}
}
