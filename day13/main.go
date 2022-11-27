package main

import (
	"aocutil"
	"math"
	"strconv"
	"strings"
)

/*
--- Day 13: Knights of the Dinner Table ---

In years past, the holiday feast with your family hasn't gone so well. Not everyone gets along! This year, you resolve, will be different. You're going to find the optimal seating arrangement and avoid all those awkward conversations.

You start by writing up a list of everyone invited and the amount their happiness would increase or decrease if they were to find themselves sitting next to each other person. You have a circular table that will be just big enough to fit everyone comfortably, and so each person will have exactly two neighbors.

For example, suppose you have only four attendees planned, and you calculate their potential happiness as follows:

Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.

Then, if you seat Alice next to David, Alice would lose 2 happiness units (because David talks so much), but David would gain 46 happiness units (because Alice is such a good listener), for a total change of 44.

If you continue around the table, you could then seat Bob next to Alice (Bob gains 83, Alice gains 54). Finally, seat Carol, who sits next to Bob (Carol gains 60, Bob loses 7) and David (Carol gains 55, David gains 41). The arrangement looks like this:

     +41 +46
+55   David    -2
Carol       Alice
+60    Bob    +54
     -7  +83

After trying every other seating arrangement in this hypothetical scenario, you find that this one is the most optimal, with a total change in happiness of 330.

What is the total change in happiness for the optimal seating arrangement of the actual guest list?

--- Part Two ---

In all the commotion, you realize that you forgot to seat yourself. At this point, you're pretty apathetic toward the whole thing, and your happiness wouldn't really go up or down regardless of who you sit next to. You assume everyone else would be just as ambivalent about sitting next to you, too.

So, add yourself to the list, and give all happiness relationships that involve you a score of 0.

What is the total change in happiness for the optimal seating arrangement that actually includes yourself?


*/

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
