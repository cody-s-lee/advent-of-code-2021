package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	lines := get_input.Lines("https://adventofcode.com/2015/day/14/input")

	reindeers := make(map[*Reindeer]bool)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		reindeers[Parse(line)] = true
	}

	partOneResult := PartOne(reindeers, 2503)
	partTwoResult := PartTwo(reindeers, 2503)

	fmt.Printf("Part one result: %d\n", partOneResult)
	fmt.Printf("Part two result: %d\n", partTwoResult)
}

func Reset(reindeers map[*Reindeer]bool) {
	for reindeer := range reindeers {
		reindeer.Reset()
	}
}

func (reindeer *Reindeer) Reset() {
	reindeer.EnduranceTimer = reindeer.Endurance
	reindeer.RestTimer = 0
}

func PartOne(reindeers map[*Reindeer]bool, seconds int) int {
	Reset(reindeers)

	distance := make(map[*Reindeer]int)
	for second := 0; second < seconds; second++ {
		for reindeer := range reindeers {
			if reindeer.EnduranceTimer > 0 {
				reindeer.EnduranceTimer--
				distance[reindeer] += reindeer.Speed
				if reindeer.EnduranceTimer <= 0 {
					reindeer.RestTimer = reindeer.Rest
				} else if reindeer.RestTimer > 0 {
					reindeer.RestTimer--
					if reindeer.RestTimer <= 0 {
						reindeer.EnduranceTimer = reindeer.Endurance
					}
				}
			} else {
				reindeer.RestTimer--
				if reindeer.RestTimer <= 0 {
					reindeer.EnduranceTimer = reindeer.Endurance
				}
			}
		}
	}

	fmt.Printf("Reindeer at distance %v after %d seconds\n", distance, seconds)

	maxDistance := distance[leader(distance)]

	return maxDistance
}

func PartTwo(reindeers map[*Reindeer]bool, seconds int) int {
	Reset(reindeers)

	distance := make(map[*Reindeer]int)
	scoreboard := make(map[*Reindeer]int)
	for second := 0; second < seconds; second++ {
		for reindeer := range reindeers {
			if reindeer.EnduranceTimer > 0 {
				reindeer.EnduranceTimer--
				distance[reindeer] += reindeer.Speed
				if reindeer.EnduranceTimer <= 0 {
					reindeer.RestTimer = reindeer.Rest
				} else if reindeer.RestTimer > 0 {
					reindeer.RestTimer--
					if reindeer.RestTimer <= 0 {
						reindeer.EnduranceTimer = reindeer.Endurance
					}
				}
			} else {
				reindeer.RestTimer--
				if reindeer.RestTimer <= 0 {
					reindeer.EnduranceTimer = reindeer.Endurance
				}
			}
		}

		for _, reindeer := range allLeaders(distance) {
			scoreboard[reindeer]++
		}
	}

	fmt.Printf("Reindeer with score %v after %d seconds\n", scoreboard, seconds)

	maxScore := scoreboard[leader(scoreboard)]

	return maxScore
}

func allLeaders(distance map[*Reindeer]int) []*Reindeer {
	leaders := make([]*Reindeer, 0)
	for reindeer := range distance {
		if len(leaders) == 0 || distance[reindeer] == distance[leaders[0]] {
			leaders = append(leaders, reindeer)
		} else if distance[reindeer] > distance[leaders[0]] {
			leaders = []*Reindeer{reindeer}
		}
	}
	return leaders
}

func leader(distance map[*Reindeer]int) *Reindeer {
	var leader *Reindeer
	for reindeer := range distance {
		if distance[reindeer] > distance[leader] {
			leader = reindeer
		}
	}
	return leader
}

type Reindeer struct {
	Name           string
	Speed          int // in km/s
	Endurance      int // in seconds
	EnduranceTimer int // in seconds remaining
	Rest           int // in seconds
	RestTimer      int // in seconds remaining
}

func (r *Reindeer) String() string {
	return fmt.Sprintf("%s", r.Name)
}

//	Vixen can fly 19 km/s for 7 seconds, but then must rest for 124 seconds.
func Parse(description string) *Reindeer {
	name, cdr := get_input.SplitPair(description, " can fly ")
	speedValue, cdr := get_input.SplitPair(cdr, " km/s for ")
	speed, _ := strconv.Atoi(speedValue)
	enduranceValue, cdr := get_input.SplitPair(cdr, " seconds, but then must rest for ")
	endurance, _ := strconv.Atoi(enduranceValue)
	restValue, _ := get_input.SplitPair(cdr, " seconds.")
	rest, _ := strconv.Atoi(restValue)

	return &Reindeer{
		Name:           name,
		Speed:          speed,
		Endurance:      endurance,
		EnduranceTimer: endurance,
		Rest:           rest,
		RestTimer:      0,
	}
}
