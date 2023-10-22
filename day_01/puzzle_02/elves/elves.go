package elves

import (
	"sort"
)

// Elf represents a single elf and its calories
type Elf struct {
	Name          string
	Calories      []int
	totalCalories int
}

// NewElf creates a new elf
func NewElf(name string) *Elf {
	return &Elf{
		Name: name,
	}
}

// AddCalories adds calories to an elf
func (e *Elf) AddCalories(calories []int) {
	e.Calories = calories
	for _, calorie := range calories {
		e.totalCalories += calorie
	}
}

// TotalCalories returns the total calories for an elf
func (e *Elf) TotalCalories() int {
	return e.totalCalories
}

// Elves represents a collection of elves
type Elves []*Elf

// NewElves creates a new collection of elves
func New() Elves {
	return make(Elves, 0)
}

// Add adds an elf to the collection
func (e *Elves) Add(elf *Elf) {
	*e = append(*e, elf)
}

// TotalCalories returns the total calories for all elves
func (e Elves) TotalCalories() int {
	var total int
	for _, elf := range e {
		total += elf.TotalCalories()
	}
	return total
}

// TopThree return the Top Three elves with most calories
func (e Elves) TopThree() Elves {
	ret := New()

	if len(e) >= 3 {
		// sort by calories in desc order
		sort.SliceStable(e, func(i, j int) bool {
			return e[i].totalCalories > e[j].totalCalories
		})

		// get the three first
		ret.Add(e[0])
		ret.Add(e[1])
		ret.Add(e[2])
	}

	return ret
}
