package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

/*
--- Day 12: JSAbacusFramework.io ---

Santa's Accounting-Elves need help balancing the books after a recent order. Unfortunately, their accounting software uses a peculiar storage format. That's where you come in.

They have a JSON document which contains a variety of things: arrays ([1,2,3]), objects ({"a":1, "b":2}), numbers, and strings. Your first job is to simply find all of the numbers throughout the document and add them together.

For example:

    [1,2,3] and {"a":2,"b":4} both have a sum of 6.
    [[[3]]] and {"a":{"b":4},"c":-1} both have a sum of 3.
    {"a":[-1,1]} and [-1,{"a":1}] both have a sum of 0.
    [] and {} both have a sum of 0.

You will not encounter any strings containing numbers.

What is the sum of all numbers in the document?

--- Part Two ---

Uh oh - the Accounting-Elves have realized that they double-counted everything red.

Ignore any object (and all of its children) which has any property with the value "red". Do this only for objects ({...}), not arrays ([...]).

    [1,2,3] still has a sum of 6.
    [1,{"c":"red","b":2},3] now has a sum of 4, because the middle object is ignored.
    {"d":"red","e":[1,2,3,4],"f":5} now has a sum of 0, because the entire structure is ignored.
    [1,"red",5] has a sum of 6, because "red" in an array has no effect.


*/

func day12_part1(data []byte) (total int) {
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

func day12_sum_numbers(data any) (sum int) {
	switch v := data.(type) {
	case float64:
		sum += int(v)
	case []any:
		for _, value := range v {
			sum += day12_sum_numbers(value)
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
				sum += day12_sum_numbers(value)
			}
		}
	}
	return
}

func day12_part2(input []byte) int {
	var data any
	json.Unmarshal(input, &data)
	return day12_sum_numbers(data)
}

func day12() (string, string) {
	data, err := os.ReadFile("input/day12.txt")
	if err != nil {
		panic(err)
	}

	return fmt.Sprint(day12_part1(data)), fmt.Sprint(day12_part2(data))
}
