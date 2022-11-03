package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

/*
--- Day 4: The Ideal Stocking Stuffer ---

Santa needs help mining some AdventCoins (very similar to bitcoins) to use as gifts for all the economically forward-thinking little girls and boys.

To do this, he needs to find MD5 hashes which, in hexadecimal, start with at least five zeroes. The input to the MD5 hash is some secret key (your puzzle input, given below) followed by a number in decimal. To mine AdventCoins, you must find Santa the lowest positive number (no leading zeroes: 1, 2, 3, ...) that produces such a hash.

For example:

-	If your secret key is abcdef, the answer is 609043, because the MD5 hash of abcdef609043 starts with five zeroes (000001dbbfa...), and it is the lowest such number to do so.
-   If your secret key is pqrstuv, the lowest number it combines with to make an MD5 hash starting with five zeroes is 1048970; that is, the MD5 hash of pqrstuv1048970 looks like 000006136ef....

--- Part Two ---

Now find one that starts with six zeroes.
*/

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func day4() (string, string) {
	data, err := os.ReadFile("input/day4.txt")
	if err != nil {
		panic(err)
	}

	i := 0
	part1Solution := 0
	part2Solution := 0
	part1SolutionFound := false

	for {
		hash := GetMD5Hash(fmt.Sprintf("%s%d", data, i))

		if part1SolutionFound {
			if strings.HasPrefix(hash, "000000") {
				part2Solution = i
				break
			}
		} else {
			if strings.HasPrefix(hash, "00000") {
				part1Solution = i
				part1SolutionFound = true
			}
		}
		i++
	}

	return fmt.Sprint(part1Solution), fmt.Sprint(part2Solution)
}
