package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/basokant/advent-of-code-2023/util"
	"github.com/samber/lo"
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
	histories := parseInput(input, false)

	extrapolatedValues := lo.Map(histories, func(history []int, _ int) int {
		val, _ := extrapolateRight(history)
		return val
	})

	return lo.Sum(extrapolatedValues)
}

func part2(input string) int {
	histories := parseInput(input, false)

	extrapolatedValues := lo.Map(histories, func(history []int, _ int) int {
		val, _ := extrapolateLeft(history)
		return val
	})

	return lo.Sum(extrapolatedValues)
}

func parseInput(input string, _ bool) [][]int {
	lines := strings.Split(input, "\n")

	histories := lo.Map(lines, func(line string, _ int) []int {
		historyInput := strings.Split(line, " ")
		return lo.Map(historyInput, func(numStr string, _ int) int {
			num, _ := strconv.Atoi(numStr)
			return num
		})
	})

	return histories
}

func computeDifferences(differences [][]int) ([][]int, error) {
	lastDifferences, err := lo.Last(differences)
	if err != nil {
		return nil, err
	} else if slices.Equal(lo.Uniq(lastDifferences), []int{0}) {
		return differences, nil
	}

	newDifferences := lo.RepeatBy(len(lastDifferences)-1, func(i int) int {
		return lastDifferences[i+1] - lastDifferences[i]
	})

	differences = append(differences, newDifferences)
	return computeDifferences(differences)
}

func extrapolateRight(history []int) (int, error) {
	allDifferences, err := computeDifferences([][]int{history})
	if err != nil {
		return 0, err
	}

	lastDifference := 0
	allDifferences = lo.Reverse(allDifferences)
	for _, differences := range allDifferences {
		lastVal, err := lo.Last(differences)
		if err != nil {
			return 0, err
		}
		lastDifference = lastVal + lastDifference
	}
	return lastDifference, nil
}

func extrapolateLeft(history []int) (int, error) {
	allDifferences, err := computeDifferences([][]int{history})
	if err != nil {
		return 0, err
	}

	lastDifference := 0
	allDifferences = lo.Reverse(allDifferences)
	for _, differences := range allDifferences {
		firstVal := differences[0]
		lastDifference = firstVal - lastDifference
	}
	return lastDifference, nil
}
