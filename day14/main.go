package main

import (
	"aocutil"
	"strconv"
	"strings"
)

/*
--- Day 14: Reindeer Olympics ---

This year is the Reindeer Olympics! Reindeer can fly at high speeds, but must rest occasionally to recover their energy. Santa would like to know which of his reindeer is fastest, and so he has them race.

Reindeer can only either be flying (always at their top speed) or resting (not moving at all), and always spend whole seconds in either state.

For example, suppose you have the following Reindeer:

    Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
    Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.

After one second, Comet has gone 14 km, while Dancer has gone 16 km. After ten seconds, Comet has gone 140 km, while Dancer has gone 160 km. On the eleventh second, Comet begins resting (staying at 140 km), and Dancer continues on for a total distance of 176 km. On the 12th second, both reindeer are resting. They continue to rest until the 138th second, when Comet flies for another ten seconds. On the 174th second, Dancer flies for another 11 seconds.

In this example, after the 1000th second, both reindeer are resting, and Comet is in the lead at 1120 km (poor Dancer has only gotten 1056 km by that point). So, in this situation, Comet would win (if the race ended at 1000 seconds).

Given the descriptions of each reindeer (in your puzzle input), after exactly 2503 seconds, what distance has the winning reindeer traveled?

--- Part Two ---

Seeing how reindeer move in bursts, Santa decides he's not pleased with the old scoring system.

Instead, at the end of each second, he awards one point to the reindeer currently in the lead. (If there are multiple reindeer tied for the lead, they each get one point.) He keeps the traditional 2503 second time limit, of course, as doing otherwise would be entirely ridiculous.

Given the example reindeer from above, after the first second, Dancer is in the lead and gets one point. He stays in the lead until several seconds into Comet's second burst: after the 140th second, Comet pulls into the lead and gets his first point. Of course, since Dancer had been in the lead for the 139 seconds before that, he has accumulated 139 points by the 140th second.

After the 1000th second, Dancer has accumulated 689 points, while poor Comet, our old champion, only has 312. So, with the new scoring system, Dancer would win (if the race ended at 1000 seconds).

Again given the descriptions of each reindeer (in your puzzle input), after exactly 2503 seconds, how many points does the winning reindeer have?


*/

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
