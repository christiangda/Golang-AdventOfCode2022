package tree

import (
	"reflect"
	"testing"
)

func TestFindByName(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		// given
		n := &Node{}

		// when
		got := n.FindByName("none", Dir)

		// want
		var want *Node

		if got != want {
			t.Errorf("Expected: %v, got: %v", want, got)
		}
	})

	t.Run("valid", func(t *testing.T) {
		// given
		n := &Node{
			Name: "root",
			Children: []*Node{
				{
					Name: "one",
					Children: []*Node{
						{Name: "three"},
					},
				},
				{Name: "two"},
			},
		}

		// when
		got := n.FindByName("three", Dir)

		// want
		want := &Node{
			Name: "three",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected: %v, got: %v", want, got)
		}
	})
}
