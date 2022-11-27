package main

import (
	"aocutil"
	"strconv"
	"strings"
)

/*
--- Day 6: Probably a Fire Hazard ---

Because your neighbors keep defeating you in the holiday house decorating contest year after year, you've decided to deploy one million lights in a 1000x1000 grid.

Furthermore, because you've been especially nice this year, Santa has mailed you instructions on how to display the ideal lighting configuration.

Lights in your grid are numbered from 0 to 999 in each direction; the lights at each corner are at 0,0, 0,999, 999,999, and 999,0. The instructions include whether to turn on, turn off, or toggle various inclusive ranges given as coordinate pairs. Each coordinate pair represents opposite corners of a rectangle, inclusive; a coordinate pair like 0,0 through 2,2 therefore refers to 9 lights in a 3x3 square. The lights all start turned off.

To defeat your neighbors this year, all you have to do is set up your lights by doing the instructions Santa sent you in order.

For example:

    turn on 0,0 through 999,999 would turn on (or leave on) every light.
    toggle 0,0 through 999,0 would toggle the first line of 1000 lights, turning off the ones that were on, and turning on the ones that were off.
    turn off 499,499 through 500,500 would turn off (or leave off) the middle four lights.

After following the instructions, how many lights are lit?

--- Part Two ---

You just finish implementing your winning light pattern when you realize you mistranslated Santa's message from Ancient Nordic Elvish.

The light grid you bought actually has individual brightness controls; each light can have a brightness of zero or more. The lights all start at zero.

The phrase turn on actually means that you should increase the brightness of those lights by 1.

The phrase turn off actually means that you should decrease the brightness of those lights by 1, to a minimum of zero.

The phrase toggle actually means that you should increase the brightness of those lights by 2.

What is the total brightness of all lights combined after following Santa's instructions?

For example:

    turn on 0,0 through 0,0 would increase the total brightness by 1.
    toggle 0,0 through 999,999 would increase the total brightness by 2000000.

*/

type Vector struct {
	x, y int
}

var (
	width, height int = 1000, 1000
)

func do_with_lights(lights []int, start, end Vector, f func(*int)) {
	for y := start.y; y <= end.y; y++ {
		for x := start.x; x <= end.x; x++ {
			f(&lights[y*width+x])
		}
	}
}

func count_enabled_lights(lights []int) (count int) {
	for _, light := range lights {
		if light > 0 {
			count++
		}
	}
	return
}

func calculate_total_brightness(lights []int) (brightness int) {
	for _, light := range lights {
		brightness += light
	}
	return
}

func parse_vector(str string) Vector {
	positions := strings.Split(str, ",")
	x, err := strconv.Atoi(positions[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(positions[1])
	if err != nil {
		panic(err)
	}
	return Vector{x, y}
}

func main() {
	lightsPart1 := make([]int, width*height)
	lightsPart2 := make([]int, width*height)

	aocutil.FileReadAllLines("day6/input.txt", func(s string) {
		s = strings.Replace(s, "turn ", "", 1)
		s = strings.Replace(s, "through ", "", 1)
		tokens := strings.Split(s, " ")

		start := parse_vector(tokens[1])
		end := parse_vector(tokens[2])

		switch tokens[0] {
		case "on":
			do_with_lights(lightsPart1, start, end, func(b *int) {
				(*b) = 1
			})
			do_with_lights(lightsPart2, start, end, func(b *int) {
				(*b)++
			})
		case "off":
			do_with_lights(lightsPart1, start, end, func(b *int) {
				(*b) = 0
			})
			do_with_lights(lightsPart2, start, end, func(b *int) {
				if (*b) > 0 {
					(*b)--
				}
			})
		case "toggle":
			do_with_lights(lightsPart1, start, end, func(b *int) {
				*b = (*b) ^ 1
			})
			do_with_lights(lightsPart2, start, end, func(b *int) {
				(*b) += 2
			})
		}
	})

	aocutil.AOCFinish(count_enabled_lights(lightsPart1), calculate_total_brightness(lightsPart2))
}
