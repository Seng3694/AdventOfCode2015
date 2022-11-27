package main

import (
	"aocutil"
)

func part1(input []byte) (floor int) {
	for _, b := range input {
		if b == '(' {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

func part2(input []byte) int {
	floor := 0

	for i, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			return i + 1
		}
	}

	return -1 //unreachable
}

func main() {
	input := aocutil.FileReadAll[[]byte]("day01/input.txt")
	aocutil.AOCFinish(part1(input), part2(input))
}
