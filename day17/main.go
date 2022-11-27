package main

import (
	"aocutil"
	"math"
	"strconv"
)

/*
--- Day 17: No Such Thing as Too Much ---

The elves bought too much eggnog again - 150 liters this time. To fit it all into your refrigerator, you'll need to move it into smaller containers. You take an inventory of the capacities of the available containers.

For example, suppose you have containers of size 20, 15, 10, 5, and 5 liters. If you need to store 25 liters, there are four ways to do it:

    15 and 10
    20 and 5 (the first 5)
    20 and 5 (the second 5)
    15, 5, and 5

Filling all containers entirely, how many different combinations of containers can exactly fit all 150 liters of eggnog?

--- Part Two ---

While playing with all the containers in the kitchen, another load of eggnog arrives! The shipping and receiving department is requesting as many containers as you can spare.

Find the minimum number of containers that can exactly fit all 150 liters of eggnog. How many different ways can you fill that number of containers and still hold exactly 150 litres?

In the example above, the minimum number of containers was two. There were three ways to use that many containers, and so the answer there would be 3.

*/

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
