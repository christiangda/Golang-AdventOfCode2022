package shapes

type Shaper interface {
	Value() int
	PlayWith(shape Shaper) int
	String() string
}

// outcome of the round
const (
	Lost int = 0
	Draw int = 3
	Won  int = 6

	Lose int = 1
	Win  int = 2
)

type Rock int

func (r Rock) Value() int {
	return int(r)
}

func (r Rock) String() string {
	return "Rock"
}

func (r Rock) PlayWith(shape Shaper) int {
	switch shape.String() {
	case "Rock":
		return Draw + shape.Value()
	case "Paper":
		return Won + shape.Value()
	case "Scissors":
		return Lost + shape.Value()
	}

	return 0
}

type Paper int

func (p Paper) Value() int {
	return int(p)
}

func (p Paper) String() string {
	return "Paper"
}

func (p Paper) PlayWith(shape Shaper) int {
	switch shape.String() {
	case "Rock":
		return Lost + shape.Value()
	case "Paper":
		return Draw + shape.Value()
	case "Scissors":
		return Won + shape.Value()
	}

	return 0
}

type Scissors int

func (s Scissors) Value() int {
	return int(s)
}

func (s Scissors) String() string {
	return "Scissors"
}

func (s Scissors) PlayWith(shape Shaper) int {
	switch shape.String() {
	case "Rock":
		return Won + shape.Value()
	case "Paper":
		return Lost + shape.Value()
	case "Scissors":
		return Draw + shape.Value()
	}

	return 0
}

// GetShape Return the shape according to the token s
func GetShape(s string) Shaper {
	switch s {
	case "A":
		return Rock(1)
	case "B":
		return Paper(2)
	case "C":
		return Scissors(3)
	}

	return nil
}

// GetShapeFromStrategy return the Shape need by t2 because t1 has a shape and the
// Strategy are defined as:
// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
func GetShapeFromStrategy(t1, t2 string) Shaper {
	switch t1 {
	case "A": // Rock
		switch t2 {
		case "X": // lose
			return Scissors(3)
		case "Y": // draw
			return Rock(1)
		case "Z": // win
			return Paper(2)
		}
	case "B": // Paper
		switch t2 {
		case "X": // lose
			return Rock(1)
		case "Y": // draw
			return Paper(2)
		case "Z": // win
			return Scissors(3)
		}
	case "C": // Scissors
		switch t2 {
		case "X": // lose
			return Paper(2)
		case "Y": // draw
			return Scissors(3)
		case "Z": // win
			return Rock(1)
		}
	}
	return nil
}
