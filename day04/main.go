package main

import (
	"aocutil"
	"fmt"
	"strings"
)

func main() {
	data := aocutil.FileReadAll[[]byte]("day04/input.txt")

	i := 0
	part1Solution := 0
	part2Solution := 0
	part1SolutionFound := false

	for {
		hash := aocutil.GetMD5Hash(fmt.Sprintf("%s%d", data, i))

		if part1SolutionFound {
			if strings.HasPrefix(hash, "000000") {
				part2Solution = i
				break
			}
		} else {
			if strings.HasPrefix(hash, "00000") {
				part1Solution = i
				part1SolutionFound = true
			}
		}
		i++
	}

	aocutil.AOCFinish(part1Solution, part2Solution)
}
