package main

import (
	"aocutil"
)

const (
	LIGHT_STATE_OFF byte = '.'
	LIGHT_STATE_ON  byte = '#'
)

func update_lights(lights []byte, buffer []byte, width, height int, cornerAlwaysOn bool) {
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

func count_lights(lights []byte) (count int) {
	for _, state := range lights {
		if state == LIGHT_STATE_ON {
			count++
		}
	}
	return
}

func main() {
	width, height := 100, 100
	lights := make([]byte, 0, width*height)
	aocutil.FileReadAllLines("day18/input.txt", func(s string) {
		for _, c := range s {
			lights = append(lights, byte(c))
		}
	})

	steps := 100
	lightsCopy := make([]byte, len(lights))
	copy(lightsCopy, lights)
	buffer := make([]byte, len(lightsCopy))
	for i := 0; i < steps; i++ {
		update_lights(lightsCopy, buffer, width, height, false)
	}
	part1 := count_lights(lightsCopy)

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
		update_lights(lightsCopy, buffer, width, height, true)
	}
	part2 := count_lights(lightsCopy)

	aocutil.AOCFinish(part1, part2)
}
