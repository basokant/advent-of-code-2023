package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	calibrationValues := parseInput(input, false)

	var sum int
	for _, value := range calibrationValues {
		sum += value
	}
	return sum
}

func part2(input string) int {
	calibrationValues := parseInput(input, true)

	var sum int
	for _, value := range calibrationValues {
		sum += value
	}
	return sum
}

func parseInput(input string, shouldReplaceDigitWords bool) []int {
	calibrationValues := []int{}

	if shouldReplaceDigitWords {
		input = replaceDigitWords(input)
	}

	for _, line := range strings.Split(input, "\n") {
		digits := []int{}
		chars := []rune(line)

		for i := 0; i < len(line); i++ {
			char := chars[i]
			str := string(char)
			if unicode.IsDigit(char) {
				digit, _ := strconv.Atoi(str)
				digits = append(digits, digit)
			}
		}

		calibrationValue := digits[0]*10 + digits[len(digits)-1]
		calibrationValues = append(calibrationValues, calibrationValue)
	}
	return calibrationValues
}

func replaceDigitWords(input string) string {
	digitToValueMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	convertedInput := strings.Clone(input)

	for digitWord, digit := range digitToValueMap {
		digitStr := fmt.Sprint(digit)
		convertedInput = strings.ReplaceAll(convertedInput, digitWord, digitWord+digitStr+digitWord)
	}
	return convertedInput
}
