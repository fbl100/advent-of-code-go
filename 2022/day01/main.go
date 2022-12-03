package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
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

	max := 0
	for _, lst := range parsed {
		calories := sum(lst)
		if calories > max {
			max = calories
		}
	}

	return max
}

func part2(input string) int {
	parsed := parseInput(input)

	var sums []int

	for _, lst := range parsed {
		calories := sum(lst)
		sums = append(sums, calories)
	}

	sort.Ints(sums)
	reverse(sums)

	for _, x := range sums {
		println(x)
	}

	return sums[0] + sums[1] + sums[2]

}

func parseInput(input string) (ans [][]int) {

	var curr []int

	for _, line := range strings.Split(input, "\n") {

		if line != "" {
			curr = append(curr, cast.ToInt(line))

		} else {
			ans = append(ans, curr)
			curr = make([]int, 0) // clear
		}
	}
	return ans
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func sum(arr []int) int {
	ans := 0
	for _, x := range arr {
		ans += x
	}
	return ans
}
