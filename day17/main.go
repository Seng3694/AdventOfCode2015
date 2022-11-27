package main

import (
	"aocutil"
	"math"
	"strconv"
)

func main() {
	barrels := make([]int, 0, 25)
	aocutil.FileReadAllLines("day17/input.txt", func(s string) {
		v, _ := strconv.Atoi(s)
		barrels = append(barrels, v)
	})

	countPart1 := 0
	countPart2 := 0
	smallest := uint32(math.MaxUint32)
	barrelCount := uint32(len(barrels))
	//each bit represents one barrel. so 2^count combinations
	combinations := uint32(math.Pow(2, float64(barrelCount)))
	for i := uint32(0); i < combinations; i++ {
		sum := 0
		for j := uint32(0); j < barrelCount; j++ {
			if aocutil.CheckBit(i, j) {
				sum += barrels[j]
			}
		}

		if sum == 150 {
			//for part 2 check how many barrels are used
			//each bit is one barrel
			bits := aocutil.BitCount32(i)
			if bits < smallest {
				//if a smaller amount is found reset the count
				smallest = bits
				countPart2 = 0
			}

			countPart1++
			if bits == smallest {
				countPart2++
			}
		}
	}

	aocutil.AOCFinish(countPart1, countPart2)
}
