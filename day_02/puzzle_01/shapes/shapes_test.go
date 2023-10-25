package shapes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestRock(t *testing.T) {
	r := Rock(1)

	if r.Value() != 1 {
		t.Fatalf("Expected 1, got %v", r.Value())
	}

	if r.String() != "Rock" {
		t.Fatalf("Expected Rock, got %v", r.String())
	}
}

func TestPaper(t *testing.T) {
	p := Paper(2)

	if p.Value() != 2 {
		t.Fatalf("Expected 2, got %v", p.Value())
	}

	if p.String() != "Paper" {
		t.Fatalf("Expected Paper, got %v", p.String())
	}
}

func TestScissors(t *testing.T) {
	s := Scissors(3)

	if s.Value() != 3 {
		t.Fatalf("Expected 2, got %v", s.Value())
	}

	if s.String() != "Scissors" {
		t.Fatalf("Expected Scissors, got %v", s.String())
	}
}

func TestPlays(t *testing.T) {
	// Rock
	t.Run("Rock vs Rock", func(t *testing.T) {
		p1 := Rock(1)
		p2 := Rock(1)

		got := p1.PlayWith(p2)

		if got != Draw+p1.Value() {
			t.Fatalf("Expected Draw, got %v", got)
		}
	})

	t.Run("Rock vs Paper", func(t *testing.T) {
		p1 := Rock(1)
		p2 := Paper(2)

		got := p1.PlayWith(p2)

		if got != Won+p2.Value() {
			t.Fatalf("Expected Won, got %v", got)
		}
	})

	t.Run("Rock vs Scissors", func(t *testing.T) {
		p1 := Rock(1)
		p2 := Scissors(3)

		got := p1.PlayWith(p2)

		if got != Lost+p2.Value() {
			t.Fatalf("Expected Won, got %v", got)
		}
	})

	// Paper
	t.Run("Paper vs Rock", func(t *testing.T) {
		p1 := Paper(2)
		p2 := Rock(1)

		got := p1.PlayWith(p2)

		if got != Lost+p2.Value() {
			t.Fatalf("Expected Draw, got %v", got)
		}
	})

	t.Run("Paper vs Paper", func(t *testing.T) {
		p1 := Paper(2)
		p2 := Paper(2)

		got := p1.PlayWith(p2)

		if got != Draw+p1.Value() {
			t.Fatalf("Expected Lost, got %v", got)
		}
	})

	t.Run("Paper vs Scissors", func(t *testing.T) {
		p1 := Paper(2)
		p2 := Scissors(3)

		got := p1.PlayWith(p2)

		if got != Won+p2.Value() {
			t.Fatalf("Expected Won, got %v", got)
		}
	})

	// Scissors
	t.Run("Scissors vs Rock", func(t *testing.T) {
		p1 := Scissors(3)
		p2 := Rock(1)

		got := p1.PlayWith(p2)

		if got != Won+p2.Value() {
			t.Fatalf("Expected Draw, got %v", got)
		}
	})

	t.Run("Scissors vs Paper", func(t *testing.T) {
		p1 := Scissors(3)
		p2 := Paper(2)

		got := p1.PlayWith(p2)

		if got != Lost+p2.Value() {
			t.Fatalf("Expected Lost, got %v", got)
		}
	})

	t.Run("Scissors vs Scissors", func(t *testing.T) {
		p1 := Scissors(3)
		p2 := Scissors(3)

		got := p1.PlayWith(p2)

		if got != Draw+p1.Value() {
			t.Fatalf("Expected Won, got %v", got)
		}
	})
}

func TestGetShape(t *testing.T) {
	input := `A Y
B X
C Z`

	rc := bufio.NewScanner(strings.NewReader(input))
	var results []int
	for rc.Scan() {
		line := rc.Text()
		round := strings.Split(line, " ")
		if len(round) != 2 {
			t.Errorf("Round do not have 2 players, has %v", len(round))
		}

		p1Token := round[0]
		p2Token := round[1]

		p1 := GetShape(p1Token)
		if p1 == nil {
			fmt.Print("Player 1 shape is null")
			os.Exit(1)
		}

		p2 := GetShape(p2Token)
		if p2 == nil {
			fmt.Print("Player 2 shape is null")
			os.Exit(1)
		}

		result := p1.PlayWith(p2)
		t.Logf("p1: {token: %v, name: %v, value: %v}, p2: {token: %v, name: %v, value: %v} -> round result: %v\n", p1Token, p1, p1.Value(), p2Token, p2, p2.Value(), result)

		results = append(results, result)
	}

	if len(results) != 3 {
		t.Errorf("Expected 3, got %v", len(results))
	}

	got1 := results[0]
	got2 := results[1]
	got3 := results[2]

	if got1 != 8 {
		t.Errorf("Expected 8, got %v", got1)
	}

	if got2 != 1 {
		t.Errorf("Expected 1, got %v", got2)
	}

	if got3 != 6 {
		t.Errorf("Expected 6, got %v", got3)
	}

	if (got1 + got2 + got3) != 15 {
		t.Errorf("Expected 15, got %v", (got1 + got2 + got3))
	}
}
