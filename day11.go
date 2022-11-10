package main

/*
--- Day 11: Corporate Policy ---

Santa's previous password expired, and he needs help choosing a new one.

To help him remember his new password after the old one expires, Santa has devised a method of coming up with a password based on the previous one. Corporate policy dictates that passwords must be exactly eight lowercase letters (for security reasons), so he finds his new password by incrementing his old password string repeatedly until it is valid.

Incrementing is just like counting with numbers: xx, xy, xz, ya, yb, and so on. Increase the rightmost letter one step; if it was z, it wraps around to a, and repeat with the next letter to the left until one doesn't wrap around.

Unfortunately for Santa, a new Security-Elf recently started, and he has imposed some additional password requirements:

    Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
    Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
    Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.

For example:

    hijklmmn meets the first requirement (because it contains the straight hij) but fails the second requirement requirement (because it contains i and l).
    abbceffg meets the third requirement (because it repeats bb and ff) but fails the first requirement.
    abbcegjk fails the third requirement, because it only has one double letter (bb).
    The next password after abcdefgh is abcdffaa.
    The next password after ghijklmn is ghjaabcc, because you eventually skip all the passwords that start with ghi..., since i is not allowed.

Given Santa's current password (your puzzle input), what should his next password be?

--- Part Two ---

Santa's password expired again. What's the next one?

*/

//generates a "aabbccdd.." sequence starting from the given index
func day11_default_sequence(input *[]byte, start int) {
	length := len(*input) - start
	c := byte('a') - 1
	for i := 0; i < length; i++ {
		if i%2 == 0 {
			c++
		}
		(*input)[i+start] = c
	}
}

func day11_increase(input *[]byte) {
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

func day11_next_password(input string) string {
	arr := []byte(input)
	day11_increase(&arr)

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
				day11_default_sequence(&arr, i)
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
			day11_increase(&arr)
			changed = true
		}
	}

	return string(arr)
}

func day11() (string, string) {
	input := "hxbxwxba"
	part1 := day11_next_password(input)
	part2 := day11_next_password(part1)
	return part1, part2
}
