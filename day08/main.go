package main

import (
	"aocutil"
)

func day8_count_string_memory(str string) (visualLength, memoryLength, encodedLength int) {
	visualLength = len(str)
	encodedLength = 6

	for i := 1; i < len(str)-1; i++ {
		switch str[i] {
		case '\\':
			if i+3 < visualLength && str[i+1] == 'x' {
				if aocutil.IsHexChar(str[i+2]) && aocutil.IsHexChar(str[i+3]) {
					i += 3
					encodedLength += 3
				}
			} else if i+1 < visualLength {
				if str[i+1] == '\\' || str[i+1] == '"' {
					i++
					encodedLength += 2
				}
			}
			encodedLength += 2
			memoryLength++
		default:
			encodedLength++
			memoryLength++
		}
	}

	return
}

func main() {
	visualTotal := 0
	memoryTotal := 0
	encodedTotal := 0

	aocutil.FileReadAllLines("day08/input.txt", func(s string) {
		visual, memory, encoded := day8_count_string_memory(s)
		visualTotal += visual
		memoryTotal += memory
		encodedTotal += encoded
	})

	aocutil.AOCFinish(visualTotal-memoryTotal, encodedTotal-visualTotal)
}
