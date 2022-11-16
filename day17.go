package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func check_bit(value uint32, index int) bool {
	return (value & (1 << index)) == (1 << index)
}

// https://stackoverflow.com/questions/109023/count-the-number-of-set-bits-in-a-32-bit-integer
// https://en.wikipedia.org/wiki/Hamming_weight
func popcount32(i uint32) int {
	i = i - ((i >> 1) & 0x55555555)                // add pairs of bits
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333) // quads
	i = (i + (i >> 4)) & 0x0F0F0F0F                // groups of 8
	return int((i * 0x01010101) >> 24)             // horizontal sum of bytes
}

func day17() (string, string) {
	file, err := os.Open("input/day17.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	barrels := make([]int, 0, 25)
	for scanner.Scan() {
		line := scanner.Text()
		v, _ := strconv.Atoi(line)
		barrels = append(barrels, v)
	}

	countPart1 := 0
	countPart2 := 0
	smallest := math.MaxInt
	//each bit represents one barrel. so 2^count combinations
	combinations := uint32(math.Pow(2, float64(len(barrels))))
	for i := uint32(0); i < combinations; i++ {
		sum := 0
		for j := 0; j < len(barrels); j++ {
			if check_bit(i, j) {
				sum += barrels[j]
			}
		}

		if sum == 150 {
			//for part 2 check how many barrels are used
			//each bit is one barrel
			bits := popcount32(i)
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

	return fmt.Sprint(countPart1), fmt.Sprint(countPart2)
}
