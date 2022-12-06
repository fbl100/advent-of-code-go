package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/algos"
	"github.com/alexchao26/advent-of-code-go/cast"
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
	count := 0
	for _, p := range parsed {
		l := p.Left.ToSet()
		r := p.Right.ToSet()

		if l.IsSuperset(r) || r.IsSuperset(l) {
			count += 1
		}
	}

	return count
}

func part2(input string) int {
	parsed := parseInput(input)
	count := 0
	for _, p := range parsed {
		l := p.Left.ToSet()
		r := p.Right.ToSet()
		i := l.Intersection(r)
		if len(i) > 0 {
			count++
		}
	}

	return count
}

type Range struct {
	From int
	To   int
}

type RangePair struct {
	Left  Range
	Right Range
}

func (r Range) ToSet() sets.Int {
	ret_val := sets.NewInt()
	for i := r.From; i <= r.To; i++ {
		ret_val.Insert(i)
	}
	return ret_val
}

func (r Range) String() string {
	return fmt.Sprintf("%d to %d", r.From, r.To)
}

func (r RangePair) String() string {
	return fmt.Sprintf("left: %s right %s", r.Left, r.Right)
	
}

func parseInput(input string) (ans []RangePair) {
	for _, line := range strings.Split(input, "\n") {

		arr := algos.SplitStringOn(line, []string{"-", ","})
		r1 := Range{
			From: cast.ToInt(arr[0]),
			To:   cast.ToInt(arr[1]),
		}
		r2 := Range{
			From: cast.ToInt(arr[2]),
			To:   cast.ToInt(arr[3]),
		}

		ans = append(ans, RangePair{Left: r1, Right: r2})
	}
	return ans
}
