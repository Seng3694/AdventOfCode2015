package main

import (
	"aocutil"
	"encoding/json"
	"strconv"
)

func part1(data []byte) (total int) {
	//manually go through the array byte by byte and convert numbers
	for i := 0; i < len(data); i++ {
		if (data[i] >= '0' && data[i] <= '9') || data[i] == '-' {
			start := i
			for i < len(data) && ((data[i] >= '0' && data[i] <= '9') || data[i] == '-') {
				i++
			}

			if v, e := strconv.Atoi(string(data[start:i])); e != nil {
				panic("couldn't convert: " + string(data[start:i]))
			} else {
				total += v
			}
		}
	}
	return
}

func sum_numbers(data any) (sum int) {
	switch v := data.(type) {
	case float64:
		sum += int(v)
	case []any:
		for _, value := range v {
			sum += sum_numbers(value)
		}
	case map[string]any:
		red := false
		for _, value := range v {
			if s, ok := value.(string); ok {
				if s == "red" {
					red = true
				}
			}
		}
		if !red {
			for _, value := range v {
				sum += sum_numbers(value)
			}
		}
	}
	return
}

func part2(input []byte) int {
	var data any
	json.Unmarshal(input, &data)
	return sum_numbers(data)
}

func main() {
	data := aocutil.FileReadAll[[]byte]("day12/input.txt")
	aocutil.AOCFinish(part1(data), part2(data))
}
