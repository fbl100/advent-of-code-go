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

	shapes := make(map[string]string)
	shapes["A"] = "R"
	shapes["B"] = "P"
	shapes["C"] = "S"
	shapes["X"] = "R"
	shapes["Y"] = "P"
	shapes["Z"] = "S"

	scores := make(map[string]int)
	scores["R"] = 1
	scores["P"] = 2
	scores["S"] = 3

	score := 0
	for _, game := range parsed {

		moves := strings.Fields(game)
		opp := shapes[moves[0]]
		me := shapes[moves[1]]

		score += scoreGame(opp, me)

	}

	return score
}

func scoreGame(opp string, me string) int {
	scores := make(map[string]int)
	scores["R"] = 1
	scores["P"] = 2
	scores["S"] = 3

	if opp == me {
		// tie
		return 3 + scores[me]
	} else if (me == "R" && opp == "S") ||
		(me == "P" && opp == "R") ||
		(me == "S" && opp == "P") {
		return 6 + scores[me]
	} else {
		// didn't tie, didn't win
		return scores[me]
	}
}

func scoreGame2(opp string, result int) int {
	scores := make(map[string]int)
	scores["R"] = 1
	scores["P"] = 2
	scores["S"] = 3

	key := map[string]map[int]int{
		"R": {
			0: scores["S"], // opp picked rock, I lost, picked scissors
			3: scores["R"], // opp picked rock, so did I
			6: scores["P"], // opp picked rock, I picked paper
		},
		"P": {
			0: scores["R"], // opp picked paper, I lost, I rock
			3: scores["P"], // opp picked paper, so did I
			6: scores["S"], // opp picked paper, I picked scissors
		},
		"S": {
			0: scores["P"], // opp picked scissors, I lost, I paper
			3: scores["S"], // opp picked scissors, so did I
			6: scores["R"], // opp picked scissors, I picked rock
		},
	}

	return result + key[opp][result]
}

func part2(input string) int {
	parsed := parseInput(input)

	shapes := make(map[string]string)
	shapes["A"] = "R"
	shapes["B"] = "P"
	shapes["C"] = "S"

	results := make(map[string]int)
	results["X"] = 0
	results["Y"] = 3
	results["Z"] = 6

	score := 0
	for _, game := range parsed {

		moves := strings.Fields(game)
		opp := shapes[moves[0]]
		result := results[moves[1]]

		score += scoreGame2(opp, result)

	}

	return score
}

func scoreShape(shape string) int {
	rocks := sets.NewString("A", "X")
	papers := sets.NewString("B", "Y")
	scissors := sets.NewString("C", "Z")

	if rocks.Has(shape) {
		return 1
	} else if papers.Has(shape) {
		return 2
	} else if scissors.Has(shape) {
		return 3
	} else {
		panic("bad shape")
	}
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")

}
