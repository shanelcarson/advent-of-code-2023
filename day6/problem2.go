package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input1.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	timeLine := strings.Join(strings.Fields(scanner.Text())[1:], "")
	time, _ := strconv.Atoi(timeLine)

	scanner.Scan()
	distanceLine := strings.Join(strings.Fields(scanner.Text())[1:], "")
	recordDistance, _ := strconv.Atoi(distanceLine)

	beatRecordTime := 0
	for i := 0; i <= time; i++ {
		if distance := (time - i) * i; distance > recordDistance {
			beatRecordTime++
		}
	}

	fmt.Println("End Product: ", beatRecordTime)
}
