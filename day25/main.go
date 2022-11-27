package main

import (
	"aocutil"
)

func main() {
	row := 2978 - 1
	column := 3083 - 1
	diagonals := row + column
	indices := ((diagonals*diagonals + diagonals) / 2) + column
	previous := 20151125
	code := 0

	for i := 0; i < indices; i++ {
		code = (previous * 252533) % 33554393
		previous = code
	}

	aocutil.AOCFinish(code)
}
