package ship

import (
	"github.com/christiangda/Golang-AdventOfCode2022/day_05/puzzle_02/queue"
)

type Crate struct {
	Content string
	Column  int
}

type Instruction struct {
	Text  string
	Moves int
	From  int
	To    int
}

type Ship struct {
	Queues       []queue.Queue[Crate]
	Instructions []Instruction
}

func New() *Ship {
	return &Ship{
		Queues:       make([]queue.Queue[Crate], 0),
		Instructions: make([]Instruction, 0),
	}
}

func (s *Ship) AddInstructions(instruction []Instruction) {
	s.Instructions = append(s.Instructions, instruction...)
}

func (s *Ship) AddCrates(crates []Crate) {
	// find the numbers of columns
	var columns int
	for _, crate := range crates {
		if crate.Column > columns {
			columns = crate.Column
		}
	}

	// add one queue per column
	s.Queues = make([]queue.Queue[Crate], columns)

	// fill the stacks with the crates using column information
	for _, crate := range crates {
		s.Queues[crate.Column-1].Enqueue(crate)
	}
}

func Crane(ship *Ship) {
	for _, i := range ship.Instructions {
		n := queue.New[Crate]()

		// fmt.Printf("\n\nship Instruction: %+v\n\n", i.Text)
		for j := 0; j < i.Moves; j++ {
			// put the extractions from the source queue in the new queue
			n.Enqueue(ship.Queues[i.From-1].Dequeue())
		}

		// enqueue in the new queue the destination queue
		for ship.Queues[i.To-1].Size() > 0 {
			n.Enqueue(ship.Queues[i.To-1].Dequeue())
		}

		// the ordered new queue (movements + current queue) is the new destination queue for the position To
		// of the instructions
		ship.Queues[i.To-1] = *n
		// fmt.Printf("\n\nship Queues: %+v\n\n", ship.Queues)
	}
}
