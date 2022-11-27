package main

import (
	"aocutil"
	"math"
	"strconv"
	"strings"
)

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
