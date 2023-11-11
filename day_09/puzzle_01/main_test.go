package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	H := Knot{Label: "H", position: [2]int{0, 0}}
	T := Knot{Label: "T", position: [2]int{0, 0}}
	Rope := NewRope(&H, &T)

	rc := bufio.NewScanner(strings.NewReader(input))
	for rc.Scan() {
		line := rc.Text()
		instructions := strings.Split(line, " ")
		if len(instructions) != 2 {
			fmt.Printf("Expected 2 elements in the line, got: %v\n", len(instructions))
			continue
		}

		direction := instructions[0]
		steps, err := strconv.Atoi(instructions[1])
		if err != nil {
			fmt.Printf("Error converting %v steps to integer\n", instructions[1])
		}

		Rope.Move(direction, steps)

		t.Logf("H -> Dir: %v, Steps: %v, Pos: %v, Mov: %v\n", H.direction, H.steps, H.position, H.movements)
		t.Logf("T -> Dir: %v, Steps: %v, Pos: %v, Mov: %v\n\n", T.direction, T.steps, T.position, T.movements)

	}

	HWantMov := 24
	if H.movements != HWantMov {
		t.Errorf("Expected: %v, got: %v\n", HWantMov, H.movements)
	}

	HWantPos := [2]int{2, 2}
	if H.position != HWantPos {
		t.Errorf("Expected: %v, got: %v\n", HWantPos, H.position)
	}

	TWantMov := 13
	if T.movements != TWantMov {
		t.Errorf("Expected: %v, got: %v\n", TWantMov, T.movements)
	}

	TWantPos := [2]int{1, 2}
	if T.position != TWantPos {
		t.Errorf("Expected: %v, got: %v\n", TWantPos, T.position)
	}
}

func TestMain(t *testing.T) {
	input := `L 2
D 1
L 2
U 2
R 1
D 2
U 1
R 2`

	H := Knot{Label: "H", position: [2]int{0, 0}}
	T := Knot{Label: "T", position: [2]int{0, 0}}
	Rope := NewRope(&H, &T)

	rc := bufio.NewScanner(strings.NewReader(input))
	for rc.Scan() {
		line := rc.Text()
		instructions := strings.Split(line, " ")
		if len(instructions) != 2 {
			fmt.Printf("Expected 2 elements in the line, got: %v\n", len(instructions))
			continue
		}

		direction := instructions[0]
		steps, err := strconv.Atoi(instructions[1])
		if err != nil {
			fmt.Printf("Error converting %v steps to integer\n", instructions[1])
		}

		Rope.Move(direction, steps)

		t.Logf("H -> Dir: %v, Steps: %v, Pos: %v, Mov: %v\n", H.direction, H.steps, H.position, H.movements)
		t.Logf("T -> Dir: %v, Steps: %v, Pos: %v, Mov: %v\n\n", T.direction, T.steps, T.position, T.movements)
	}

	HWantMov := 13
	if H.movements != HWantMov {
		t.Errorf("Expected: %v, got: %v\n", HWantMov, H.movements)
	}

	HWantPos := [2]int{-1, 0}
	if H.position != HWantPos {
		t.Errorf("Expected: %v, got: %v\n", HWantPos, H.position)
	}

	TWantMov := 6
	if T.movements != TWantMov {
		t.Errorf("Expected: %v, got: %v\n", TWantMov, T.movements)
	}

	TWantPos := [2]int{-2, 0}
	if T.position != TWantPos {
		t.Errorf("Expected: %v, got: %v\n", TWantPos, T.position)
	}
}
