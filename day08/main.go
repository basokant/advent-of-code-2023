package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math/big"
	"regexp"
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
	instructions, nodeMap := parseInput(input, false)
	return numStepsToZZZ("AAA", nodeMap, instructions)
}

func part2(input string) int64 {
	instructions, nodeMap := parseInput(input, false)

	nodes := lo.Filter(lo.Keys(nodeMap), func(node string, _ int) bool {
		return node[2] == 'A'
	})

	numStepsToZ := lo.Map(nodes, func(node string, _ int) *big.Int {
		return numStepsToZ(node, nodeMap, instructions)
	})

	return LCM(numStepsToZ...).Int64()
}

type Pair[T any] struct {
	left  T
	right T
}

func parseInput(input string, _ bool) (string, map[string]Pair[string]) {
	re := regexp.MustCompile(`\b\w+\b`)

	lines := strings.Split(input, "\n")
	instructions := lines[0]

	nodeMap := make(map[string]Pair[string])
	for _, line := range lines[2:] {
		matches := re.FindAllString(line, -1)
		node, left, right := matches[0], matches[1], matches[2]
		nodeMap[node] = Pair[string]{left, right}
	}

	return instructions, nodeMap
}

func numStepsToZZZ(node string, nodeMap map[string]Pair[string], instructions string) int {
	next := 0
	numSteps := 0
	for node != "ZZZ" {
		step := instructions[next]
		switch step {
		case 'L':
			node = nodeMap[node].left
		case 'R':
			node = nodeMap[node].right
		}

		next = (next + 1) % len(instructions)
		numSteps += 1
	}

	return numSteps
}

func numStepsToZ(node string, nodeMap map[string]Pair[string], instructions string) *big.Int {
	next := 0
	var numSteps int64 = 0

	for node[2] != 'Z' {
		step := instructions[next]
		switch step {
		case 'L':
			node = nodeMap[node].left
		case 'R':
			node = nodeMap[node].right
		}

		next = (next + 1) % len(instructions)
		numSteps += 1
	}
	return big.NewInt(numSteps)
}

// gcd calculates the Greatest Common Divisor using Euclid's algorithm
func gcd(a, b *big.Int) *big.Int {
	for b.Cmp(big.NewInt(0)) != 0 {
		temp := new(big.Int)
		temp.Set(b)
		b.Mod(a, b)
		a.Set(temp)
	}
	return a
}

// lcm calculates the Least Common Multiple using the formula: LCM(a, b) = |a * b| / GCD(a, b)
func lcm(a, b *big.Int) *big.Int {
	product := new(big.Int).Mul(a, b)
	gcdVal := gcd(a, b)
	if gcdVal.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}
	return product.Div(product, gcdVal)
}

func LCM(nums ...*big.Int) *big.Int {
	if len(nums) == 0 {
		return big.NewInt(0)
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = lcm(result, nums[i])
	}
	return result
}
