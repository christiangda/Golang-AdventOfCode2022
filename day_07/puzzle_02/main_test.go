package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/christiangda/Golang-AdventOfCode2022/day_07/puzzle_02/stack"
	"github.com/christiangda/Golang-AdventOfCode2022/day_07/puzzle_02/tree"
)

func TestGetSumOfFilesSize(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	root := &tree.Node{}

	rc := bufio.NewScanner(strings.NewReader(input))
	stack := stack.New[*tree.Node]()

	for rc.Scan() {
		line := rc.Text()

		value := ParseLine(line)
		if value == nil {
			// cannot be parser, jump
			continue
		}
		BuildTree(stack, root, value)
	}

	gotSize := root.GetSumOfFilesSize()

	wantSize := 48381165
	t.Logf("Tree: %+v", root)

	if gotSize != wantSize {
		t.Errorf("Expected: %+v, got: %+v", wantSize, gotSize)
	}

	if len(root.Children) != 4 {
		t.Errorf("Expected: %+v, got: %+v", 4, len(root.Children))
	}

	if root.Children[0].Name != "a" {
		t.Errorf("Expected: %+v, got: %+v", "a", root.Children[0].Name)
	}
}

func TestUpdateDirSize(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	root := &tree.Node{}

	rc := bufio.NewScanner(strings.NewReader(input))
	stack := stack.New[*tree.Node]()

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
	gotSize := root.Size

	wantSize := 48381165
	t.Logf("Tree: %+v", root)

	if gotSize != wantSize {
		t.Errorf("Expected: %+v, got: %+v", wantSize, gotSize)
	}
}

func TestGetSumOfSize(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	root := &tree.Node{}

	rc := bufio.NewScanner(strings.NewReader(input))
	stack := stack.New[*tree.Node]()

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
	got := root.GetGreaterOrEqual(8381165)

	t.Logf("Tree: %+v", root)

	wantSize := 24933642

	if got.Size != wantSize {
		t.Errorf("Expected: %+v, got: %+v", wantSize, got.Size)
	}

	if got.Name != "d" {
		t.Errorf("Expected: %+v, got: %+v", "d", got.Name)
	}
}
