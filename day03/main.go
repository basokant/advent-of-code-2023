package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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

func part1(input string) int {
	partNums := findPartNumbers(input)

	sum := 0
	for _, partNum := range partNums {
		sum += partNum
	}

	return sum
}

func part2(input string) int {
	gearRatios := findGearRatios(input)

	sum := 0
	for _, gearRatio := range gearRatios {
		sum += gearRatio
	}
	return sum
}

func isAdjacentToSymbol(grid [][]rune, rowNum int, colNums []int) bool {
	directions := [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
	}

	for _, colNum := range colNums {
		for _, direction := range directions {
			dx, dy := direction[0], direction[1]

			isValidAdjacent := 0 <= rowNum+dy && rowNum+dy < len(grid) && 0 <= colNum+dx && colNum+dx < len(grid[rowNum])
			if !isValidAdjacent {
				continue
			}

			adjacentChar := grid[rowNum+dy][colNum+dx]
			if !unicode.IsDigit(adjacentChar) && adjacentChar != '.' {
				return true
			}
		}
	}

	return false
}

func getGearRatio(grid [][]rune, potentialParts []Part, rowNum int, colNum int) int {
	directions := [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
	}

	parts := make([]Part, len(potentialParts))
	copy(parts, potentialParts)

	adjacentParts := []Part{}
	for _, direction := range directions {
		dx, dy := direction[0], direction[1]

		isValidAdjacent := 0 <= rowNum+dy && rowNum+dy < len(grid) && 0 <= colNum+dx && colNum+dx < len(grid[rowNum])
		if !isValidAdjacent {
			continue
		}

		adjacentChar := grid[rowNum+dy][colNum+dx]
		if unicode.IsDigit(adjacentChar) {
			for i, p := range parts {
				isAdjacentPart := p.rowNum == rowNum+dy && colNum+dx >= p.colNums[0] && colNum+dx <= p.colNums[len(p.colNums)-1]
				if !isAdjacentPart {
					continue
				}
				adjacentParts = append(adjacentParts, p)
				parts = slices.Delete(parts, i, i+1)

				if len(adjacentParts) >= 2 {
					return 0
				}
			}
		}
	}

	if len(adjacentParts) == 2 {
		return adjacentParts[0].number * adjacentParts[1].number
	}

	return 0
}

type Part struct {
	colNums []int
	number  int
	rowNum  int
}

func getPotentialParts(lines []string) []Part {
	potentialParts := []Part{}
	re := regexp.MustCompile("[0-9]+")

	for rowNum, line := range lines {
		nums := re.FindAllString(line, -1)
		numCols := re.FindAllStringIndex(line, -1)

		for i, numStr := range nums {
			numCols[i][1] -= 1
			num, _ := strconv.Atoi(numStr)
			potentialPart := Part{
				number:  num,
				rowNum:  rowNum,
				colNums: numCols[i],
			}
			potentialParts = append(potentialParts, potentialPart)
		}
	}

	return potentialParts
}

func findPartNumbers(input string) []int {
	lines := strings.Split(input, "\n")
	grid := parseInput(input)

	partNumbers := []int{}
	potentialParts := getPotentialParts(lines)

	for _, potentialPart := range potentialParts {
		if isAdjacentToSymbol(grid, potentialPart.rowNum, potentialPart.colNums) {
			partNumbers = append(partNumbers, potentialPart.number)
		}
	}

	return partNumbers
}

func findGearRatios(input string) []int {
	lines := strings.Split(input, "\n")
	grid := parseInput(input)

	gearRatios := []int{}
	potentialParts := getPotentialParts(lines)
	re := regexp.MustCompile("[*]")

	for rowNum := range grid {
		gearIdxs := re.FindAllStringIndex(lines[rowNum], -1)

		for _, colNums := range gearIdxs {
			gearRatio := getGearRatio(grid, potentialParts, rowNum, colNums[0])
			gearRatios = append(gearRatios, gearRatio)
		}
	}

	return gearRatios
}

func parseInput(input string) [][]rune {
	lines := strings.Split(input, "\n")

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		row := make([]rune, len(line))
		for j, char := range line {
			row[j] = char
		}
		grid[i] = row
	}

	return grid
}
