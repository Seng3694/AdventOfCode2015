package main

import (
	"aocutil"
)

func say_convert(numberStr string) string {
	output := make([]byte, 0, len(numberStr))
	var count byte = 0

	for i := 0; i < len(numberStr); i++ {
		count++
		//interrupt sequence when at end or the next character is different
		if i == len(numberStr)-1 || numberStr[i] != numberStr[i+1] {
			output = append(output, count+'0')
			output = append(output, numberStr[i])
			count = 0
		}
	}
	return string(output)
}

func main() {
	input := "3113322113"
	solution1 := 0
	for i := 0; i < 50; i++ {
		input = say_convert(input)
		if i == 39 {
			solution1 = len(input)
		}
	}
	aocutil.AOCFinish(solution1, len(input))
}
