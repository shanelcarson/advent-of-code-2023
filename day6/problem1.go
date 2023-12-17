package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time           int
	recordDistance int
}

func newRace(time, recordDistance int) *Race {
	r := Race{time: time, recordDistance: recordDistance}
	return &r
}

func main() {
	f, _ := os.Open("input1.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	timesLine := strings.Fields(scanner.Text())[1:]

	scanner.Scan()
	distancesLine := strings.Fields(scanner.Text())[1:]

	races := make([]*Race, 0)
	for i := 0; i < len(timesLine); i++ {
		time, _ := strconv.Atoi(timesLine[i])
		distance, _ := strconv.Atoi(distancesLine[i])
		races = append(races, newRace(time, distance))
	}

	runningProduct := 1
	for _, race := range races {
		raceTotalTime := race.time
		beatRecordTime := 0

		for i := 0; i <= raceTotalTime; i++ {
			if distance := (raceTotalTime - i) * i; distance > race.recordDistance {
				beatRecordTime++
			}
		}
		runningProduct *= beatRecordTime
	}
	fmt.Println("End Product: ", runningProduct)
}
