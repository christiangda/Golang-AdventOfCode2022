package elves

// Elf represents a single elf and its calories
type Elf struct {
	Name     string
	Calories []int
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
}

// TotalCalories returns the total calories for an elf
func (e Elf) TotalCalories() int {
	var total int
	for _, calorie := range e.Calories {
		total += calorie
	}
	return total
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

// MostCalories returns the elf with the most calories
func (e Elves) MostCalories() *Elf {
	var mostCalories *Elf

	for _, elf := range e {
		if mostCalories == nil {
			mostCalories = elf
			continue
		}

		if elf.TotalCalories() > mostCalories.TotalCalories() {
			mostCalories = elf
		}
	}
	return mostCalories
}
