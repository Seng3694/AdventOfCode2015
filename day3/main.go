package main

import (
	"aocutil"
	"fmt"
)

type Vector struct {
	x, y int
}

func (v *Vector) Move(c byte) {
	switch c {
	case '<':
		v.x--
	case '>':
		v.x++
	case '^':
		v.y--
	case 'v':
		v.y++
	default:
		fmt.Printf("Unexpected input: '%v'\n", c)
	}
}

func day3_solution(santaCount int, data []byte) int {
	positions := make([]Vector, santaCount)
	visits := make(map[Vector]int)
	visits[Vector{0, 0}] = santaCount

	for i, c := range data {
		pos := &positions[i%santaCount]
		pos.Move(c)
		visits[*pos]++
	}

	return len(visits)
}

func main() {
	input := aocutil.FileReadAll[[]byte]("day3/input.txt")
	aocutil.AOCFinish(day3_solution(1, input), day3_solution(2, input))
}
