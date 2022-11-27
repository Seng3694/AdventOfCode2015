package main

import (
	"aocutil"
	"strconv"
	"strings"
)

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

	aocutil.FileReadAllLines("day06/input.txt", func(s string) {
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
