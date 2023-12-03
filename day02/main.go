package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/basokant/advent-of-code-2023/util"
)

type Colour int

const (
	Blue Colour = iota
	Red
	Green
)

func (c Colour) String() string {
	switch c {
	case Blue:
		return "blue"
	case Red:
		return "red"
	case Green:
		return "green"
	}
	return "unknown"
}

func StringToColour(s string) (Colour, error) {
	switch s {
	case "blue":
		return Blue, nil
	case "red":
		return Red, nil
	case "green":
		return Green, nil
	}
	return -1, nil
}

type Game struct {
	cubeSets []map[Colour]int
	id       int
}

func (g Game) isPossible(cubeCounts map[Colour]int) bool {
	for _, cubeSet := range g.cubeSets {
		for colour, count := range cubeCounts {
			if cubeSet[colour] > count {
				return false
			}
		}
	}
	return true
}

func (g Game) getMinCubeCounts() map[Colour]int {
	minCubeCounts := map[Colour]int{
		Red:   0,
		Green: 0,
		Blue:  0,
	}

	for _, cubeSet := range g.cubeSets {
		for colour, minCount := range minCubeCounts {
			minCubeCounts[colour] = max(cubeSet[colour], minCount)
		}
	}

	return minCubeCounts
}

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
	cubeCounts := map[Colour]int{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	games := parseInput(input)

	possibleGameIds := make([]int, len(games))
	for i, game := range games {
		if game.isPossible(cubeCounts) {
			possibleGameIds[i] = game.id
		}
	}

	sum := 0
	for _, gameId := range possibleGameIds {
		sum += gameId
	}

	return sum
}

func part2(input string) int {
	games := parseInput(input)

	powers := make([]int, len(games))
	for i, game := range games {
		minCubeCounts := game.getMinCubeCounts()

		power := 1
		for _, count := range minCubeCounts {
			power *= count
		}

		powers[i] = power
	}

	sum := 0
	for _, power := range powers {
		sum += power
	}

	return sum
}

func parseInput(input string) []Game {
	games := []Game{}
	for _, line := range strings.Split(input, "\n") {
		game := Game{}

		tokens := strings.Split(line, ": ")
		gameInfo, gameSets := tokens[0], tokens[1]
		gameId, err := strconv.Atoi(strings.Split(gameInfo, " ")[1])

		if err == nil {
			game.id = gameId
		}

		game.cubeSets = createCubeSets(gameSets)
		games = append(games, game)
	}

	return games
}

func createCubeSets(gameSetsStr string) []map[Colour]int {
	cubeSets := []map[Colour]int{}
	for _, set := range strings.Split(gameSetsStr, "; ") {
		cubeSet := map[Colour]int{}
		for _, subSet := range strings.Split(set, ", ") {
			tokens := strings.Split(subSet, " ")
			countStr, colourStr := tokens[0], tokens[1]

			count, _ := strconv.Atoi(countStr)
			colour, _ := StringToColour(colourStr)
			cubeSet[colour] = count
		}
		cubeSets = append(cubeSets, cubeSet)
	}

	return cubeSets
}
