package main

import (
	"aocutil"
	"math"
	"strconv"
	"strings"
)

type Person struct {
	id        int
	happiness map[int]int
}

var (
	currentPerson  int
	persons        map[string]Person = make(map[string]Person)
	personsMapping map[int]string    = make(map[int]string)
)

func add_person_if_not_exist(person string) {
	if _, ok := persons[person]; !ok {
		persons[person] = Person{id: currentPerson, happiness: make(map[int]int)}
		personsMapping[currentPerson] = person
		currentPerson++
	}
}

func add_person_rule(person, partner string, happiness int) {
	add_person_if_not_exist(person)
	add_person_if_not_exist(partner)
	persons[person].happiness[persons[partner].id] = happiness
}

func parse_line(line string) {
	split := strings.Split(line[:len(line)-1], " ")
	value, _ := strconv.Atoi(split[3])
	if split[2] == "lose" {
		value *= (-1)
	}

	add_person_rule(split[0], split[10], value)
}

func wrap(value, min, max int) int {
	if value > max {
		return min
	} else if value < min {
		return max
	} else {
		return value
	}
}

func permutations(currentIndex int, highestHappiness *int, personsIndices []int) {
	if currentIndex == len(persons) {
		happiness := 0
		// for each person check the person left and right
		for i, v := range personsIndices {
			leftIndex := wrap(i-1, 0, len(personsIndices)-1)
			rightIndex := wrap(i+1, 0, len(personsIndices)-1)

			currentPerson := persons[personsMapping[v]]
			leftPerson := persons[personsMapping[personsIndices[leftIndex]]]
			rightPerson := persons[personsMapping[personsIndices[rightIndex]]]

			happiness += currentPerson.happiness[leftPerson.id]
			happiness += currentPerson.happiness[rightPerson.id]
		}

		if happiness > *highestHappiness {
			*highestHappiness = happiness
		}
		return
	}
	for i := 0; i < len(personsIndices); i++ {
		if currentIndex > 0 {
			doContinue := false
			for j := 0; j < currentIndex; j++ {
				if personsIndices[j] == i {
					doContinue = true
					break
				}
			}

			if doContinue {
				continue
			}
		}

		personsIndices[currentIndex] = i
		permutations(currentIndex+1, highestHappiness, personsIndices)
	}
}

func main() {
	aocutil.FileReadAllLines("day13/input.txt", func(s string) {
		parse_line(s)
	})

	happinessPart1 := math.MinInt
	indices := make([]int, len(persons))

	permutations(0, &happinessPart1, indices)

	happinessPart2 := math.MinInt
	names := make([]string, 0, len(persons))
	for k := range persons {
		names = append(names, k)
	}
	for _, v := range names {
		add_person_rule("Me", v, 0)
		add_person_rule(v, "Me", 0)
	}
	indices = make([]int, len(persons))
	permutations(0, &happinessPart2, indices)

	aocutil.AOCFinish(happinessPart1, happinessPart2)
}
