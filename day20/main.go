package main

import (
	"aocutil"
	"math"
)

func part1(input int) int {
	presentsPerElf := 10
	presents := 0
	house := int(math.Sqrt(float64(input))) - 1
	for presents < input {
		house++
		presents = 0
		max := int(math.Sqrt(float64(house)))
		for i := 1; i <= max; i++ {
			if house%i == 0 {
				//house = 100, i = 5 (add for presents for both 5 and 20)
				factor := house / i
				if factor != i {
					presents += (factor * presentsPerElf)
				}
				presents += (i * presentsPerElf)
			}
		}
	}
	return house
}

func part2(input int) int {
	presentsPerElf := 11
	presents := 0
	house := int(math.Sqrt(float64(input))) - 1
	elves := make(map[int]int, 300)
	for presents < input {
		house++
		presents = 0
		max := int(math.Sqrt(float64(house)))
		for i := 1; i <= max; i++ {
			if house%i == 0 {
				factor := house / i
				if factor != i {
					if _, ok := elves[factor]; !ok {
						elves[factor] = 0
					}
					if elves[factor] < 50 {
						elves[factor]++
						presents += (factor * presentsPerElf)
					}
				}
				if _, ok := elves[i]; !ok {
					elves[i] = 0
				}
				if elves[i] < 50 {
					elves[i]++
					presents += (i * presentsPerElf)
				}
			}
		}
	}
	return house
}

func main() {
	input := 34000000
	aocutil.AOCFinish(part1(input), part2(input))
}
