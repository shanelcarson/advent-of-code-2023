package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input1.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	instructions := scanner.Text()
	scanner.Scan()
	scanner.Text()

	islandsMap := make(map[string][]string)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		key := fields[0]
		destinations := []string{fields[2][1 : len(fields[2])-1], fields[3][:len(fields[3])-1]}
		islandsMap[key] = destinations
	}

	currentIsland := "AAA"
	steps := 0
	for currentIsland != "ZZZ" {
		for _, dir := range instructions {
			if dir == 'L' {
				currentIsland = islandsMap[currentIsland][0]
			} else {
				currentIsland = islandsMap[currentIsland][1]
			}
			steps++

			if currentIsland == "ZZZ" {
				break
			}
		}
	}
	fmt.Println("Total steps:", steps)
}
