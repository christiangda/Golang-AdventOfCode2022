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

func GetShape(s string) Shaper {
	switch s {
	case "X", "A":
		return Rock(1)
	case "Y", "B":
		return Paper(2)
	case "Z", "C":
		return Scissors(3)
	}

	return nil
}
