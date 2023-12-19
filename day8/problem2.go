package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type IslandHopper struct {
	currentIsland  string
	onEndingIsland int
}

func newIslandHopper(currentIsland string) *IslandHopper {
	i := IslandHopper{currentIsland: currentIsland, onEndingIsland: -1}
	return &i
}

func main() {
	f, _ := os.Open("input2.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	instructions := scanner.Text()
	scanner.Scan()
	scanner.Text()

	islandHoppers := make([]*IslandHopper, 0)

	islandsMap := make(map[string][]string)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		key := fields[0]
		destinations := []string{fields[2][1 : len(fields[2])-1], fields[3][:len(fields[3])-1]}
		islandsMap[key] = destinations

		if key[len(key)-1] == 'A' {
			islandHoppers = append(islandHoppers, newIslandHopper(key))
		}
	}

	for _, islandHopper := range islandHoppers {
		fmt.Println(islandHopper)
	}
	steps := 0
	islandHoppersDone := 0
	for islandHoppersDone < len(islandHoppers) {
		for _, dir := range instructions {
			nextIslandIdx := 0

			if dir == 'R' {
				nextIslandIdx = 1
			}

			for _, islandHopper := range islandHoppers {
				islandHopper.currentIsland = islandsMap[islandHopper.currentIsland][nextIslandIdx]
				if islandHopper.currentIsland[len(islandHopper.currentIsland)-1] == 'Z' {
					islandHopper.onEndingIsland = 1
					islandHoppersDone++
				} else if islandHopper.onEndingIsland == 1 {
					islandHopper.onEndingIsland = -1
					islandHoppersDone--
				}
			}
			fmt.Println(islandHoppers)
			steps++
		}
	}
	fmt.Println("Total steps:", steps)
}
