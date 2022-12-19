package main

import (
	_ "embed"
	"flag"
	"fmt"
	"golang.org/x/exp/slices"
	"sort"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

type PacketArray []string

func (p PacketArray) Len() int {
	//TODO implement me
	return len(p)
}

func (p PacketArray) Less(i, j int) bool {
	return compare(p[i], p[j]) == CORRECT
}

func (p PacketArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type CompareType int

const (
	CORRECT   CompareType = 0
	INCORRECT             = 1
	NA                    = 2
)

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

	ans := 0
	index := 1

	for i := 0; i < len(parsed); i += 3 {
		fmt.Println(parsed[i])
		fmt.Println(parsed[i+1])

		if compare(parsed[i], parsed[i+1]) == CORRECT {
			ans += index
			fmt.Println(index, " - ", ans)
		}
		fmt.Println("--")
		index++
	}
	return ans
}

func part2(input string) int {
	parsed := parseInput(input)
	filtered := PacketArray{}
	for _, s := range parsed {
		if strings.TrimSpace(s) != "" {
			filtered = append(filtered, s)
		}
	}

	filtered = append(filtered, "[[2]]")
	filtered = append(filtered, "[[6]]")

	sort.Sort(filtered)

	a := slices.IndexFunc(filtered, func(s string) bool { return s == "[[2]]" })
	b := slices.IndexFunc(filtered, func(s string) bool { return s == "[[6]]" })

	ans := (a + 1) * (b + 1)
	return ans
}

func parseInput(input string) (ans PacketArray) {

	return strings.Split(input, "\n")
}

func parseNode(s string) []string {
	// if s starts with '[' and ends with ']', then this is a list
	isList := strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]")

	retVal := []string{}

	if isList {

		if s == "[]" {
			return retVal
		}

		lastBracketIndex := len(s) - 1
		// get the content and look for commas, splitting only when we find commas at the top level
		startIndex := 1
		endIndex := startIndex
		level := 0

		for ; endIndex < lastBracketIndex; endIndex++ {
			//fmt.Println(string(s[endIndex]), ",", level, ",", endIndex)
			//if endIndex == 11 {
			//	println("11")
			//}
			if s[endIndex] == '[' {
				level++
			} else if s[endIndex] == ']' {
				level--
			} else if s[endIndex] == ',' && level == 0 {
				// extract from startIndex to endIndex
				x := s[startIndex:endIndex]
				retVal = append(retVal, x)
				startIndex = endIndex + 1
			}
		}

		// last item
		x := s[startIndex:endIndex]
		retVal = append(retVal, x)
		startIndex = endIndex + 1

	}
	return retVal
}

func compare(left, right string) CompareType {
	// both values are integers
	// if left < right - true
	// if left == right - confinue
	// if left > right - false
	leftArray := parseNode(left)
	rightArray := parseNode(right)

	i := 0

	for ; i < len(leftArray) && i < len(rightArray); i++ {
		fmt.Printf("Compare %v vs %v\n", leftArray[i], rightArray[i])

		left_i := leftArray[i]
		right_i := rightArray[i]

		leftIsList := isList(left_i)
		rightIsList := isList(right_i)

		if leftIsList && rightIsList {
			compare_i := compare(left_i, right_i)
			if compare_i != NA {
				return compare_i
			} else {
				continue
			}
		}

		li, li_err := strconv.Atoi(left_i)
		ri, ri_err := strconv.Atoi(right_i)

		if li_err == nil && ri_err != nil {
			fmt.Printf("Mixed types; convert left to [%v] and retry comparison\n", left_i)
			// left is int, right is not
			compare_i := compare(list(left_i), right_i)
			if compare_i != NA {
				return compare_i
			} else {
				continue
			}
		}

		if li_err != nil && ri_err == nil {
			fmt.Printf("Mixed types; convert left to [%v] and retry comparison\n", right_i)
			// left is int, right is not
			compare_i := compare(left_i, list(right_i))
			if compare_i != NA {
				return compare_i
			} else {
				continue
			}
		}

		if li_err != nil && ri_err != nil {
			panic("now what?")
		}

		if li > ri {
			fmt.Println("Left > Right, INCORRECT")
			return INCORRECT
		}

		if li < ri {
			fmt.Println("Left < Right, CORRECT")
			return CORRECT
		}
	}

	if i < len(rightArray) {
		fmt.Println("Left ran out of items first, CORRECT")
		return CORRECT
	}

	if i < len(leftArray) {
		fmt.Println("Right ran out of items first, INCORRECT")
		return INCORRECT
	}

	return NA

}

func list(s string) string {
	return "[" + s + "]"
}

func isList(s string) bool {
	return strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]")
}
