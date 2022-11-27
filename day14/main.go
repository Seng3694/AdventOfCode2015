package main

import (
	"aocutil"
	"strconv"
	"strings"
)

type ReindeerData struct {
	kms, seconds, rest int
}

func parse_line(line string) ReindeerData {
	split := strings.Split(line, " ")
	kms, _ := strconv.Atoi(split[3])
	forSeconds, _ := strconv.Atoi(split[6])
	withRest, _ := strconv.Atoi(split[13])

	return ReindeerData{
		kms:     kms,
		seconds: forSeconds,
		rest:    withRest,
	}
}

type ReindeerRaceData struct {
	isResting                       bool
	points, km, timer, timerElapsed int
}

func calc_race_results(data []ReindeerData, seconds int) (maxDistance, maxPoints int) {
	raceData := make([]ReindeerRaceData, len(data))

	for i := 0; i < len(raceData); i++ {
		raceData[i].timer = data[i].seconds
	}

	for i := 0; i < seconds; i++ {
		leader := -1
		leaderKm := 0

		for j := 0; j < len(data); j++ {
			rd := &raceData[j]
			d := &data[j]

			rd.timerElapsed++
			if !rd.isResting {
				rd.km += d.kms
				if rd.timer == rd.timerElapsed {
					rd.timer = d.rest
					rd.timerElapsed = 0
					rd.isResting = true
				}
			} else {
				if rd.timer == rd.timerElapsed {
					rd.timer = d.seconds
					rd.timerElapsed = 0
					rd.isResting = false
				}
			}

			if rd.km > leaderKm {
				leaderKm = rd.km
				leader = j
			}
		}

		raceData[leader].points++
	}

	for _, rd := range raceData {
		if rd.points > maxPoints {
			maxPoints = rd.points
		}
		if rd.km > maxDistance {
			maxDistance = rd.km
		}
	}
	return
}

func main() {
	data := make([]ReindeerData, 0, 10)
	aocutil.FileReadAllLines("day14/input.txt", func(s string) {
		data = append(data, parse_line(s))
	})

	dist, points := calc_race_results(data, 2503)
	aocutil.AOCFinish(dist, points)
}
