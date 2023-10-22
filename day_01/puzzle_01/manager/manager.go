package manager

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/christiangda/Golang-AdventOfCode2022/day_01/puzzle_01/elves"
)

type Manager struct {
	elves *elves.Elves
	items io.Reader
}

func New(elves *elves.Elves, items io.Reader) *Manager {
	return &Manager{
		elves: elves,
		items: items,
	}
}

func (m *Manager) Assign() error {
	rc := bufio.NewScanner(m.items)

	calories := make([]int, 0)
	elfNumber := 1
	for rc.Scan() {
		line := rc.Text()

		// if line is not empty, add calories to the slice
		if line != "" {
			calorieString := strings.TrimSpace(line)
			calorie, err := strconv.Atoi(calorieString)
			if err != nil {
				return err
			}
			calories = append(calories, calorie)

			// jump to next line
			continue
		}

		// if line is empty, add elf to the slice
		elf := elves.NewElf(fmt.Sprintf("Elf %d", elfNumber))
		elf.AddCalories(calories)
		m.elves.Add(elf)

		// next elf values
		calories = make([]int, 0)
		elfNumber++
	}

	if err := rc.Err(); err != nil {
		if err != io.EOF {
			return err
		}
	}

	// add last elf, the scanner doesn't read the last line
	elf := elves.NewElf(fmt.Sprintf("Elf %d", elfNumber))
	elf.AddCalories(calories)
	m.elves.Add(elf)

	return nil
}
