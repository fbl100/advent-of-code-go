package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/data-structures/grid"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
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
	parsed := parseInput(input)

	cube := grid.NewInfiniteCube(".")
	for _, c := range parsed {
		cube.PutWithCoord(c, "#")
	}

	count := 0
	for _, c := range parsed {
		for _, n := range c.Neighbors() {
			value, _ := cube.GetWithCoord(n)
			if value == cube.Default {
				count++
			}
		}
	}

	return count
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []grid.Coord3d) {
	for _, line := range strings.Split(input, "\n") {
		tokens := strings.Split(line, ",")
		ans = append(ans, grid.Coord3d{cast.ToInt(tokens[0]),
			cast.ToInt(tokens[1]),
			cast.ToInt(tokens[2])})
	}
	return ans
}
