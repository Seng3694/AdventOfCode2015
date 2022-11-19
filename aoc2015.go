package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func print_day(day int, part1 string, part2 string) {
	fmt.Printf("Solution for day %d:\n  Part1: '%s'\n  Part2: '%s'\n", day, part1, part2)
}

func main() {
	days := []func() (string, string){
		day1,
		day2,
		day3,
		day4,
		day5,
		day6,
		day7,
		day8,
		day9,
		day10,
		day11,
		day12,
		day13,
		day14,
		day15,
		day16,
		day17,
		day18,
		day19,
		day20,
		day21,
		day22,
	}

	start := time.Now()

	if args := os.Args[1:]; len(args) > 0 {
		for _, arg := range args {
			if day, err := strconv.Atoi(arg); err != nil {
				fmt.Printf("Can't convert argument to int '%s'\n", arg)
				return
			} else if day < 1 || day > 25 {
				fmt.Printf("Invalid day argument [1-25]: '%d'\n", day)
				return
			} else if day > len(days) {
				fmt.Printf("Day not solved yet: '%d'\n", day)
			} else {
				pt1, pt2 := days[day-1]()
				print_day(day, pt1, pt2)
			}
		}
	} else {
		pt1, pt2 := days[len(days)-1]()
		print_day(len(days), pt1, pt2)
	}

	elapsed := time.Since(start)
	fmt.Printf("finished in %v seconds\n", elapsed.Seconds())
}
