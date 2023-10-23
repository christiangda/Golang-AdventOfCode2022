package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestGetPriority(t *testing.T) {
	// given
	chars := []string{"a", "A", "b", "Z", "c", "Y"}
	want := []int{1, 27, 2, 52, 3, 51}

	for i, c := range chars {

		got := Priority(c)

		if got != want[i] {
			t.Errorf("Expected %v, got %v", want[i], got)
		}
	}
}

func TestCommonItem(t *testing.T) {
	t.Run("first rucksack", func(t *testing.T) {
		// given
		rucksack1 := "vJrwpWtwJgWrhcsFMMfFFhFp"
		rucksack2 := "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"
		rucksack3 := "PmmdzqPrVvPwwTWBwg"

		want := "r"

		// got
		got := CommonItem(rucksack1, rucksack2, rucksack3)

		if got != want {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("first rucksack", func(t *testing.T) {
		// given
		rucksack1 := "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"
		rucksack2 := "ttgJtRGJQctTZtZT"
		rucksack3 := "CrZsJsPPZsGzwwsLwLmpwMDw"

		want := "Z"

		// got
		got := CommonItem(rucksack1, rucksack2, rucksack3)

		if got != want {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("first rucksack", func(t *testing.T) {
		// given
		rucksack1 := "shzsFcPssFhjFssBzdpRcNHNZrpdJdJVJZ"
		rucksack2 := "fwvMCntfCCbSbSbtDgDNrDtDtJHZVH"
		rucksack3 := "GbCwwbwwnGrLhBzjFFFsWPhL"

		want := "r"

		// got
		got := CommonItem(rucksack1, rucksack2, rucksack3)

		if got != want {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}

func TestSumPriorities(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

	rc := bufio.NewScanner(strings.NewReader(input))
	var got []int
	var rucksacks []string
	groupOf := 3
	line := 1
	for rc.Scan() {

		if line%groupOf != 0 || line%groupOf == 0 {
			rucksacks = append(rucksacks, rc.Text())
		}

		if len(rucksacks) == groupOf {
			ci := CommonItem(rucksacks[0], rucksacks[1], rucksacks[2])
			p := Priority(ci)
			got = append(got, p)

			// clean
			rucksacks = make([]string, 0)
		}
		line += 1
	}

	// r, Z
	want := []int{18, 52}

	if len(got) != len(want) {
		t.Errorf("len of arrays: Expected %v, got %v\n", len(want), len(got))
	}

	var sumGot int
	for i, g := range got {
		if g != want[i] {
			t.Errorf("Expected %v, got %v\n", want[i], g)
		}
		sumGot += g
	}

	if sumGot != 70 {
		t.Errorf("Expected %v, got %v\n", 70, sumGot)
	}
}
