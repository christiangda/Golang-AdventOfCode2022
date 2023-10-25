package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestAreFullyContained(t *testing.T) {
	t.Run("Input From Webpage", func(t *testing.T) {
		input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

		rc := bufio.NewScanner(strings.NewReader(input))
		var got int
		for rc.Scan() {
			sections := strings.Split(rc.Text(), ",")

			if len(sections) != 2 {
				t.Fatalf("bad entry from sections, expected 2, got %v", len(sections))
			}

			ok, err := areFullyContained(sections[0], sections[1])
			if err != nil {
				t.Fatalf("Error returned from function, err: %v", err)
			}

			if ok {
				got += 1
			}
		}

		// 2-8,3-7 and 6-6,4-6
		want := 2

		if want != got {
			t.Errorf("Expected %v, got %v\n", want, got)
		}
	})

	t.Run("Input with corner case", func(t *testing.T) {
		// 2-4,2-4 mutual contained, but count only 1
		// 8-8,8-8 mutual contained and only 1 section, this return only 1
		// 2-7,2-4 first contain second, same start, this return only 1
		// 7-7,1-7 second contain first and share corner, this return 1
		// 1-1,8-8 return 0
		// 7-8,1-2 return 0

		input := `2-4,2-4
8-8,8-8
2-7,2-4
7-7,1-7
1-1,8-8
7-8,1-2`

		rc := bufio.NewScanner(strings.NewReader(input))
		var got int
		for rc.Scan() {
			sections := strings.Split(rc.Text(), ",")

			if len(sections) != 2 {
				t.Fatalf("bad entry from sections, expected 2, got %v", len(sections))
			}

			ok, err := areFullyContained(sections[0], sections[1])
			if err != nil {
				t.Fatalf("Error returned from function, err: %v", err)
			}

			if ok {
				got += 1
			}
		}

		want := 4

		if want != got {
			t.Errorf("Expected %v, got %v\n", want, got)
		}
	})
}
