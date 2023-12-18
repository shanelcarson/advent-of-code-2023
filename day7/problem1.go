package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards    string
	handType int
	bid      int
}

func newHand(cards string, handType, bid int) *Hand {
	h := Hand{cards: cards, handType: handType, bid: bid}
	return &h
}

var cardVals = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func main() {
	f, _ := os.Open("input1.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	hands := make([]*Hand, 0)
	for scanner.Scan() {
		hand := strings.Fields(scanner.Text())
		cards := hand[0]
		bid, _ := strconv.Atoi(hand[1])

		handMap := make(map[rune]int)
		for _, card := range cards {
			if _, ok := handMap[card]; ok {
				handMap[card]++
			} else {
				handMap[card] = 1
			}
		}

		spread := make([]int, 0)
		for _, count := range handMap {
			spread = append(spread, count)
		}

		sort.Slice(spread, func(i, j int) bool {
			return spread[i] > spread[j]
		})

		spreadLen := len(spread)

		handType := -1
		if spreadLen == 1 {
			handType = 7
		} else if spreadLen == 2 {
			if spread[0] == 4 {
				handType = 6
			} else {
				handType = 5
			}

		} else if spreadLen == 3 {
			if spread[0] == 3 && spread[1] == 1 {
				handType = 4
			} else {
				handType = 3
			}

		} else if spreadLen == 4 {
			handType = 2
		} else {
			handType = 1
		}

		newCard := newHand(cards, handType, bid)
		hands = append(hands, newCard)

	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			for cardIdx := 0; cardIdx < len(hands[i].cards); cardIdx++ {
				cardIValue := cardVals[rune(hands[i].cards[cardIdx])]
				cardJValue := cardVals[rune(hands[j].cards[cardIdx])]
				if cardIValue == cardJValue {
					continue
				}
				res := cardVals[rune(hands[i].cards[cardIdx])] < cardVals[rune(hands[j].cards[cardIdx])]
				return res
			}
		}
		return hands[i].handType < hands[j].handType
	})
	winnings := 0
	for i := range hands {
		winnings += hands[i].bid * (i + 1)
	}
	fmt.Println("Total winnings are", winnings)
}
