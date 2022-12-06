package main

import (
	_ "embed"
	"flag"
	"fmt"
	util2 "github.com/alexchao26/advent-of-code-go/2022/day05/util"
	"golang.org/x/exp/slices"
	"regexp"
	"strconv"
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

func part1(input string) string {
	board, moves := parseInput(input)

	println("Starting State")
	board.Print()

	for _, m := range moves {
		board.Apply(m)
		println()
		board.Print()
	}

	println("Ending State")
	board.Print()

	_ = board
	_ = moves

	return board.Top()
}

func part2(input string) string {
	board, moves := parseInput(input)

	println("Starting State")
	board.Print()

	for _, m := range moves {
		board.Apply2(m)
		println()
		board.Print()
	}

	println("Ending State")
	board.Print()

	_ = board
	_ = moves

	return board.Top()
}

func parseInput(input string) (*util2.Board, []*util2.Move) {
	// read the input until we get to a blank line
	lines := strings.Split(input, "\n")

	firstBlank := slices.IndexFunc(lines, func(s string) bool { return len(s) == 0 })

	//00000000001111111111222222222233333
	//01234567890123456789012345678901234
	//[G] [G] [G] [N] [V] [V] [T] [Q] [F]
	// 1   2   3   4   5   6   7   8   9

	indices := []int{1, 5, 9, 13, 17, 21, 25, 29, 33}

	bottom := firstBlank - 2

	b := util2.NewBoard(9)

	for i := bottom; i >= 0; i-- {
		line := strings.TrimRight(lines[i], " ")
		for stack_index, string_index := range indices {
			if string_index < len(line) {
				x := string(line[string_index])
				if x != " " {
					b.Stacks[stack_index].Push(string(line[string_index]))
				}
			}
		}
	}
	moves := []*util2.Move{}

	regex := *regexp.MustCompile(`move (\d+) from (\d) to (\d)`)
	for i := firstBlank + 1; i < len(lines); i++ {
		res := regex.FindStringSubmatch(lines[i])
		count, _ := strconv.Atoi(res[1])
		from, _ := strconv.Atoi(res[2])
		to, _ := strconv.Atoi(res[3])

		m := util2.Move{
			Count: count,
			From:  from,
			To:    to,
		}

		moves = append(moves, &m)
	}

	return b, moves
	//for _, line := range strings.Split(input, "\n") {
	//	ans = append(ans, cast.ToInt(line))
	//}
	//return ans
}
