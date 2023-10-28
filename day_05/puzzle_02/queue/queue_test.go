package queue

import "testing"

func TestQueue(t *testing.T) {
	t.Run("test new", func(t *testing.T) {
		// given
		q := New[any]()

		// when

		// then
		l := q.size
		e := q.e

		if l != 0 {
			t.Errorf("Expected: %v, got: %v\n", 0, l)
		}

		if len(e) != 0 {
			t.Errorf("Expected: %v, got: %v\n", 0, len(e))
		}
	})

	t.Run("test Enqueue", func(t *testing.T) {
		// given
		q := New[any]()

		// when
		q.Enqueue("A")
		q.Enqueue("B")
		q.Enqueue("C")

		// then
		l := q.size
		e := q.e

		if l != 3 {
			t.Errorf("Expected: %v, got: %v\n", 3, l)
		}

		if len(e) != 3 {
			t.Errorf("Expected: %v, got: %v\n", 3, len(e))
		}

		if e[0] != "A" {
			t.Errorf("Expected: %v, got: %v\n", "A", e[0])
		}

		if e[1] != "B" {
			t.Errorf("Expected: %v, got: %v\n", "B", e[1])
		}

		if e[2] != "C" {
			t.Errorf("Expected: %v, got: %v\n", "C", e[2])
		}
	})

	t.Run("test Dequeue", func(t *testing.T) {
		// given
		q := New[any]()

		// when
		q.Enqueue("A")
		q.Enqueue("B")
		q.Enqueue("C")

		a := q.Dequeue()
		b := q.Dequeue()
		c := q.Dequeue()

		// then
		l := q.size
		e := q.e

		if l != 0 {
			t.Errorf("Expected: %v, got: %v\n", 0, l)
		}

		if len(e) != 0 {
			t.Errorf("Expected: %v, got: %v\n", 0, len(e))
		}

		if a != "A" {
			t.Errorf("Expected: %v, got: %v\n", "A", a)
		}

		if b != "B" {
			t.Errorf("Expected: %v, got: %v\n", "B", b)
		}

		if c != "C" {
			t.Errorf("Expected: %v, got: %v\n", "C", c)
		}
	})

	t.Run("test Size", func(t *testing.T) {
		// given
		q := New[any]()

		// when
		q.Enqueue("A")
		q.Enqueue("B")
		q.Enqueue("C")

		a := q.Dequeue()
		b := q.Dequeue()

		// then
		l := q.size
		e := q.e
		s := q.Size()

		if l != s && s != 2 {
			t.Errorf("Expected: %v, got: %v\n", 2, s)
		}

		if len(e) != 2 && len(e) != s {
			t.Errorf("Expected: %v, got: %v\n", 2, s)
		}

		if a != "A" {
			t.Errorf("Expected: %v, got: %v\n", "A", a)
		}

		if b != "B" {
			t.Errorf("Expected: %v, got: %v\n", "B", b)
		}
	})
}
