package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"

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

func part1(input string) int {
	winningSets, mySets := parseInput(input)

	sum := 0
	for i, mySet := range mySets {
		winSet := winningSets[i]
		myWinningNums := mySet.Intersect(winSet)

		exponent := float64(myWinningNums.Cardinality() - 1)

		var points float64
		if exponent >= 0 {
			points = math.Pow(2, float64(myWinningNums.Cardinality())-1)
		} else {
			points = 0
		}

		sum += int(math.Round(points))
	}
	return sum
}

func part2(input string) int {
	winningSets, mySets := parseInput(input)
	numCardsMap := make(map[int]int)
	for i := range winningSets {
		numCardsMap[i] = 1
	}

	for cardIdx := 0; cardIdx < len(winningSets); cardIdx++ {
		numCards := numCardsMap[cardIdx]
		winSet := winningSets[cardIdx]
		mySet := mySets[cardIdx]
		myWinningNums := mySet.Intersect(winSet)
		numMatches := myWinningNums.Cardinality()

		for i := 1; i <= numMatches; i++ {
			if cardIdx+i < len(numCardsMap) {
				numCardsMap[cardIdx+i] += numCards
			}
		}
	}

	totalNumCards := 0
	for _, numCards := range numCardsMap {
		totalNumCards += numCards
	}
	return totalNumCards
}

func parseInput(input string) ([]mapset.Set[int], []mapset.Set[int]) {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile("[0-9]+")

	winningSets := make([]mapset.Set[int], len(lines))
	mySets := make([]mapset.Set[int], len(lines))

	for i, line := range lines {
		cards := strings.Split(strings.Split(line, ": ")[1], " | ")
		winNumsStr := re.FindAllString(cards[0], -1)
		myCardNumsStr := re.FindAllString(cards[1], -1)

		winSet, mySet := mapset.NewSet[int](), mapset.NewSet[int]()

		for _, winNumStr := range winNumsStr {
			winNum, _ := strconv.Atoi(winNumStr)
			winSet.Add(winNum)
		}

		for _, myCardNumStr := range myCardNumsStr {
			myCardNum, _ := strconv.Atoi(myCardNumStr)
			mySet.Add(myCardNum)
		}

		winningSets[i] = winSet
		mySets[i] = mySet
	}

	return winningSets, mySets
}
