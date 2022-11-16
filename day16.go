package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
--- Day 16: Aunt Sue ---

Your Aunt Sue has given you a wonderful gift, and you'd like to send her a thank you card. However, there's a small problem: she signed it "From, Aunt Sue".

You have 500 Aunts named "Sue".

So, to avoid sending the card to the wrong person, you need to figure out which Aunt Sue (which you conveniently number 1 to 500, for sanity) gave you the gift. You open the present and, as luck would have it, good ol' Aunt Sue got you a My First Crime Scene Analysis Machine! Just what you wanted. Or needed, as the case may be.

The My First Crime Scene Analysis Machine (MFCSAM for short) can detect a few specific compounds in a given sample, as well as how many distinct kinds of those compounds there are. According to the instructions, these are what the MFCSAM can detect:

    children, by human DNA age analysis.
    cats. It doesn't differentiate individual breeds.
    Several seemingly random breeds of dog: samoyeds, pomeranians, akitas, and vizslas.
    goldfish. No other kinds of fish.
    trees, all in one group.
    cars, presumably by exhaust or gasoline or something.
    perfumes, which is handy, since many of your Aunts Sue wear a few kinds.

In fact, many of your Aunts Sue have many of these. You put the wrapping from the gift into the MFCSAM. It beeps inquisitively at you a few times and then prints out a message on ticker tape:

children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1

You make a list of the things you can remember about each Aunt Sue. Things missing from your list aren't zero - you simply don't remember the value.

What is the number of the Sue that got you the gift?

--- Part Two ---

As you're about to send the thank you note, something in the MFCSAM's instructions catches your eye. Apparently, it has an outdated retroencabulator, and so the output from the machine isn't exact values - some of them indicate ranges.

In particular, the cats and trees readings indicates that there are greater than that many (due to the unpredictable nuclear decay of cat dander and tree pollen), while the pomeranians and goldfish readings indicate that there are fewer than that many (due to the modial interaction of magnetoreluctance).

What is the number of the real Aunt Sue?
*/

const (
	MFCSAM_CHILDREN int = iota
	MFCSAM_CATS
	MFCSAM_SAMOYEDS
	MFCSAM_POMERANIANS
	MFCSAM_AKITAS
	MFCSAM_VIZLAS
	MFCSAM_GOLDFISH
	MFCSAM_TREES
	MFCSAM_CARS
	MFCSAM_PERFUMES
)

func day16_parse_line(line string) []int {
	split := strings.Split(strings.Replace(line, ",", "", -1), " ")
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = -1 //sentinel value
	}

	//skip first 2 entries 'Sue' '123:'
	for i := 2; i < len(split); i += 2 {
		switch split[i] {
		case "children:":
			data[MFCSAM_CHILDREN], _ = strconv.Atoi(split[i+1])
		case "cats:":
			data[MFCSAM_CATS], _ = strconv.Atoi(split[i+1])
		case "samoyeds:":
			data[MFCSAM_SAMOYEDS], _ = strconv.Atoi(split[i+1])
		case "pomeranians:":
			data[MFCSAM_POMERANIANS], _ = strconv.Atoi(split[i+1])
		case "akitas:":
			data[MFCSAM_AKITAS], _ = strconv.Atoi(split[i+1])
		case "vizslas:":
			data[MFCSAM_VIZLAS], _ = strconv.Atoi(split[i+1])
		case "goldfish:":
			data[MFCSAM_GOLDFISH], _ = strconv.Atoi(split[i+1])
		case "trees:":
			data[MFCSAM_TREES], _ = strconv.Atoi(split[i+1])
		case "cars:":
			data[MFCSAM_CARS], _ = strconv.Atoi(split[i+1])
		case "perfumes:":
			data[MFCSAM_PERFUMES], _ = strconv.Atoi(split[i+1])
		}
	}

	return data
}

func day16() (string, string) {
	file, err := os.Open("input/day16.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	aunts := make([][]int, 0, 500)
	for scanner.Scan() {
		line := scanner.Text()
		aunts = append(aunts, day16_parse_line(line))
	}

	data := make([]int, 10)
	data[MFCSAM_CHILDREN] = 3
	data[MFCSAM_CATS] = 7
	data[MFCSAM_SAMOYEDS] = 2
	data[MFCSAM_POMERANIANS] = 3
	data[MFCSAM_AKITAS] = 0
	data[MFCSAM_VIZLAS] = 0
	data[MFCSAM_GOLDFISH] = 5
	data[MFCSAM_TREES] = 3
	data[MFCSAM_CARS] = 2
	data[MFCSAM_PERFUMES] = 1

	maxScore := 0
	solutionPt1 := -1

	for i, aunt := range aunts {
		score := 0
		for j := 0; j < len(aunt); j++ {
			if aunt[j] == data[j] {
				score++
			}
		}
		if score > maxScore {
			maxScore = score
			solutionPt1 = i + 1
		}
	}

	maxScore = 0
	solutionPt2 := -1

	for i, aunt := range aunts {
		score := 0
		for j := 0; j < len(aunt); j++ {
			if aunt[j] == -1 {
				continue
			}
			if j == MFCSAM_CATS || j == MFCSAM_TREES {
				if aunt[j] > data[j] {
					score++
				}
			} else if j == MFCSAM_POMERANIANS || j == MFCSAM_GOLDFISH {
				if aunt[j] < data[j] {
					score++
				}
			} else {
				if aunt[j] == data[j] {
					score++
				}
			}

		}
		if score > maxScore {
			maxScore = score
			solutionPt2 = i + 1
		}
	}

	return fmt.Sprint(solutionPt1), fmt.Sprint(solutionPt2)
}
