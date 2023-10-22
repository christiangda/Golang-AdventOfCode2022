# AdventOfCode2022

Golang Implementation of Advent of Code 2022 [https://adventofcode.com/2022](https://adventofcode.com/2022)

## Usage

Having [Go (Golang)](https://go.dev/) installed, and this repository cloned, you can run the following commands to execute the solutions.

### Unit Tests

All the puzzles have their own unit tests, so you can run them with the following command:

```bash
go test ./...
```

### Day 01

#### Puzzle 01

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_01/puzzle_01/main.go -input day_01/puzzle_01/input.txt
```

Output:

```text
Total Elves: 266
Total Calories: 11900800
Elf with most Calories: Elf 6, Calories: 69281
```

#### Puzzle 02

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_01/puzzle_02/main.go -input day_01/puzzle_02/input.txt
```

Output:

```text
Total Elves: 266
Total Calories: 11900800

Top Three Elf most Calories:
-> Name: Elf 6, Calories: 69281
-> Name: Elf 154, Calories: 67653
-> Name: Elf 180, Calories: 64590

Total carrying calories (Top Three Elf): 201524
```

### Day 02

#### Puzzle 01

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_02/puzzle_01/main.go -input day_02/puzzle_01/input.txt

# or with debug mode
go run day_02/puzzle_01/main.go -input day_02/puzzle_01/input.txt -debug
```

Output:

```text
Number of rounds: 2500
Total score according to my strategy: 11666
```

#### Puzzle 02

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_02/puzzle_02/main.go -input day_02/puzzle_02/input.txt

# or with debug mode
go run day_02/puzzle_02/main.go -input day_02/puzzle_02/input.txt -debug
```

Output:

```text
Number of rounds: 2500
Total score according to my strategy: 12767
```
