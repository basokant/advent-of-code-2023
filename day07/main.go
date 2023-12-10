package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/basokant/advent-of-code-2023/util"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

type HandClass int

const (
	HighCard HandClass = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	cards []rune
	class HandClass
	bid   int
}

func getHandClass(cards []rune, useJokers bool) HandClass {
	if len(cards) != 5 {
		panic("Hand does not have 5 cards, " + string(cards))
	}

	cardCounts := make(map[rune]int)
	jokers := 0
	for _, card := range cards {
		if useJokers && card == 'J' {
			jokers += 1
		} else {
			cardCounts[card] += 1
		}
	}

	maxOfAKind := getMaxOfAKind(cardCounts, jokers)
	switch true {
	case maxOfAKind == 5:
		return FiveOfAKind
	case maxOfAKind == 4:
		return FourOfAKind
	case isFullHouse(cardCounts, jokers):
		return FullHouse
	case maxOfAKind == 3:
		return ThreeOfAKind
	case isTwoPair(cardCounts, jokers):
		return TwoPair
	case maxOfAKind == 2:
		return OnePair
	default:
		return HighCard
	}
}

func getMaxOfAKind(cardCounts map[rune]int, jokers int) int {
	counts := util.Values(cardCounts)
	maxOfAKind := 0
	if len(counts) > 0 {
		maxOfAKind = slices.Max(counts)
	}
	maxOfAKind += jokers
	return maxOfAKind
}

func isFullHouse(cardCounts map[rune]int, jokers int) bool {
	useJokers := jokers > 0
	if (!useJokers && len(cardCounts) != 2) || (useJokers && 1 <= len(cardCounts) && 3 <= len(cardCounts)) {
		return false
	}

	singleCount, pairCount, threesCount := 0, 0, 0
	for _, count := range cardCounts {
		switch count {
		case 1:
			singleCount += 1
		case 2:
			pairCount += 1
		case 3:
			threesCount += 1
		}
	}

	if !useJokers && len(cardCounts) == 2 || useJokers && jokers == 0 {
		return pairCount == 1 && threesCount == 1
	} else if useJokers {
		switch jokers {
		case 1:
			return threesCount == 1 || pairCount == 2
		case 2:
			return threesCount == 1 || pairCount == 1
		case 3, 4, 5:
			return true
		}
	}

	return false
}

func isTwoPair(cardCounts map[rune]int, jokers int) bool {
	useJokers := jokers > 0
	if (!useJokers && len(cardCounts) != 3) || (useJokers && 1 <= len(cardCounts) && 4 <= len(cardCounts)) {
		return false
	}

	singleCount, pairCount, threesCount := 0, 0, 0
	for _, count := range cardCounts {
		switch count {
		case 1:
			singleCount += 1
		case 2:
			pairCount += 1
		case 3:
			threesCount += 1
		}
	}

	if !useJokers && len(cardCounts) == 3 || useJokers && jokers == 0 {
		return pairCount == 2 && singleCount == 1
	} else if useJokers {
		switch jokers {
		case 1:
			return pairCount == 1
		case 2:
			return threesCount == 0
		case 3, 4, 5:
			return true
		}
	}

	return false
}

func getCardRank(card rune, useJokers bool) (int, error) {
	if unicode.IsDigit(card) {
		return int(card - '0'), nil
	}

	switch true {
	case card == 'T':
		return 10, nil
	case !useJokers && card == 'J':
		return 11, nil
	case card == 'Q':
		return 12, nil
	case card == 'K':
		return 13, nil
	case card == 'A':
		return 14, nil
	default:
		return -1, nil
	}
}

func compareCards(a rune, b rune) int {
	aRank, _ := getCardRank(a, false)
	bRank, _ := getCardRank(b, false)

	if aRank == bRank {
		return 0
	} else if aRank > bRank {
		return 1
	} else {
		return -1
	}
}

func compareCardsWithJokers(a rune, b rune) int {
	aRank, _ := getCardRank(a, true)
	bRank, _ := getCardRank(b, true)

	if aRank == bRank {
		return 0
	} else if aRank > bRank {
		return 1
	} else {
		return -1
	}
}

func compareHands(a Hand, b Hand) int {
	if a.class > b.class {
		return 1
	} else if a.class < b.class {
		return -1
	}

	for i, aCard := range a.cards {
		bCard := b.cards[i]
		comp := compareCards(aCard, bCard)
		if comp != 0 {
			return comp
		}
	}

	return 0
}

func compareHandsWithJokers(a Hand, b Hand) int {
	if a.class > b.class {
		return 1
	} else if a.class < b.class {
		return -1
	}

	for i, aCard := range a.cards {
		bCard := b.cards[i]
		comp := compareCardsWithJokers(aCard, bCard)
		if comp != 0 {
			return comp
		}
	}

	return 0
}

func part1(input string) int {
	hands := parseInput(input, false)
	slices.SortFunc[[]Hand, Hand](hands, compareHands)

	for _, hand := range hands {
		fmt.Printf("hand %s, class %d, bid %d\n", string(hand.cards), hand.class, hand.bid)
	}

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.bid
	}

	return totalWinnings
}

func part2(input string) int {
	hands := parseInput(input, true)
	slices.SortFunc[[]Hand, Hand](hands, compareHandsWithJokers)

	for _, hand := range hands {
		fmt.Printf("hand %s, class %d, bid %d\n", string(hand.cards), hand.class, hand.bid)
	}

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.bid
	}

	return totalWinnings
}

func parseInput(input string, useJokers bool) []Hand {
	lines := strings.Split(input, "\n")
	hands := []Hand{}

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		handInput, bidInput := tokens[0], tokens[1]

		bid, _ := strconv.Atoi(bidInput)
		cards := []rune(handInput)
		class := getHandClass(cards, useJokers)

		hand := Hand{
			cards: cards,
			class: class,
			bid:   bid,
		}
		hands = append(hands, hand)
	}

	return hands
}
