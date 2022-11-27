package main

import (
	"aocutil"
	"regexp"
	"sort"
	"strings"
)

/*
--- Day 19: Medicine for Rudolph ---

Rudolph the Red-Nosed Reindeer is sick! His nose isn't shining very brightly, and he needs medicine.

Red-Nosed Reindeer biology isn't similar to regular reindeer biology; Rudolph is going to need custom-made medicine. Unfortunately, Red-Nosed Reindeer chemistry isn't similar to regular reindeer chemistry, either.

The North Pole is equipped with a Red-Nosed Reindeer nuclear fusion/fission plant, capable of constructing any Red-Nosed Reindeer molecule you need. It works by starting with some input molecule and then doing a series of replacements, one per step, until it has the right molecule.

However, the machine has to be calibrated before it can be used. Calibration involves determining the number of molecules that can be generated in one step from a given starting point.

For example, imagine a simpler machine that supports only the following replacements:

H => HO
H => OH
O => HH

Given the replacements above and starting with HOH, the following molecules could be generated:

    HOOH (via H => HO on the first H).
    HOHO (via H => HO on the second H).
    OHOH (via H => OH on the first H).
    HOOH (via H => OH on the second H).
    HHHH (via O => HH).

So, in the example above, there are 4 distinct molecules (not five, because HOOH appears twice) after one replacement from HOH. Santa's favorite molecule, HOHOHO, can become 7 distinct molecules (over nine replacements: six from H, and three from O).

The machine replaces without regard for the surrounding characters. For example, given the string H2O, the transition H => OO would result in OO2O.

Your puzzle input describes all of the possible replacements and, at the bottom, the medicine molecule for which you need to calibrate the machine. How many distinct molecules can be created after all the different ways you can do one replacement on the medicine molecule?

--- Part Two ---

Now that the machine is calibrated, you're ready to begin molecule fabrication.

Molecule fabrication always begins with just a single electron, e, and applying replacements one at a time, just like the ones during calibration.

For example, suppose you have the following replacements:

e => H
e => O
H => HO
H => OH
O => HH

If you'd like to make HOH, you start with e, and then make the following replacements:

    e => O to get O
    O => HH to get HH
    H => OH (on the second H) to get HOH

So, you could make HOH after 3 steps. Santa's favorite molecule, HOHOHO, can be made in 6 steps.

How long will it take to make the medicine? Given the available replacements and the medicine molecule in your puzzle input, what is the fewest number of steps to go from e to the medicine molecule?


*/

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
