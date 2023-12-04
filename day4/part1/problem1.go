package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input1.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		tokens := strings.Split(line, ": ")
		numStrs := strings.Split(tokens[1], " | ")
		cardNums := strings.Split(numStrs[0], " ")
		myNums := strings.Split(numStrs[1], " ")

		numMap := make(map[string]bool)
		for _, num := range cardNums {
			num = strings.Trim(num, " ")
			if num == "" {
				continue
			}
			numMap[num] = false
		}
		for _, num := range myNums {
			num = strings.Trim(num, " ")
			if num == "" {
				continue
			}
			if value, ok := numMap[num]; ok && !value {
				numMap[num] = true
			}
		}

		winningMatches := 0.0
		for _, match := range numMap {
			if match {
				winningMatches += 1.0
			}
		}
		if winningMatches > 0.0 {
			sum += int(math.Pow(2.0, winningMatches-1.0))
		}

	}
	fmt.Println("Sum is", sum)

	readFile.Close()
}
