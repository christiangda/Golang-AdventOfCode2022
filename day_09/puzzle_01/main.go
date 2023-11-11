package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	input string
	debug bool
)

func init() {
	flag.StringVar(&input, "input", "input.txt", "input: a file path with the content of your https://adventofcode.com/2022/day/8/input")
	flag.BoolVar(&debug, "debug", false, "debug (default: false)")
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Visited(position [2]int, direction string, steps int) [][2]int {
	var ret [][2]int

	// visited the las position
	ret = append(ret, position)

	// rebuild the previous positions based in the steps and the last position
	switch direction {
	case Up:
		for i := 1; i < steps; i++ {
			ret = append(ret, [2]int{position[X], position[Y] - i})
		}
	case Down:
		for i := 1; i < steps; i++ {
			ret = append(ret, [2]int{position[X], position[Y] + i})
		}
	case Right:
		for i := 1; i < steps; i++ {
			ret = append(ret, [2]int{position[X] - i, position[Y]})
		}
	case Left:
		for i := 1; i < steps; i++ {
			ret = append(ret, [2]int{position[X] + i, position[Y]})
		}
	}

	return ret
}

// Position
const (
	X int = iota
	Y
)

// direction
const (
	Right string = "R"
	Left  string = "L"
	Up    string = "U"
	Down  string = "D"
)

type Knot struct {
	Label     string
	position  [2]int
	movements int

	direction string
	steps     int
	visited   map[[2]int]bool
}

func (k *Knot) MoveUp(steps int) {
	k.position[Y] += steps
	k.movements += steps
	k.steps = steps
}

func (k *Knot) MoveDown(steps int) {
	k.position[Y] -= steps
	k.movements += steps
	k.steps = steps
}

func (k *Knot) MoveRight(steps int) {
	k.position[X] += steps
	k.movements += steps
	k.steps = steps
}

func (k *Knot) MoveLeft(steps int) {
	k.position[X] -= steps
	k.movements += steps
	k.steps = steps
}

func (k *Knot) Move(direction string, steps int) {
	k.direction = direction

	switch direction {
	case Up:
		k.MoveUp(steps)
	case Down:
		k.MoveDown(steps)
	case Right:
		k.MoveRight(steps)
	case Left:
		k.MoveLeft(steps)
	default:
		fmt.Printf("unknown direction: %v\n", direction)
	}
}

type Rope struct {
	head *Knot
	tail *Knot
}

func NewRope(Head, Tail *Knot) *Rope {
	return &Rope{
		head: Head,
		tail: Tail,
	}
}

func (r *Rope) Move(direction string, steps int) {
	r.head.Move(direction, steps)

	switch direction {
	case Up:
		// steps moved
		dy := Abs(r.head.position[Y] - r.tail.position[Y])
		if dy >= 1 {
			r.tail.steps = dy - 1
		}
		r.tail.movements += r.tail.steps

		// tail position
		if dy > 1 {
			r.tail.position[X] = r.head.position[X]
			r.tail.position[Y] = r.head.position[Y] - 1
		}

		visited := Visited(r.tail.position, Up, r.tail.steps)
		for _, v := range visited {
			r.tail.visited[v] = true
		}

	case Down:
		// steps moved
		dy := Abs(r.head.position[Y] - r.tail.position[Y])
		if dy >= 1 {
			r.tail.steps = dy - 1
		}
		r.tail.movements += r.tail.steps

		// tail position
		if dy > 1 {
			r.tail.position[X] = r.head.position[X]
			r.tail.position[Y] = r.head.position[Y] - 1
		}

		visited := Visited(r.tail.position, Down, r.tail.steps)
		for _, v := range visited {
			r.tail.visited[v] = true
		}
	case Right:
		// steps moved
		dx := Abs(r.head.position[X] - r.tail.position[X])
		if dx >= 1 {
			r.tail.steps = dx - 1
		}
		r.tail.movements += r.tail.steps

		// tail position
		if dx > 1 {
			r.tail.position[X] = r.head.position[X] - 1
			r.tail.position[Y] = r.head.position[Y]
		}

		visited := Visited(r.tail.position, Right, r.tail.steps)
		for _, v := range visited {
			r.tail.visited[v] = true
		}
	case Left:
		// steps moved
		dx := Abs(r.head.position[X] - r.tail.position[X])
		if dx >= 1 {
			r.tail.steps = dx - 1
		}
		r.tail.movements += r.tail.steps

		// tail position
		if dx > 1 {
			r.tail.position[X] = r.head.position[X] + 1
			r.tail.position[Y] = r.head.position[Y]
		}

		visited := Visited(r.tail.position, Left, r.tail.steps)
		for _, v := range visited {
			r.tail.visited[v] = true
		}
	default:
		fmt.Printf("unknown direction: %v\n", direction)
	}

	r.tail.direction = direction
}

type KnotMap struct {
	pos [][]rune
}

func NewKnotMap() *KnotMap {
	return &KnotMap{
		pos: make([][]rune, 0),
	}
}

func (km *KnotMap) Render(points map[[2]int]bool) {
	// initialize positions
	minX, minY, maxX, maxY := MinXMinYMaxXMaxY(points)

	row := (maxY - minY) + 1
	col := (maxX - minX) + 2

	fmt.Printf("row: %v, col:%v\n", row, col)
	fmt.Printf("minX: %v, minY: %v, maxX:%v, maxY: %v\n", minX, minY, maxX, maxY)

	km.pos = make([][]rune, row)

	for r := 0; r < row; r++ {
		km.pos[r] = make([]rune, 0, col)

		for c := 0; c < col; c++ {
			km.pos[r] = append(km.pos[r], []rune(".")[0])
		}
	}

	// fill the matrix with the map positions
	for p := range points {
		// normalize
		x := p[X] + Abs(minX)
		y := p[Y] + Abs(minY)

		km.pos[y][x] = []rune("#")[0]
	}

	buf := bytes.Buffer{}

	for y := len(km.pos) - 1; y > -1; y-- {
		var line string

		for x := 0; x < len(km.pos[y]); x++ {
			line += string(km.pos[y][x])
		}

		buf.WriteString(line)
		buf.WriteString("\n")
	}

	fmt.Print(buf.String())
}

func MinXMinYMaxXMaxY(points map[[2]int]bool) (int, int, int, int) {
	var maxX int
	var minX int
	var maxY int
	var minY int

	for k := range points {
		// min
		if k[X] < minX {
			minX = k[X]
		}
		if k[Y] < minY {
			minY = k[Y]
		}

		// max
		if k[X] > maxX {
			maxX = k[X]
		}
		if k[Y] > maxY {
			maxY = k[Y]
		}
	}

	return minX, minY, maxX, maxY
}

func main() {
	flag.Parse()

	if input == "" {
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open(input)
	if err != nil {
		fmt.Printf("Error: Looks like the input file is not valid %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	H := Knot{Label: "H", position: [2]int{0, 0}, visited: map[[2]int]bool{{0, 0}: true}}
	T := Knot{Label: "T", position: [2]int{0, 0}, visited: map[[2]int]bool{{0, 0}: true}}
	Rope := NewRope(&H, &T)

	rc := bufio.NewScanner(file)
	for rc.Scan() {
		line := rc.Text()
		instructions := strings.Split(line, " ")
		if len(instructions) != 2 {
			fmt.Printf("Expected 2 elements in the line, got: %v\n", len(instructions))
			continue
		}

		direction := instructions[0]
		steps, err := strconv.Atoi(instructions[1])
		if err != nil {
			fmt.Printf("Error converting %v steps to integer\n", instructions[1])
		}

		Rope.Move(direction, steps)

		if debug {
			fmt.Printf("H -> Dir: %v, Steps: %v, Pos: %v, Mov: %v\n", H.direction, H.steps, H.position, H.movements)
			fmt.Printf("T -> Dir: %v, Steps: %v, Pos: %v, Mov: %v\n\n", T.direction, T.steps, T.position, T.movements)
		}
	}

	fmt.Println("-----------------------------------------------------------------")
	fmt.Printf("Number of movements of Head: %v\n", H.movements)
	fmt.Printf("Last position of Head: %v\n\n", H.position)
	fmt.Printf("Number of movements of Tail: %v\n", T.movements)
	fmt.Printf("Last position of Tail: %v\n", T.position)

	if debug {
		// fmt.Printf("Positions visited by Tail: %+v\n\n", T.visited)

		m := NewKnotMap()
		fmt.Println("-------------------------------------------------------------------------")
		m.Render(T.visited)
		fmt.Println("-------------------------------------------------------------------------")

	}

	fmt.Printf("How many positions does the tail of the rope visit at least once?: %v\n", len(T.visited))
}
