package main

import (
	"aocutil"
	"math"
	"strconv"
)

func sum_group(group *[]int) (sum int) {
	for _, w := range *group {
		sum += w
	}
	return
}

func prod_group(group *[]int) (prod int) {
	prod = 1
	for _, w := range *group {
		prod *= w
	}
	return
}

func solution(groupCount int) int {
	weights := make([]int, 0, 30)

	aocutil.FileReadAllLines("day24/input.txt", func(s string) {
		weight, _ := strconv.Atoi(s)
		weights = append(weights, weight)
	})

	weightSum := sum_group(&weights)
	groupWeight := weightSum / groupCount

	weightLength := uint32(len(weights))

	group := make([]int, 0, weightLength)
	minQuantumEntanglement := math.MaxInt
	minGroupSize := uint32(0)
	sum := 0
	for sum < groupWeight {
		//they are sorted in ascending order so go in reverse order
		sum += weights[weightLength-1-minGroupSize]
		minGroupSize++
	}
	minGroupSize++

	//assumption: other groups can still be built when the first one was built successfully
	//with that assumption it's only necessary to go through all combinations for a single group
	//go through all combinations in bits
	combinations := uint32(math.Pow(2, float64(weightLength)))
	for i := uint32(0); i < combinations; i++ {
		//if there are more bits than the minGroupSize it can't be the solution
		if aocutil.BitCount32(i) > minGroupSize {
			continue
		}

		//each bit corresponds to an index in the weights list
		for j := uint32(0); j < weightLength; j++ {
			if aocutil.CheckBit(i, j) {
				group = append(group, weights[j])
			}
		}

		if sum_group(&group) == groupWeight && uint32(len(group)) <= minGroupSize {
			minGroupSize = uint32(len(group))
			if prod := prod_group(&group); prod < minQuantumEntanglement {
				minQuantumEntanglement = prod
			}
		}

		//clear
		group = group[:0]
	}

	return minQuantumEntanglement
}

func main() {
	aocutil.AOCFinish(solution(3), solution(4))
}
