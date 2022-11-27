package main

import "aocutil"

//generates a "aabbccdd.." sequence starting from the given index
func default_sequence(input *[]byte, start int) {
	length := len(*input) - start
	c := byte('a') - 1
	for i := 0; i < length; i++ {
		if i%2 == 0 {
			c++
		}
		(*input)[i+start] = c
	}
}

func increase(input *[]byte) {
	last := len(*input) - 1
	for i := last; i > 0; i-- {
		if (*input)[i] == 'z' {
			(*input)[i] = 'a'
		} else {
			(*input)[i]++
			break
		}
	}
}

func next_password(input string) string {
	arr := []byte(input)
	increase(&arr)

	changed := true
	for changed {
		changed = false
		//check for invalid characters
		for i := 0; i < len(arr); i++ {
			if arr[i] == 'i' || arr[i] == 'o' || arr[i] == 'l' {
				//if invalid increase char by 1 (i/o/l cannot wrap)
				arr[i]++
				//and change following sequence to aabbcc...
				i++
				default_sequence(&arr, i)
				changed = true
				break
			}
		}

		//has no invalid characters anymore at this point
		//check for increasing straight pattern "abc" "bcd" etc
		hasIncreasingPattern := false
		increasingPatternCount := 1
		for i := 0; i < len(arr)-1; i++ {
			if arr[i]+1 == arr[i+1] {
				increasingPatternCount++
			} else {
				if increasingPatternCount >= 3 {
					hasIncreasingPattern = true
					break
				}
				increasingPatternCount = 1
			}
		}
		//if the whole password is a sequence it will never enter the else branch
		//check again here
		if increasingPatternCount >= 3 {
			hasIncreasingPattern = true
		}

		differentPairsCount := 0
		previous := byte(0)
		for i := 0; i < len(arr)-1; i++ {
			if previous != arr[i] && arr[i] == arr[i+1] {
				differentPairsCount++
				previous = arr[i]
				i++
			}
		}

		if !hasIncreasingPattern || differentPairsCount < 2 {
			increase(&arr)
			changed = true
		}
	}

	return string(arr)
}

func main() {
	input := "hxbxwxba"
	part1 := next_password(input)
	part2 := next_password(part1)
	aocutil.AOCFinish(part1, part2)
}
