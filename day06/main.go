package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"

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

func computeDistance(speed int, time int) int {
	return speed * (time - speed)
}

func part1(input string) int {
	races := parseInput(input)

	numWinPossibilities := 1
	for _, race := range races {
		numWinSpeeds := 0
		for speed := 0; speed <= race.time; speed++ {
			distance := computeDistance(speed, race.time)
			if distance > race.distance {
				numWinSpeeds += 1
			}
		}
		numWinPossibilities *= numWinSpeeds
	}

	return numWinPossibilities
}

func part2(input string) int {
	race := parseInputPart2(input)

	numWinSpeeds := 0
	for speed := 0; speed <= race.time; speed++ {
		distance := computeDistance(speed, race.time)
		if distance > race.distance {
			numWinSpeeds += 1
		}
	}

	return numWinSpeeds
}

type Race struct {
	time     int
	distance int
}

func parseInputPart2(input string) Race {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile("[0-9]+")

	timesDigits := re.FindAllString(lines[0], -1)
	distancesDigits := re.FindAllString(lines[1], -1)

	time, _ := strconv.Atoi(strings.Join(timesDigits, ""))
	distance, _ := strconv.Atoi(strings.Join(distancesDigits, ""))

	return Race{time, distance}
}

func parseInput(input string) []Race {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile("[0-9]+")

	timesStr := re.FindAllString(lines[0], -1)
	distancesStr := re.FindAllString(lines[1], -1)

	races := []Race{}

	for i, timeStr := range timesStr {
		time, _ := strconv.Atoi(timeStr)
		distance, _ := strconv.Atoi(distancesStr[i])
		race := Race{time, distance}
		races = append(races, race)
	}

	return races
}
