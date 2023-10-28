package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	input := `mjqjpqmgbljsphdztnvjfqwrcgsmlb
bvwbjplbgvbhsrlpgdmjqwftvncz
nppdvjthqldpwncqszvftbrmjlhg
nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg
zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`

	t.Run("test", func(t *testing.T) {
		windowSize := 4
		var results []int

		rc := bufio.NewScanner(strings.NewReader(input))
		for rc.Scan() {
			line := rc.Text()
			t.Logf("line: %v\n", line)

			i := DetectSignal(line, windowSize)
			results = append(results, i)
		}

		// Then
		want := []int{7, 5, 6, 10, 11}

		if len(results) != len(want) {
			t.Errorf("Expected: %v, got: %v\n", len(want), len(results))
		}

		for i, got := range results {
			if got != want[i] {
				t.Errorf("Expected: %v, got: %v\n", want[i], got)
			}
		}
	})
}
