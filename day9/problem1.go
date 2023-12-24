package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sendExtrapolateVals(ch chan<- int, inputVals []int) {
	valToSend := extrapolateVals(inputVals)
	ch <- valToSend
}
func extrapolateVals(inputVals []int) int {
	allZeroes := true
	newVals := make([]int, 0)
	for i := 0; i < len(inputVals)-1; i++ {
		val1 := inputVals[i]
		val2 := inputVals[i+1]
		if val1 != 0 || val2 != 0 && allZeroes {
			allZeroes = false
		}
		newVals = append(newVals, val2-val1)
	}
	if allZeroes {
		return 0
	} else {
		childEndVal := extrapolateVals(newVals)
		return inputVals[len(inputVals)-1] + childEndVal
	}
}

func main() {
	f, _ := os.Open("input1.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	allVals := make([][]int, 0)
	for scanner.Scan() {
		valsLine := strings.Fields(scanner.Text())
		someVals := make([]int, 0)
		for _, val := range valsLine {
			num, _ := strconv.Atoi(val)
			someVals = append(someVals, num)
		}
		allVals = append(allVals, someVals)
	}

	ch := make(chan int, 0)

	for _, inputVals := range allVals {
		go sendExtrapolateVals(ch, inputVals)
	}
	sum := 0
	for i := 0; i < len(allVals); i++ {
		val := <-ch
		sum += val
	}
	fmt.Println("Sum is", sum)
}
