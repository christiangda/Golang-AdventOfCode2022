package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/christiangda/Golang-AdventOfCode2022/day_07/puzzle_01/tree"
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
	BuildTree(rc, root)

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

	BuildTree(rc, root)

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
	BuildTree(rc, root)

	root.UpdateNodesDirSize()
	gotSize := root.GetSumOfSize(100000)

	t.Logf("Tree: %+v", root)

	wantSize := 95437

	if gotSize != wantSize {
		t.Errorf("Expected: %+v, got: %+v", wantSize, gotSize)
	}
}
