package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/christiangda/Golang-AdventOfCode2022/day_05/puzzle_02/ship"
)

func TestReadHeader(t *testing.T) {
	input := `
    [M]             [Z]     [V]
    [Z]     [P]     [L]     [Z] [J]
[S] [D]     [W]     [W]     [H] [Q]
[P] [V] [N] [D]     [P]     [C] [V]
[H] [B] [J] [V] [B] [M]     [N] [P]
[V] [F] [L] [Z] [C] [S] [P] [S] [G]
[F] [J] [M] [G] [R] [R] [H] [R] [L]
[G] [G] [G] [N] [V] [V] [T] [Q] [F]
 1   2   3   4   5   6   7   8   9

move 6 from 9 to 3
move 2 from 2 to 1
move 1 from 8 to 2
move 3 from 7 to 2
move 7 from 6 to 9
move 1 from 9 to 5
move 3 from 5 to 7
move 6 from 8 to 6
`

	// when
	rc := bufio.NewScanner(strings.NewReader(input))
	var crates []ship.Crate
	var instructions []ship.Instruction
	for rc.Scan() {
		if rc.Text() == "" {
			continue
		}
		c, i := Parse(rc.Text())
		crates = append(crates, c...)
		instructions = append(instructions, i...)
	}

	t.Logf("\nCreates -> %v\n\n", crates)
	t.Logf("\nInstructions -> %v\n\n", instructions)

	if len(crates) != 56 {
		t.Errorf("Expected: %v, got:%v", 56, len(crates))
	}

	if len(instructions) != 8 {
		t.Errorf("Expected: %v, got:%v", 8, len(instructions))
	}

	if instructions[0].To != 3 {
		t.Errorf("Expected: %v, got:%v", 3, instructions[0].To)
	}
}

func TestA(t *testing.T) {
	input := `
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`
	rc := bufio.NewScanner(strings.NewReader(input))
	var crates []ship.Crate
	var instructions []ship.Instruction
	for rc.Scan() {
		if rc.Text() == "" {
			continue
		}
		c, i := Parse(rc.Text())
		crates = append(crates, c...)
		instructions = append(instructions, i...)
	}

	t.Logf("\n\nCreates -> %v\n\n", crates)
	t.Logf("\n\nInstructions -> %v\n\n", instructions)

	if len(crates) != 6 {
		t.Errorf("Expected: %v, got:%v", 6, len(crates))
	}

	if len(instructions) != 4 {
		t.Errorf("Expected: %v, got:%v", 4, len(instructions))
	}

	if instructions[0].To != 1 {
		t.Errorf("Expected: %v, got:%v", 1, instructions[0].To)
	}

	s := ship.New()
	s.AddCrates(crates)
	s.AddInstructions(instructions)

	t.Logf("\n\n Ship Queues: -> %+v\n\n", s.Queues)
	t.Logf("\n\n Ship Instructions: -> %+v\n\n", s.Instructions)

	if len(s.Queues) != 3 {
		t.Errorf("Expected: %v, got: %v", 3, len(s.Queues))
	}

	if len(s.Instructions) != 4 {
		t.Errorf("Expected: %v, got: %v", 4, len(s.Instructions))
	}

	got := s.Queues[1].Dequeue().Content
	if got != "D" {
		t.Errorf("Expected: %v, got: %v", "D", got)
	}

	got = s.Queues[1].Dequeue().Content
	if got != "C" {
		t.Errorf("Expected: %v, got: %v", "C", got)
	}

	got = s.Queues[1].Dequeue().Content
	if got != "M" {
		t.Errorf("Expected: %v, got: %v", "M", got)
	}

	got = s.Queues[2].Dequeue().Content
	if got != "P" {
		t.Errorf("Expected: %v, got: %v", "P", got)
	}

	got = s.Queues[0].Dequeue().Content
	if got != "N" {
		t.Errorf("Expected: %v, got: %v", "N", got)
	}

	to := s.Instructions[0].To
	from := s.Instructions[0].From
	if to != 1 {
		t.Errorf("Expected: %v, got: %v", 1, to)
	}
	if from != 2 {
		t.Errorf("Expected: %v, got: %v", 2, from)
	}

	to = s.Instructions[1].To
	from = s.Instructions[1].From
	if to != 3 {
		t.Errorf("Expected: %v, got: %v", 3, to)
	}
	if from != 1 {
		t.Errorf("Expected: %v, got: %v", 1, from)
	}
}

func TestB(t *testing.T) {
	input := `
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`
	rc := bufio.NewScanner(strings.NewReader(input))
	var crates []ship.Crate
	var instructions []ship.Instruction
	for rc.Scan() {
		if rc.Text() == "" {
			continue
		}
		c, i := Parse(rc.Text())
		crates = append(crates, c...)
		instructions = append(instructions, i...)
	}

	t.Logf("\n\nCreates -> %v\n\n", crates)
	t.Logf("\n\nInstructions -> %v\n\n", instructions)

	if len(crates) != 6 {
		t.Errorf("Expected: %v, got:%v", 3, len(crates))
	}

	if len(instructions) != 4 {
		t.Errorf("Expected: %v, got:%v", 4, len(instructions))
	}

	if instructions[0].To != 1 {
		t.Errorf("Expected: %v, got:%v", 1, instructions[0].To)
	}

	s := ship.New()
	s.AddCrates(crates)
	s.AddInstructions(instructions)

	if len(s.Queues) != 3 {
		t.Errorf("Expected: %v, got: %v", 3, len(s.Queues))
	}

	if len(s.Instructions) != 4 {
		t.Errorf("Expected: %v, got: %v", 4, len(s.Instructions))
	}

	t.Logf("\n\n Ship Queues: -> %+v\n\n", s.Queues)
	t.Logf("\n\n Ship Instructions: -> %+v\n\n", s.Instructions)

	ship.Crane(s)

	var end []string
	for _, q := range s.Queues {
		end = append(end, q.Dequeue().Content)
	}

	if len(end) != 3 {
		t.Errorf("Expected: %v, got: %v", 3, len(end))
	}

	if end[0] != "M" {
		t.Errorf("Expected: %v, got: %v", "M", end[0])
	}
	if end[1] != "C" {
		t.Errorf("Expected: %v, got: %v", "C", end[1])
	}
	if end[2] != "D" {
		t.Errorf("Expected: %v, got: %v", "D", end[2])
	}

	if strings.Join(end, "") != "MCD" {
		t.Errorf("Expected: %v, got: %v", "MCD", strings.Join(end, ""))
	}
}
