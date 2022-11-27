package main

import (
	"aocutil"
	"regexp"
	"sort"
	"strings"
)

type Day19Rule struct {
	search      string
	replacement string
}

func main() {
	molecule := ""
	rules := make([]Day19Rule, 0, 50)

	aocutil.FileReadAllLines("day19/input.txt", func(s string) {
		splits := strings.Split(s, " => ")
		if len(splits) == 2 { //rules "search => replace"
			rules = append(rules, Day19Rule{search: splits[0], replacement: splits[1]})
		} else if len(s) > 0 {
			molecule = s
		}
	})

	//sort rules in ascending replacement length order (for part 2)
	sort.Slice(rules, func(i, j int) bool {
		return len(rules[i].replacement) > len(rules[j].replacement)
	})

	//precompile regexes
	searchRegexes := make([]*regexp.Regexp, len(rules))
	replaceRegexes := make([]*regexp.Regexp, len(rules))
	for i, r := range rules {
		searchRegexes[i] = regexp.MustCompile(r.search)
		replaceRegexes[i] = regexp.MustCompile(r.replacement)
	}

	//part 1
	set := make(map[string]bool)
	copy := molecule
	for i, r := range rules {
		for _, i := range searchRegexes[i].FindAllStringIndex(copy, -1) {
			copy := copy[:i[0]] + r.replacement + copy[i[1]:]
			if _, ok := set[copy]; !ok {
				set[copy] = true
			}
		}
	}

	//part 2
	copy = molecule
	steps := 0
	for {
		for i, r := range rules {
			matches := replaceRegexes[i].FindAllStringIndex(copy, -1)
			//replace matches starting from the last one
			for j := len(matches) - 1; j >= 0; j-- {
				if j < 0 {
					break
				}
				copy = copy[:matches[j][0]] + r.search + copy[matches[j][1]:]
				steps++
			}

			//if something was replaced make another step starting with the biggest rule again
			if len(matches) > 0 {
				break
			}
		}

		if len(copy) == 1 && copy[0] == 'e' {
			break
		}
	}

	aocutil.AOCFinish(len(set), steps)
}
