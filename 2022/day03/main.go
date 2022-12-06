package main

import (
	_ "embed"
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/util/sets"
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
	_ = parsed

	sum := 0
	for _, line := range parsed {

		length := len(line)
		if length%2 == 1 {
			panic("odd numbered input string!")
		}
		idx := length / 2
		first := line[0:idx]
		second := line[idx:length]
		if len(first)+len(second) != length {
			panic("bad split")
		}

		c := commonLetters(first, second)
		sum += c.List()[0]
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	sum := 0
	for i := 0; i < len(parsed); i += 3 {
		a := digits(parsed[i])
		b := digits(parsed[i+1])
		c := digits(parsed[i+2])

		x := a.Intersection(b)
		x = x.Intersection(c)

		if len(x) != 1 {
			panic("expected one value")
		}

		sum += x.List()[0]

	}
	return sum
	//for _, line := range parsed {
	//
	//	length := len(line)
	//	if length%2 == 1 {
	//		panic("odd numbered input string!")
	//	}
	//	idx := length / 2
	//	first := line[0:idx]
	//	second := line[idx:length]
	//	if len(first)+len(second) != length {
	//		panic("bad split")
	//	}
	//
	//	c := commonLetters(first, second)
	//	sum += c.List()[0]
	//}

	return sum
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")

}

func commonLetters(a string, b string) sets.Int {

	digitsA := digits(a)
	digitsB := digits(b)

	common := digitsA.Intersection(digitsB)
	return common

}

func digits(a string) sets.Int {
	retVal := sets.NewInt()
	for _, c := range a {
		retVal.Insert(priority(c))
	}
	return retVal

}

func priority(a rune) int {
	if a >= 'a' && a <= 'z' {
		return int(a) - int('a') + 1
	} else if a >= 'A' && a <= 'Z' {
		return int(a) - int('A') + 27
	} else {
		panic("bad character")
	}
}
