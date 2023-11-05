package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestAddRowOfTrees(t *testing.T) {
	input := `30373
25512
65332
33549
35390`

	woods := NewWoods()
	lineNum := 1
	rc := bufio.NewScanner(strings.NewReader(input))
	for rc.Scan() {
		woods.AddRowOfTrees(rc.Text())
		lineNum += 1
	}

	if len(woods.trees) != 5 {
		t.Errorf("Expected: %v, got: %v\n", 5, len(woods.trees))
	}

	if len(woods.trees[0]) != 5 {
		t.Errorf("Expected: %v, got: %v\n", 5, len(woods.trees[0]))
	}

	gotLine1 := woods.trees[0]
	wantLine1 := [...]int{3, 0, 3, 7, 3}
	if reflect.DeepEqual(gotLine1, wantLine1) {
		t.Errorf("Expected: %v, got: %v\n", wantLine1, gotLine1)
	}

	if len(woods.trees[1]) != 5 {
		t.Errorf("Expected: %v, got: %v\n", 5, len(woods.trees[1]))
	}
	gotLine2 := woods.trees[1]
	wantLine2 := [...]int{2, 5, 5, 1, 2}
	if reflect.DeepEqual(gotLine2, wantLine2) {
		t.Errorf("Expected: %v, got: %v\n", wantLine2, gotLine2)
	}

	if len(woods.trees[2]) != 5 {
		t.Errorf("Expected: %v, got: %v\n", 5, len(woods.trees[2]))
	}
	gotLine3 := woods.trees[2]
	wantLine3 := [...]int{6, 5, 3, 3, 2}
	if reflect.DeepEqual(gotLine3, wantLine3) {
		t.Errorf("Expected: %v, got: %v\n", wantLine3, gotLine3)
	}

	gotLine1 = woods.trees[0]
	wantLine1 = [...]int{6, 5, 3, 3, 2}
	if reflect.DeepEqual(gotLine1, wantLine1) {
		t.Errorf("Expected: %v, got: %v\n", wantLine1, gotLine1)
	}

	gotLine2 = woods.trees[1]
	wantLine2 = [...]int{3, 3, 5, 4, 9}
	if reflect.DeepEqual(gotLine2, wantLine2) {
		t.Errorf("Expected: %v, got: %v\n", wantLine2, gotLine2)
	}

	gotLine3 = woods.trees[2]
	wantLine3 = [...]int{3, 5, 3, 9, 0}
	if reflect.DeepEqual(gotLine3, wantLine3) {
		t.Errorf("Expected: %v, got: %v\n", wantLine3, gotLine3)
	}
}

func TestMainLogic(t *testing.T) {
	input := `30373
25512
65332
33549
35390`

	woods := NewWoods()
	lineNum := 1

	rc := bufio.NewScanner(strings.NewReader(input))
	for rc.Scan() {
		woods.AddRowOfTrees(rc.Text())
		lineNum += 1
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
	t.Logf("lines processed: %v\n", lineNum)
	t.Logf("Woods: %v\n", woods)
	t.Logf("Visible trees inside: %v\n", visibleTreesInside)
	t.Logf("Visible trees outside: %v\n", visibleTreesOutside)
	t.Logf("Visible trees: %v\n", visibleTrees)

	if visibleTrees != 21 {
		t.Errorf("Expected: %v, got: %v\n", 21, visibleTrees)
	}
}
