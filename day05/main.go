package main

import (
	"aocutil"
	"strings"
)

func is_nice_part1(str string) bool {
	var (
		vowelCount      int
		hasDoubleLetter bool
		previousLetter  rune
		vowels          string = "aeiou"
	)

	badStrings := []string{
		"ab", "cd", "pq", "xy",
	}

	//check bad strings
	for _, s := range badStrings {
		if strings.Contains(str, s) {
			return false
		}
	}

	for _, c := range str {
		if strings.ContainsRune(vowels, c) {
			vowelCount++
		}
		if c == previousLetter {
			hasDoubleLetter = true
		}
		previousLetter = c
	}

	return vowelCount >= 3 && hasDoubleLetter
}

func is_nice_part2(str string) bool {
	var (
		hasDoubleLetterPair bool
		hasRepeatingLetter  bool
	)

	for i := 0; i < len(str); i++ {
		if !hasDoubleLetterPair && i < len(str)-1 {
			if strings.Count(str, str[i:i+2]) >= 2 {
				hasDoubleLetterPair = true
			}
		}

		if i >= 2 && str[i] == str[i-2] {
			hasRepeatingLetter = true
		}

	}

	return hasDoubleLetterPair && hasRepeatingLetter
}

func main() {
	niceStringCount1 := 0
	niceStringCount2 := 0

	aocutil.FileReadAllLines("day05/input.txt", func(s string) {
		if is_nice_part1(s) {
			niceStringCount1++
		}
		if is_nice_part2(s) {
			niceStringCount2++
		}
	})

	aocutil.AOCFinish(niceStringCount1, niceStringCount2)
}
