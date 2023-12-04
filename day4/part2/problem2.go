package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input2.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	totalCards := make(map[int]int)
	playing := true
	for fileScanner.Scan() {
		line := fileScanner.Text()
		tokens := strings.Split(line, ": ")
		cardId := strings.Split(tokens[0], " ")
		id, _ := strconv.Atoi(cardId[len(cardId)-1])
		// fmt.Println("Card Id", id)
		if _, ok := totalCards[id]; ok {
			totalCards[id] = totalCards[id] + 1
		} else {
			totalCards[id] = 1
		}

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

			winningMatches := 0
			for _, match := range numMap {
				if match {
					winningMatches += 1
				}
			}
			if winningMatches > 0 {	
				for i := 1; i <= winningMatches; i++ {
					if _, ok := totalCards[id+i]; !ok {
						totalCards[id+i] = 0
					}
					totalCards[id+i] += totalCards[id]
				}
			}
	}
	fmt.Println(totalCards)
	for _, cardCount := range totalCards {
		sum += cardCount
	}

	fmt.Println("Sum is", sum)

	readFile.Close()
}
