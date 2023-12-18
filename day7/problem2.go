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

type CardInHand struct {
	kind  rune
	count int
}

func newCardInHand(kind rune, count int) *CardInHand {
	c := CardInHand{kind: kind, count: count}
	return &c
}

var cardVals = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
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
	f, _ := os.Open("input2.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	hands := make([]*Hand, 0)
	for scanner.Scan() {
		hand := strings.Fields(scanner.Text())
		cards := hand[0]
		bid, _ := strconv.Atoi(hand[1])

		handMap := make(map[rune]int)
		jokers := 0
		for _, card := range cards {
			if card == 'J' {
				jokers++
				continue
			}
			if _, ok := handMap[card]; ok {
				handMap[card]++
			} else {
				handMap[card] = 1
			}
		}

		spread := make([]*CardInHand, 0)
		for kind, count := range handMap {
			spread = append(spread, newCardInHand(kind, count))
		}

		sort.Slice(spread, func(i, j int) bool {
			if spread[i].count == spread[j].count {
				return cardVals[spread[i].kind] > cardVals[spread[j].kind]
			}
			return spread[i].count > spread[j].count
		})

		if len(spread) == 0 {
			spread = append(spread, newCardInHand('A', 0))
		}
		spread[0].count += jokers
		strings.Replace(cards, "J", string(spread[0].kind), jokers)

		spreadLen := len(spread)

		handType := -1
		if spreadLen == 1 {
			handType = 7
		} else if spreadLen == 2 {
			if spread[0].count == 4 {
				handType = 6
			} else {
				handType = 5
			}

		} else if spreadLen == 3 {
			if spread[0].count == 3 && spread[1].count == 1 {
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
