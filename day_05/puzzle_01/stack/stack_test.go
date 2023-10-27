package stack

import "testing"

func TestStack(t *testing.T) {
	t.Run("test new", func(t *testing.T) {
		// given
		s := New[any]()

		// when

		// want
		w := 0
		top := 0

		// got

		g := len(s.e)

		if w != g {
			t.Fatalf("Expected %v, got %v", w, g)
		}

		if top != s.top {
			t.Fatalf("Expected %v, got %v", w, g)
		}
	})

	t.Run("test IsEmpty", func(t *testing.T) {
		// given
		s := New[any]()

		// when

		// want
		w := 0

		// got

		g := len(s.e)
		e := s.IsEmpty()

		if w != g {
			t.Fatalf("Expected %v, got %v", w, g)
		}

		if e != true {
			t.Fatalf("Expected %v, got %v", w, g)
		}
	})

	t.Run("test push", func(t *testing.T) {
		// given
		s := New[any]()

		// when
		s.Push("A")
		s.Push("B")
		s.Push("C")
		s.Push("D")

		// want
		w := 4

		// got

		g := len(s.e)

		if w != g {
			t.Fatalf("Expected %v, got %v", w, g)
		}
	})

	t.Run("test pop", func(t *testing.T) {
		// given
		s := New[any]()

		// when
		v := s.Pop()
		s.Pop()
		s.Pop()

		// want
		l := 0

		// got

		g := len(s.e)

		if l != g {
			t.Fatalf("Expected %v, got %v", l, g)
		}

		if v != nil {
			t.Fatalf("Expected %v, got %v", v, g)
		}
	})

	t.Run("test push/pop", func(t *testing.T) {
		// given
		s := New[string]()

		// when
		s.Push("A")
		s.Push("B")
		s.Push("C")
		s.Push("D")

		s.Pop()
		s.Pop()
		s.Pop()

		s.Push("E")
		s.Push("F")
		s.Pop()
		s.Push("G")
		s.Push("H")
		s.Push("I")

		// want
		i := s.Pop()
		w := 4

		// got
		g := len(s.e)

		if w != g {
			t.Fatalf("Expected %v, got %v", w, g)
		}

		if i != "I" {
			t.Fatalf("Expected %v, got %v", "I", i)
		}
	})
}
