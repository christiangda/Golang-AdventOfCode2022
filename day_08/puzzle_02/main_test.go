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

	var maxScenicScore int
	rows := len(woods.trees)
	cols := len(woods.trees[0])

	for y := 1; y <= rows-2; y++ {
		for x := 1; x <= cols-2; x++ {
			tree := Tree{x: x, y: y}
			scenicScore := tree.ScenicScore(woods.trees)
			t.Logf("tree: %v, sc: %v\n", woods.trees[y][x], scenicScore)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	t.Logf("lines processed: %v\n", lineNum)
	t.Logf("Woods: %v\n", woods)

	want := 13
	if maxScenicScore != want {
		t.Errorf("Expected: %v, got: %v\n", want, maxScenicScore)
	}
}

func TestGetBlockerOrEdgeIndex(t *testing.T) {
	type args struct {
		treeHeight int
		treeRow    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nil",
			args: args{
				treeHeight: 3,
				treeRow:    nil,
			},
			want: 0,
		},
		{
			name: "empty",
			args: args{
				treeHeight: 3,
				treeRow:    make([]int, 0),
			},
			want: 0,
		},
		{
			name: "same height next to",
			args: args{
				treeHeight: 3,
				treeRow:    []int{3, 1, 2, 3},
			},
			want: 1,
		},
		{
			name: "higher next to it",
			args: args{
				treeHeight: 3,
				treeRow:    []int{5, 1, 2, 3},
			},
			want: 1,
		},
		{
			name: "same in the border",
			args: args{
				treeHeight: 3,
				treeRow:    []int{1, 1, 2, 3},
			},
			want: 4,
		},
		{
			name: "higher in the border",
			args: args{
				treeHeight: 3,
				treeRow:    []int{1, 1, 2, 5},
			},
			want: 4,
		},
		{
			name: "all lower",
			args: args{
				treeHeight: 3,
				treeRow:    []int{1, 1, 2, 2, 2, 1, 1},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBlockerOrEdgeIndex(tt.args.treeHeight, tt.args.treeRow); got != tt.want {
				t.Errorf("GetBlockerOrEdgeIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
