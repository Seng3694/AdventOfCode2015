package main

import (
	"aocutil"
	"math"
	"strconv"
	"strings"
)

/*
--- Day 9: All in a Single Night ---

Every year, Santa manages to deliver all of his presents in a single night.

This year, however, he has some new locations to visit; his elves have provided him the distances between every pair of locations. He can start and end at any two (different) locations he wants, but he must visit each location exactly once. What is the shortest distance he can travel to achieve this?

For example, given the following distances:

London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141

The possible routes are therefore:

Dublin -> London -> Belfast = 982
London -> Dublin -> Belfast = 605
London -> Belfast -> Dublin = 659
Dublin -> Belfast -> London = 659
Belfast -> Dublin -> London = 605
Belfast -> London -> Dublin = 982

The shortest of these is London -> Dublin -> Belfast = 605, and so the answer is 605 in this example.

What is the distance of the shortest route?

--- Part Two ---

The next year, just to show off, Santa decides to take the route with the longest distance instead.

He can still start and end at any two (different) locations he wants, and he still must visit each location exactly once.

For example, given the distances above, the longest route would be 982 via (for example) Dublin -> London -> Belfast.

What is the distance of the longest route?


*/

type Location struct {
	name  string
	paths []Path
}

type Path struct {
	to       *Location
	distance int
}

func parse_line(str string, locationMap map[string]*Location) {
	arr := strings.Split(str, " ")
	from := arr[0]
	to := arr[2]
	dist, _ := strconv.Atoi(arr[4])

	var (
		loc1 *Location
		loc2 *Location
	)

	if v, hasKey := locationMap[from]; !hasKey {
		loc1 = &Location{name: from, paths: make([]Path, 0)}
		locationMap[from] = loc1
	} else {
		loc1 = v
	}

	if v, hasKey := locationMap[to]; !hasKey {
		loc2 = &Location{name: to, paths: make([]Path, 0)}
		locationMap[to] = loc2
	} else {
		loc2 = v
	}

	loc1.paths = append(loc1.paths, Path{distance: dist, to: loc2})
	loc2.paths = append(loc2.paths, Path{distance: dist, to: loc1})
}

func permutations(currentIndex int, smallestDistance *int, longestDistance *int, locations []string, indices []int, locationMap map[string]*Location) {
	if currentIndex == len(locations) {

		distance := 0
		for i := 0; i < len(indices)-1; i++ {
			from := locationMap[locations[indices[i]]]
			to := locationMap[locations[indices[i+1]]]

			for _, p := range from.paths {
				if p.to.name == to.name {
					distance += p.distance
					break
				}
			}
		}

		if distance < *smallestDistance {
			*smallestDistance = distance
		} else if distance > *longestDistance {
			*longestDistance = distance
		}

		return
	}

	for i := 0; i < len(locations); i++ {
		if currentIndex > 0 {
			doContinue := false
			for j := 0; j < currentIndex; j++ {
				if indices[j] == i {
					doContinue = true
					break
				}
			}

			if doContinue {
				continue
			}
		}

		indices[currentIndex] = i
		permutations(currentIndex+1, smallestDistance, longestDistance, locations, indices, locationMap)
	}
}

func main() {
	locationMap := make(map[string]*Location)

	aocutil.FileReadAllLines("day9/input.txt", func(s string) {
		parse_line(s, locationMap)
	})

	//assumption all nodes are connected with all other nodes
	locations := make([]string, 0, len(locationMap))
	for k := range locationMap {
		locations = append(locations, k)
	}

	shortestdistance := math.MaxInt
	longestDistance := 0
	indices := make([]int, len(locations))
	permutations(0, &shortestdistance, &longestDistance, locations, indices, locationMap)

	aocutil.AOCFinish(shortestdistance, longestDistance)
}
