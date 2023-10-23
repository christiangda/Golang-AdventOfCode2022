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
		rucksack := "vJrwpWtwJgWrhcsFMMfFFhFp"
		length := len(rucksack)

		want := "p"

		c1 := rucksack[:length/2]
		c2 := rucksack[length/2:]

		// got
		got := CommonItem(c1, c2)

		if got != want {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("second rucksack", func(t *testing.T) {
		// given
		rucksack := "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"
		length := len(rucksack)

		want := "L"

		c1 := rucksack[:length/2]
		c2 := rucksack[length/2:]

		// got
		got := CommonItem(c1, c2)

		if got != want {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("third rucksack", func(t *testing.T) {
		// given
		rucksack := "PmmdzqPrVvPwwTWBwg"
		length := len(rucksack)

		want := "P"

		c1 := rucksack[:length/2]
		c2 := rucksack[length/2:]

		// got
		got := CommonItem(c1, c2)

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

	var got []int
	rc := bufio.NewScanner(strings.NewReader(input))
	for rc.Scan() {
		rucksack := rc.Text()
		length := len(rucksack)

		compartment1 := rucksack[:length/2]
		compartment2 := rucksack[length/2:]

		ci := CommonItem(compartment1, compartment2)
		p := Priority(ci)

		got = append(got, p)
	}

	want := []int{16, 38, 42, 22, 20, 19}

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

	if sumGot != 157 {
		t.Errorf("Expected %v, got %v\n", sumGot, 157)
	}
}
