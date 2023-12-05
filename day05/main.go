package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/basokant/advent-of-code-2023/util"
)

//go:embed input.txt
var input string

type Mapping struct {
	src      int
	dest     int
	rangeLen int
}

func (m Mapping) isInRange(src int) bool {
	if m.src <= src && src < m.src+m.rangeLen {
		return true
	}
	return false
}

func (m Mapping) getDest(src int) int {
	if m.isInRange(src) {
		diff := src - m.src
		return m.dest + diff
	}
	return src
}

func compareMappingSrcRange(a Mapping, b Mapping) int {
	if a.isInRange(b.src) && b.isInRange(a.src) {
		return 0
	} else if a.src > b.src {
		return 1
	}
	return -1
}

func compareMappingWithSrc(mapping Mapping, src int) int {
	if mapping.isInRange(src) {
		return 0
	} else if mapping.src > src {
		return 1
	}
	return -1
}

type Map struct {
	mappings []Mapping
}

func (m Map) getDest(src int) int {
	mappingIndex, found := slices.BinarySearchFunc(m.mappings, src, compareMappingWithSrc)
	if !found {
		return src
	}
	return m.mappings[mappingIndex].getDest(src)
}

func NewMap(mappings []Mapping) Map {
	slices.SortFunc(mappings, compareMappingSrcRange)
	return Map{mappings}
}

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
	seeds, maps := parseInput(input)

	locations := make([]int, len(seeds))
	for i, seed := range seeds {
		output := seed
		for _, m := range maps {
			output = m.getDest(output)
		}
		locations[i] = output
	}

	lowestLocation := slices.Min(locations)
	return lowestLocation
}

func part2(input string) int {
	seedRanges, maps := parseInput(input)
	numSeedRanges := len(seedRanges) / 2

	locations := []int{}
	for i := 0; i < numSeedRanges; i++ {
		startSeed := seedRanges[i]
		rangeLen := seedRanges[i+1]

		for seed := startSeed; seed < startSeed+rangeLen; seed++ {
			output := seed
			for _, m := range maps {
				output = m.getDest(output)
			}
			locations = append(locations, output)
		}
	}

	lowestLocation := slices.Min(locations)
	return lowestLocation
}

func parseInput(input string) ([]int, []Map) {
	re := regexp.MustCompile("[0-9]+")
	seedsLine := strings.Split(input, "\n")[0]

	seedsStr := re.FindAllString(seedsLine, -1)
	seeds := make([]int, len(seedsStr))

	for i, seedStr := range seedsStr {
		seedNum, _ := strconv.Atoi(seedStr)
		seeds[i] = seedNum
	}

	mapsInput := strings.Split(input, "\n\n")
	maps := make([]Map, len(mapsInput)-1)

	for i, mapInput := range mapsInput[1:] {
		maps[i] = parseMapInput(mapInput)
	}

	return seeds, maps
}

func parseMapInput(mapInput string) Map {
	lines := strings.Split(mapInput, "\n")

	mappings := make([]Mapping, len(lines)-1)

	for i, line := range lines[1:] {
		nums := strings.Split(line, " ")
		dest, _ := strconv.Atoi(nums[0])
		src, _ := strconv.Atoi(nums[1])
		rangeLen, _ := strconv.Atoi(nums[2])
		mappings[i] = Mapping{
			src,
			dest,
			rangeLen,
		}
	}

	return NewMap(mappings)
}
