package main

import (
	"aocutil"
	"sort"
	"strconv"
	"strings"
)

type Dimensions struct {
	w, h, l int
}

func parse_dimensions(str string) (dim Dimensions) {
	for i, stringValue := range strings.Split(str, "x") {
		if value, err := strconv.Atoi(stringValue); err != nil {
			panic(err)
		} else {
			switch i {
			case 0:
				dim.w = value
			case 1:
				dim.h = value
			case 2:
				dim.l = value
			}
		}
	}
	return
}

func calculate_total_area(dim Dimensions) int {
	area1 := dim.l * dim.w
	area2 := dim.w * dim.h
	area3 := dim.h * dim.l
	slack := aocutil.Min(area1, aocutil.Min(area2, area3))
	return 2*area1 + 2*area2 + 2*area3 + slack
}

func calculate_ribbon_length(dim Dimensions) int {
	arr := []int{dim.w, dim.h, dim.l}
	sort.Ints(arr)
	return arr[0] + arr[0] + arr[1] + arr[1] + dim.w*dim.h*dim.l
}

func main() {
	totalArea := 0
	totalRibbonLength := 0
	aocutil.FileReadAllLines("day2/input.txt", func(s string) {
		dim := parse_dimensions(s)
		totalArea += calculate_total_area(dim)
		totalRibbonLength += calculate_ribbon_length(dim)
	})
	aocutil.AOCFinish(totalArea, totalRibbonLength)
}
