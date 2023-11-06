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

### Day 03

#### Puzzle 01

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_03/puzzle_01/main.go -input day_03/puzzle_01/input.txt

# or with debug mode
go run day_03/puzzle_01/main.go -input day_03/puzzle_01/input.txt -debug
go run day_03/puzzle_01/main.go -input day_03/puzzle_01/input_test.txt -debug
```

Output:

```text
Number of repeated Items: 300
Sum of the priorities: 8185
```

#### Puzzle 02

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_03/puzzle_02/main.go -input day_03/puzzle_02/input.txt

# or with debug mode
go run day_03/puzzle_02/main.go -input day_03/puzzle_02/input.txt -debug
go run day_03/puzzle_02/main.go -input day_03/puzzle_02/input_test.txt -debug
```

Output:

```text
Number of repeated Items: 100
Sum of the priorities: 2817
```

### Day 04

#### Puzzle 01

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_04/puzzle_01/main.go -input day_04/puzzle_01/input.txt

# or with debug mode
go run day_04/puzzle_01/main.go -input day_04/puzzle_01/input.txt -debug
go run day_04/puzzle_01/main.go -input day_04/puzzle_01/input_test.txt -debug
```

Output:

```text
Number of pairs of sections checked: 1000
Number of sections checked: 2000
Number of assignment pairs that one range fully contain the other: 569
```

#### Puzzle 02

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_04/puzzle_02/main.go -input day_04/puzzle_02/input.txt

# or with debug mode
go run day_04/puzzle_02/main.go -input day_04/puzzle_02/input.txt -debug
go run day_04/puzzle_02/main.go -input day_04/puzzle_02/input_test.txt -debug
```

Output:

```text
Number of pairs of sections checked: 1000
Number of sections checked: 2000
Number of pairs that the ranges overlap: 936
```

### Day 05

#### Puzzle 01

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_05/puzzle_01/main.go -input day_05/puzzle_01/input.txt

# or with debug mode
go run day_05/puzzle_01/main.go -input day_05/puzzle_01/input.txt -debug
go run day_05/puzzle_01/main.go -input day_05/puzzle_01/input_test.txt -debug
```

Output:

```text
Number of crates: 56
Number of instructions: 502
Number of stacks: 9
Which crate will end up on top of each stack: FCVRLMVQP
```

#### Puzzle 02

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_05/puzzle_02/main.go -input day_05/puzzle_02/input.txt

# or with debug mode
go run day_05/puzzle_02/main.go -input day_05/puzzle_02/input.txt -debug
go run day_05/puzzle_02/main.go -input day_05/puzzle_02/input_test.txt -debug
```

Output:

```text
Number of crates: 56
Number of instructions: 502
Number of queues: 9
Which crate will end up on top of each queue: RWLWGJGFD
```

### Day 06

#### Puzzle 01

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_06/puzzle_01/main.go -input day_06/puzzle_01/input.txt

# or with debug mode
go run day_06/puzzle_01/main.go -input day_06/puzzle_01/input.txt -debug
go run day_06/puzzle_01/main.go -input day_06/puzzle_01/input_test.txt -debug
```

Output:

```text
Numbers of lines processed: 1
Line number: 1
Characters that need to be processed before the first start-of-packet marker is detected: 1655
```

#### Puzzle 02

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_06/puzzle_02/main.go -input day_06/puzzle_02/input.txt

# or with debug mode
go run day_06/puzzle_02/main.go -input day_06/puzzle_02/input.txt -debug
go run day_06/puzzle_02/main.go -input day_06/puzzle_02/input_test.txt -debug
```

Output:

```text
Numbers of lines processed: 1
Line number: 1
Characters that need to be processed before the first start-of-packet marker is detected: 2665
```

### Day 07

#### Puzzle 01

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_07/puzzle_01/main.go -input day_07/puzzle_01/input.txt

# or with debug mode
go run day_07/puzzle_01/main.go -input day_07/puzzle_01/input.txt -debug
go run day_07/puzzle_01/main.go -input day_07/puzzle_01/input_test.txt -debug
```

Output:

```text
What is the sum of the total sizes of those directories?: 1443806
```

#### Puzzle 02

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_07/puzzle_02/main.go -input day_07/puzzle_02/input.txt

# or with debug mode
go run day_07/puzzle_02/main.go -input day_07/puzzle_02/input.txt -debug
go run day_07/puzzle_02/main.go -input day_07/puzzle_02/input_test.txt -debug
```

Output:

```text
The total disk space available to the filesystem: 70000000
Used in / (root): 40913445
Size Available: 29086555
Update needed space size: 30000000

Needed space to be free (at least): 913445

What is the Node in the tree selected?: lfrctthp
What is the total size of that directory?: 942298
```

### Day 08

#### Puzzle 01

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_08/puzzle_01/main.go -input day_08/puzzle_01/input.txt

# or with debug mode
go run day_08/puzzle_01/main.go -input day_08/puzzle_01/input.txt -debug
go run day_08/puzzle_01/main.go -input day_08/puzzle_01/input_test.txt -debug
```

Output:

```text
Number of rows: 99
Number of cols: 99
Number of trees: 9801

Number of visible trees inside: 1151
Number of visible trees outside: 392

how many trees are visible from outside the grid?: 1543
```

#### Puzzle 02

Remember that each puzzle has its own input file, so you need to specify it as a parameter.

Using my input file:

```bash
go run day_08/puzzle_02/main.go -input day_08/puzzle_02/input.txt

# or with debug mode
go run day_08/puzzle_02/main.go -input day_08/puzzle_02/input.txt -debug
go run day_08/puzzle_02/main.go -input day_08/puzzle_02/input_test.txt -debug
```

Output:

```text
Number of rows: 99
Number of cols: 99
Number of trees: 9801

What is the highest scenic score possible for any tree?: 595080
```
