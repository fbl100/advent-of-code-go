package main

import (
	_ "embed"
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/util/sets"
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

	//println(evaluate("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	//println(evaluate("nppdvjthqldpwncqszvftbrmjlhg"))
	//println(evaluate("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	//println(evaluate("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
	return evaluate(input, 14)

}

func evaluate(input string, n int) int {

	startIndex := n - 1
	for marker := startIndex; marker < len(input); marker++ {

		s := sets.NewByte()
		//println(input[marker-3 : marker+1])
		windowStart := marker - n + 1
		windowEnd := marker + 1
		s.Insert([]byte(input)[windowStart:windowEnd]...)
		if len(s) == n {
			return marker + 1
		}

	}

	return -1

	//startIndex := n - 1
	//for marker := startIndex; marker < len(input); marker++ {
	//
	//	s := sets.NewByte()
	//	//println(input[marker-3 : marker+1])
	//	windowStart := marker-(n+1)
	//	windowEnd := marker+1
	//	s.Insert([]byte(input)[marker-3 : marker+1]...)
	//	if len(s) == 4 {
	//		return marker + 1
	//	}
	//
	//}
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []int) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, cast.ToInt(line))
	}
	return ans
}
