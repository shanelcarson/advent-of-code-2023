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

type MetricMapInterval struct {
	src      int
	srcStart int
	srcEnd   int
	dest     int
	net      int
}

func newInterval(dest int, src int, rng int) *MetricMapInterval {
	interval := MetricMapInterval{src: src, dest: dest, net: dest - src, srcStart: src, srcEnd: src + rng - 1}
	return &interval
}

type SeedInterval struct {
	start int
	end   int
}

func newSeedInterval(start int, end int) *SeedInterval {
	seedInterval := SeedInterval{start: start, end: end}
	return &seedInterval
}

func generateLowestLocationNumber(maps [][]*MetricMapInterval, seedIntervals []*SeedInterval) []*SeedInterval {
	for _, metricMap := range maps {
		for i := 0; i < len(seedIntervals); i++ {

			currentInterval := seedIntervals[i]

			for _, metricMapInterval := range metricMap {
				if metricMapInterval.srcStart > currentInterval.end {
					break
				} else if metricMapInterval.srcEnd < currentInterval.start {
					continue
				} else {
					var leftSeedInterval *SeedInterval = nil
					var rightSeedInterval *SeedInterval = nil

					if currentInterval.start < metricMapInterval.srcStart {
						leftSeedInterval = newSeedInterval(currentInterval.start, metricMapInterval.srcStart-1)
					}

					if metricMapInterval.srcEnd < currentInterval.end {
						rightSeedInterval = newSeedInterval(metricMapInterval.srcEnd+1, currentInterval.end)
					}

					newStart := int(math.Max(float64(metricMapInterval.srcStart), float64(currentInterval.start)))
					newEnd := int(math.Min(float64(metricMapInterval.srcEnd), float64(currentInterval.end)))
					currentInterval.start = newStart + metricMapInterval.net
					currentInterval.end = newEnd + metricMapInterval.net

					if leftSeedInterval != nil {
						seedIntervals = append(seedIntervals, leftSeedInterval)
					}
					if rightSeedInterval != nil {
						seedIntervals = append(seedIntervals, rightSeedInterval)
					}
					break
				}
			}
		}
		sort.Slice(seedIntervals, func(i, j int) bool {
			return seedIntervals[i].start < seedIntervals[j].start
		})
	}
	return seedIntervals
}

func main() {
	maps := make([][]*MetricMapInterval, 7)

	f, _ := os.Open("input1.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	seedLine := scanner.Text()
	seeds := strings.Split(seedLine, " ")[1:]
	startSeedIntervals := make([]*SeedInterval, 0)
	for i := 0; i < len(seeds); i += 2 {
		seedIntStart, _ := strconv.Atoi(seeds[i])
		seedIntRange, _ := strconv.Atoi(seeds[i+1])
		startSeedIntervals = append(startSeedIntervals, newSeedInterval(seedIntStart, seedIntStart+seedIntRange-1))
	}
	scanner.Text()

	mapIdx := -1

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if len(line) == 0 {
			continue
		} else if strings.Contains(line, "map") {
			mapIdx++
			maps[mapIdx] = make([]*MetricMapInterval, 0, 10)
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
	sort.Slice(startSeedIntervals, func(i, j int) bool {
		return startSeedIntervals[i].start < startSeedIntervals[j].start
	})
	endSeedIntervals := generateLowestLocationNumber(maps, startSeedIntervals)
	lowestLocationNumber := endSeedIntervals[0].start

	fmt.Println("Lowest Location Number is", lowestLocationNumber)

}
