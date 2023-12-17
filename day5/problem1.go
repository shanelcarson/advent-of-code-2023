package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	src  int
	dest int
	rng  int
	net  int
}

func newInterval(dest int, src int, rng int) *Interval {
	interval := Interval{src: src, dest: dest, rng: rng, net: dest - src}
	return &interval
}

func generateLowestLocationNumber(maps [][]*Interval, seeds []int) {
	for _, metricMap := range maps {
		for idx, current := range seeds {
			for _, interval := range metricMap {
				if current < interval.src {
					break
				} else if current-interval.src < interval.rng {
					seeds[idx] = current + interval.net
					break
				}
			}
		}
	}

}

func main() {
	maps := make([][]*Interval, 7)

	f, _ := os.Open("input1.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	seedLine := scanner.Text()
	seeds := strings.Split(seedLine, " ")[1:]
	startSeeds := make([]int, 0)
	for _, seed := range seeds {
		seedInt, _ := strconv.Atoi(seed)
		startSeeds = append(startSeeds, seedInt)
	}
	scanner.Text()

	mapIdx := -1

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if len(line) == 0 {
			continue
		} else if strings.Contains(line, "map") {
			mapIdx++
			maps[mapIdx] = make([]*Interval, 0, 10)
			continue
		}

		intervalVals := strings.Split(line, " ")
		dest, _ := strconv.Atoi(intervalVals[0])
		src, _ := strconv.Atoi(intervalVals[1])
		rng, _ := strconv.Atoi(intervalVals[2])
		intervalStruct := newInterval(dest, src, rng)
		maps[mapIdx] = append(maps[mapIdx], intervalStruct)
	}
	for _, slice := range maps {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i].src < slice[j].src
		})
	}

	lowestLocationNumber := math.MaxFloat64
	generateLowestLocationNumber(maps, startSeeds)
	for _, location := range startSeeds {
		lowestLocationNumber = math.Min(lowestLocationNumber, float64(location))
	}
	fmt.Println("Lowest Location Number is", int(lowestLocationNumber))

}
