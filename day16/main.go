package main

import (
	"aocutil"
	"strconv"
	"strings"
)

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

func parse_line(line string) []int {
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

func main() {
	aunts := make([][]int, 0, 500)
	aocutil.FileReadAllLines("day16/input.txt", func(s string) {
		aunts = append(aunts, parse_line(s))
	})

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

	aocutil.AOCFinish(solutionPt1, solutionPt2)
}
