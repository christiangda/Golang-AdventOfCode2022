package manager

import (
	"io"
	"strings"
	"testing"

	"github.com/christiangda/Golang-AdventOfCode2022/day_01/puzzle_01/elves"
)

func TestNew(t *testing.T) {
	m := New(nil, nil)

	if m == nil {
		t.Fatalf("Expected an instance of Manager, got %v", m)
	}
}

func TestAssign(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`
	t.Run("ValidInput", func(t *testing.T) {
		elves := elves.New()
		items := io.NopCloser(strings.NewReader(input))

		m := New(&elves, items)

		err := m.Assign()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		got := len(elves)
		if got != 5 {
			t.Fatalf("Expected 5 elves, got %d", got)
		}

		got = elves.TotalCalories()
		if got != 55000 {
			t.Fatalf("Expected 55000 calories, got %d", got)
		}

		got = elves.MostCalories().TotalCalories()
		if got != 24000 {
			t.Fatalf("Expected 10000 calories, got %d", got)
		}
	})
}
