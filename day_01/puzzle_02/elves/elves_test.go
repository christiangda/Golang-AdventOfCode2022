package elves

import (
	"bufio"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestElf(t *testing.T) {
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

	t.Run("NewElf", func(t *testing.T) {
		elf := NewElf("Elf 1")
		if elf.Name != "Elf 1" {
			t.Fatalf("Expected Elf 1, got %s", elf.Name)
		}

		if len(elf.Calories) != 0 {
			t.Fatalf("Expected 0 calories, got %d", len(elf.Calories))
		}
	})

	t.Run("AddCalories", func(t *testing.T) {
		items := bufio.NewScanner(strings.NewReader(input))

		elves := New()
		calories := make([]int, 0)

		elfNumber := 1
		for items.Scan() {
			line := items.Text()

			if line != "" {
				calorieString := strings.TrimSpace(line)
				calorie, err := strconv.Atoi(calorieString)
				if err != nil {
					t.Fatal(err)
				}

				t.Logf("Reading Calorie: %d\n", calorie)
				calories = append(calories, calorie)
				continue
			}

			// new elf name
			name := fmt.Sprintf("Elf %d", elfNumber)
			elf := NewElf(name)
			t.Logf("Adding: %s, Calories: %+v\n", name, calories)
			elf.AddCalories(calories)
			elves = append(elves, elf)

			// next elf values
			calories = make([]int, 0)
			elfNumber++
		}

		// add last elf, the scanner doesn't read the last line
		name := fmt.Sprintf("Elf %d", elfNumber)
		elf := NewElf(name)
		t.Logf("Adding: %s, Calories: %+v\n", name, calories)
		elf.AddCalories(calories)
		elves = append(elves, elf)

		if err := items.Err(); err != nil {
			t.Fatal(err)
		}

		if len(elves) != 5 {
			t.Fatalf("Expected 5 elves, got %d", len(elves))
		}

		// testing calories
		totalCalories := 0

		// elf 1
		for _, calorie := range elves[0].Calories {
			totalCalories += calorie
		}

		if totalCalories != 6000 {
			t.Fatalf("Expected 6000 calories, got %d", totalCalories)
		}

		// elf 2
		totalCalories = 0
		for _, calorie := range elves[1].Calories {
			totalCalories += calorie
		}

		if totalCalories != 4000 {
			t.Fatalf("Expected 4000 calories, got %d", totalCalories)
		}

		// elf 3
		totalCalories = 0
		for _, calorie := range elves[2].Calories {
			totalCalories += calorie
		}

		if totalCalories != 11000 {
			t.Fatalf("Expected 11000 calories, got %d", totalCalories)
		}

		// elf 4
		totalCalories = 0
		for _, calorie := range elves[3].Calories {
			totalCalories += calorie
		}

		if totalCalories != 24000 {
			t.Fatalf("Expected 24000 calories, got %d", totalCalories)
		}

		// elf 5
		totalCalories = 0
		for _, calorie := range elves[4].Calories {
			totalCalories += calorie
		}

		if totalCalories != 10000 {
			t.Fatalf("Expected 10000 calories, got %d", totalCalories)
		}
	})

	t.Run("TotalCalories", func(t *testing.T) {
		elves := New()

		// elf 1
		elf := NewElf("Elf 1")
		elf.AddCalories([]int{1000, 2000, 3000})
		elves = append(elves, elf)

		// elf 2
		elf = NewElf("Elf 2")
		elf.AddCalories([]int{4000})
		elves = append(elves, elf)

		// elf 3
		elf = NewElf("Elf 3")
		elf.AddCalories([]int{5000, 6000})
		elves = append(elves, elf)

		// elf 4
		elf = NewElf("Elf 4")
		elf.AddCalories([]int{7000, 8000, 9000})
		elves = append(elves, elf)

		// elf 5
		elf = NewElf("Elf 5")
		elf.AddCalories([]int{10000})
		elves = append(elves, elf)

		// elf 1
		totalCalories := elves[0].TotalCalories()
		if totalCalories != 6000 {
			t.Fatalf("Expected 6000 calories, got %d", totalCalories)
		}

		// elf 2
		totalCalories = elves[1].TotalCalories()
		if totalCalories != 4000 {
			t.Fatalf("Expected 4000 calories, got %d", totalCalories)
		}

		// elf 3
		totalCalories = elves[2].TotalCalories()
		if totalCalories != 11000 {
			t.Fatalf("Expected 11000 calories, got %d", totalCalories)
		}

		// elf 4
		totalCalories = elves[3].TotalCalories()
		if totalCalories != 24000 {
			t.Fatalf("Expected 24000 calories, got %d", totalCalories)
		}

		// elf 5
		totalCalories = elves[4].TotalCalories()
		if totalCalories != 10000 {
			t.Fatalf("Expected 10000 calories, got %d", totalCalories)
		}
	})
}

func TestElves(t *testing.T) {
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

	t.Run("New", func(t *testing.T) {
		elves := New()

		if len(elves) != 0 {
			t.Fatalf("Expected 0 elves, got %d", len(elves))
		}

		want := Elves{}
		if !reflect.DeepEqual(elves, want) {
			t.Fatalf("Expected %+v, got %+v", elves, want)
		}
	})

	t.Run("Add", func(t *testing.T) {
		elves := New()
		_ = input
		elves.Add(NewElf("name 1"))
		elves.Add(NewElf("name 2"))
		elves.Add(NewElf("name 3"))

		if len(elves) != 3 {
			t.Fatalf("Expected 3, got %d", len(elves))
		}

		if elves[0].Name != "name 1" {
			t.Fatalf("Expected name 1, got %s", elves[0].Name)
		}

		if elves[1].Name != "name 2" {
			t.Fatalf("Expected name 2, got %s", elves[1].Name)
		}

		if elves[2].Name != "name 3" {
			t.Fatalf("Expected name 3, got %s", elves[2].Name)
		}
	})

	t.Run("TotalCalories", func(t *testing.T) {
		elves := New()
		_ = input
		// elf 1
		elf1 := NewElf("Elf 1")
		elf1.AddCalories([]int{1000, 2000, 3000})

		// elf 2
		elf2 := NewElf("Elf 2")
		elf2.AddCalories([]int{4000})

		// elf 3
		elf3 := NewElf("Elf 3")
		elf3.AddCalories([]int{5000, 6000})

		// elf 4
		elf4 := NewElf("Elf 4")
		elf4.AddCalories([]int{7000, 8000, 9000})

		// elf 5
		elf5 := NewElf("Elf 5")
		elf5.AddCalories([]int{10000})

		elves.Add(elf1)
		elves.Add(elf2)
		elves.Add(elf3)
		elves.Add(elf4)
		elves.Add(elf5)

		if len(elves) != 5 {
			t.Fatalf("Expected 5, got %d", len(elves))
		}

		if elves[0].Name != "Elf 1" {
			t.Fatalf("Expected Elf 1, got %s", elves[0].Name)
		}

		if elves[1].Name != "Elf 2" {
			t.Fatalf("Expected Elf 2, got %s", elves[1].Name)
		}

		if elves[2].Name != "Elf 3" {
			t.Fatalf("Expected Elf 3, got %s", elves[2].Name)
		}

		if elves[3].Name != "Elf 4" {
			t.Fatalf("Expected Elf 4, got %s", elves[3].Name)
		}

		if elves[4].Name != "Elf 5" {
			t.Fatalf("Expected Elf 5, got %s", elves[4].Name)
		}

		if elves.TotalCalories() != 55000 {
			t.Fatalf("Expected 55000, got %v", elves.TotalCalories())
		}
	})

	t.Run("TopThree", func(t *testing.T) {
		elves := New()
		_ = input
		// elf 1
		elf1 := NewElf("Elf 1")
		elf1.AddCalories([]int{1000, 2000, 3000})

		// elf 2
		elf2 := NewElf("Elf 2")
		elf2.AddCalories([]int{4000})

		// elf 3
		elf3 := NewElf("Elf 3")
		elf3.AddCalories([]int{5000, 6000})

		// elf 4
		elf4 := NewElf("Elf 4")
		elf4.AddCalories([]int{7000, 8000, 9000})

		// elf 5
		elf5 := NewElf("Elf 5")
		elf5.AddCalories([]int{10000})

		elves.Add(elf1)
		elves.Add(elf2)
		elves.Add(elf3)
		elves.Add(elf4)
		elves.Add(elf5)

		if len(elves) != 5 {
			t.Fatalf("Expected 5, got %d", len(elves))
		}

		if elves[0].Name != "Elf 1" {
			t.Fatalf("Expected Elf 1, got %s", elves[0].Name)
		}

		if elves[1].Name != "Elf 2" {
			t.Fatalf("Expected Elf 2, got %s", elves[1].Name)
		}

		if elves[2].Name != "Elf 3" {
			t.Fatalf("Expected Elf 3, got %s", elves[2].Name)
		}

		if elves[3].Name != "Elf 4" {
			t.Fatalf("Expected Elf 4, got %s", elves[3].Name)
		}

		if elves[4].Name != "Elf 5" {
			t.Fatalf("Expected Elf 5, got %s", elves[4].Name)
		}

		if elves.TotalCalories() != 55000 {
			t.Fatalf("Expected 55000, got %v", elves.TotalCalories())
		}

		got := elves.TopThree()

		if got.TotalCalories() != 45000 {
			t.Fatalf("Expected 45000, got %v", got.TotalCalories())
		}
	})
}
