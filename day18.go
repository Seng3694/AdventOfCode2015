package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
--- Day 18: Like a GIF For Your Yard ---

After the million lights incident, the fire code has gotten stricter: now, at most ten thousand lights are allowed. You arrange them in a 100x100 grid.

Never one to let you down, Santa again mails you instructions on the ideal lighting configuration. With so few lights, he says, you'll have to resort to animation.

Start by setting your lights to the included initial configuration (your puzzle input). A # means "on", and a . means "off".

Then, animate your grid in steps, where each step decides the next configuration based on the current one. Each light's next state (either on or off) depends on its current state and the current states of the eight lights adjacent to it (including diagonals). Lights on the edge of the grid might have fewer than eight neighbors; the missing ones always count as "off".

For example, in a simplified 6x6 grid, the light marked A has the neighbors numbered 1 through 8, and the light marked B, which is on an edge, only has the neighbors marked 1 through 5:

1B5...
234...
......
..123.
..8A4.
..765.

The state a light should have next is based on its current state (on or off) plus the number of neighbors that are on:

    A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
    A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.

All of the lights update simultaneously; they all consider the same current state before moving to the next.

Here's a few steps from an example configuration of another 6x6 grid:

Initial state:
.#.#.#
...##.
#....#
..#...
#.#..#
####..

After 1 step:
..##..
..##.#
...##.
......
#.....
#.##..

After 2 steps:
..###.
......
..###.
......
.#....
.#....

After 3 steps:
...#..
......
...#..
..##..
......
......

After 4 steps:
......
......
..##..
..##..
......
......

After 4 steps, this example has four lights on.

In your grid of 100x100 lights, given your initial configuration, how many lights are on after 100 steps?

--- Part Two ---

You flip the instructions over; Santa goes on to point out that this is all just an implementation of Conway's Game of Life. At least, it was, until you notice that something's wrong with the grid of lights you bought: four lights, one in each corner, are stuck on and can't be turned off. The example above will actually run like this:

Initial state:
##.#.#
...##.
#....#
..#...
#.#..#
####.#

After 1 step:
#.##.#
####.#
...##.
......
#...#.
#.####

After 2 steps:
#..#.#
#....#
.#.##.
...##.
.#..##
##.###

After 3 steps:
#...##
####.#
..##.#
......
##....
####.#

After 4 steps:
#.####
#....#
...#..
.##...
#.....
#.#..#

After 5 steps:
##.###
.##..#
.##...
.##...
#.#...
##...#

After 5 steps, this example now has 17 lights on.

In your grid of 100x100 lights, given your initial configuration, but with the four corners always in the on state, how many lights are on after 100 steps?
*/

const (
	LIGHT_STATE_OFF byte = '.'
	LIGHT_STATE_ON  byte = '#'
)

func day18_update_lights(lights []byte, buffer []byte, width, height int, cornerAlwaysOn bool) {
	copy(buffer, lights)

	topLeft := 0
	topRight := width - 1
	bottomRight := (width-1)*width + (width - 1)
	bottomLeft := (width - 1) * width

	for i := 0; i < len(buffer); i++ {
		lightX := (i % width)
		lightY := (i / width)

		if cornerAlwaysOn {
			if i == topLeft || i == topRight || i == bottomRight || i == bottomLeft {
				continue
			}
		}

		/*
			go through all neighbors
			0 1 2
			3 4 5
			6 7 8
			with 4 being the light (skip)
		*/
		litNeighbors := 0
		for j := 0; j < 9; j++ {
			if j == 4 {
				continue
			}
			//standard offset: x-1, y-1
			offsetX := -1
			offsetY := -1
			neighborX := lightX + offsetX + (j % 3)
			neighborY := lightY + offsetY + (j / 3)
			neighborIndex := neighborY*width + neighborX
			if neighborX >= 0 && neighborX < width && neighborY >= 0 && neighborY < height {
				if lights[neighborIndex] == '#' {
					litNeighbors++
				}
			}
		}

		switch lights[i] {
		case LIGHT_STATE_ON:
			if litNeighbors < 2 || litNeighbors > 3 {
				buffer[i] = LIGHT_STATE_OFF
			}
		case LIGHT_STATE_OFF:
			if litNeighbors == 3 {
				buffer[i] = LIGHT_STATE_ON
			}
		}
	}
	copy(lights, buffer)
}

func day18_count_lights(lights []byte) (count int) {
	for _, state := range lights {
		if state == LIGHT_STATE_ON {
			count++
		}
	}
	return
}

func day18_print_lights(lights []byte, width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Printf("%c", lights[y*width+x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func day18() (string, string) {
	file, err := os.Open("input/day18.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	width, height := 100, 100
	lights := make([]byte, 0, width*height)
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			lights = append(lights, byte(c))
		}
	}

	steps := 100
	lightsCopy := make([]byte, len(lights))
	copy(lightsCopy, lights)
	buffer := make([]byte, len(lightsCopy))
	for i := 0; i < steps; i++ {
		day18_update_lights(lightsCopy, buffer, width, height, false)
	}
	part1 := day18_count_lights(lightsCopy)

	//reset lights
	copy(lightsCopy, lights)
	//turn on corner lights
	topLeft := 0
	topRight := width - 1
	bottomRight := (width-1)*width + (width - 1)
	bottomLeft := (width - 1) * width
	lightsCopy[topLeft] = LIGHT_STATE_ON
	lightsCopy[topRight] = LIGHT_STATE_ON
	lightsCopy[bottomRight] = LIGHT_STATE_ON
	lightsCopy[bottomLeft] = LIGHT_STATE_ON
	for i := 0; i < steps; i++ {
		day18_update_lights(lightsCopy, buffer, width, height, true)
	}
	part2 := day18_count_lights(lightsCopy)

	return fmt.Sprint(part1), fmt.Sprint(part2)
}
