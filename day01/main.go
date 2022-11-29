package main

import (
	"aocutil"
)

func part1(input []byte) (floor int) {
	//divide work and assign each to a goroutine
	const parts int = 4
	part := len(input) / parts
	channel := make(chan int, parts)
	for i := 0; i < parts; i++ {
		go func(from, to int, out chan int) {
			f := 0
			for i := from; i < to; i++ {
				if input[i] == '(' {
					f++
				} else {
					f--
				}
			}
			out <- f
		}(i*part, (i+1)*part, channel)
	}
	for i := 0; i < parts; i++ {
		floor += <-channel
	}
	return
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
